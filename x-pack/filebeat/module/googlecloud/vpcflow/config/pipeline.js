// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

var vpcflow = (function () {
    var processor = require("processor");

    // The pub/sub input writes the Stackdriver LogEntry object into the message
    // field. The message needs decoded as JSON.
    var decodeJson = new processor.DecodeJSONFields({
        fields: ["message"],
        target: "json",
    });

    var dropPubSubFields = function(evt) {
        evt.Delete("message");
        evt.Delete("labels");
    };

    var categorizeEvent = new processor.AddFields({
        target: "event",
        fields: {
            category: "network_traffic",
            type: "flow",
        },
    });

    // Use the LogEntry object's timestamp. VPC flow logs are structured so the
    // LogEntry includes a jsonPayload field.
    // https://cloud.google.com/logging/docs/reference/v2/rest/v2/LogEntry
    var convertLogEntry = new processor.Convert({
        fields: [
            {from: "json.timestamp", to: "@metadata.timestamp"},
            {from: "json.jsonPayload", to: "json"},
        ],
        mode: "rename",
    });

    // Set @timetamp the LogEntry's timestamp.
    var parseTimestamp = function(evt) {
        var ts = evt.Get("@metadata.timestamp");
        if (!ts) {
            return;
        }
        evt.Delete("@metadata.timestamp");

        var ms = Date.parse(ts);
        if (isNaN(ms)) {
            return;
        }

        var date = new Date(ms);
        evt.Put("@timestamp", date);
    };

    // The LogEntry's jsonPayload is moved to the json field. The jsonPayload
    // contains the structured VPC flow log fields.
    // https://cloud.google.com/vpc/docs/using-flow-logs#record_format
    var convertJsonPayload = new processor.Convert({
        fields: [
            {from: "json.connection.dest_ip", to: "destination.address"},
            {from: "json.connection.dest_port", to: "destination.port", type: "long"},
            {from: "json.connection.protocol", to: "network.iana_number", type: "string"},
            {from: "json.connection.src_ip", to: "source.address"},
            {from: "json.connection.src_port", to: "source.port", type: "long"},

            {from: "json.src_instance.vm_name", to: "source.domain"},
            {from: "json.dest_instance.vm_name", to: "destination.domain"},

            {from: "json.bytes_sent", to: "source.bytes", type: "long"},
            {from: "json.packets_sent", to: "source.packets", type: "long"},

            {from: "json.start_time", to: "event.start"},
            {from: "json.end_time", to: "event.end"},

            {from: "json.dest_location.asn", to: "destination.as.number", type: "long"},
            {from: "json.dest_location.continent", to: "destination.geo.continent_name"},
            {from: "json.dest_location.country", to: "destination.geo.country_name"},
            {from: "json.dest_location.region", to: "destination.geo.region_name"},
            {from: "json.dest_location.city", to: "destination.geo.city_name"},

            {from: "json.src_location.asn", to: "source.as.number", type: "long"},
            {from: "json.src_location.continent", to: "source.geo.continent_name"},
            {from: "json.src_location.country", to: "source.geo.country_name"},
            {from: "json.src_location.region", to: "source.geo.region_name"},
            {from: "json.src_location.city", to: "source.geo.city_name"},

            {from: "json.dest_instance", to: "json.destination.instance"},
            {from: "json.dest_vpc", to: "json.destination.vpc"},
            {from: "json.src_instance", to: "json.source.instance"},
            {from: "json.src_vpc", to: "json.source.vpc"},

            {from: "json.rtt_msec", to: "json.rtt.ms", type: "long"},
            {from: "json", to: "googlecloud.vpcflow"},
        ],
        mode: "rename",
        ignore_missing: true,
    });

    // Delete emtpy object's whose fields have been renamed leaving them childless.
    var dropEmptyObjects = function (evt) {
        evt.Delete("googlecloud.vpcflow.connection");
        evt.Delete("googlecloud.vpcflow.dest_location");
        evt.Delete("googlecloud.vpcflow.src_location");
    };

    // Copy the source/destination.address to source/destination.ip if they are
    // valid IP addresses.
    var copyAddressFields = new processor.Convert({
        fields: [
            {from: "source.address", to: "source.ip", type: "ip"},
            {from: "destination.address", to: "destination.ip", type: "ip"},
        ],
        fail_on_error: false,
    });

    var setCloudFromDestInstance = new processor.Convert({
        fields: [
            {from: "googlecloud.vpcflow.destination.instance.project_id", to: "cloud.project.id"},
            {from: "googlecloud.vpcflow.destination.instance.vm_name", to: "cloud.instance.name"},
            {from: "googlecloud.vpcflow.destination.instance.region", to: "cloud.region"},
            {from: "googlecloud.vpcflow.destination.instance.zone", to: "cloud.availability_zone"},
            {from: "googlecloud.vpcflow.destination.vpc.subnetwork_name", to: "network.name"},
        ],
        ignore_missing: true,
    });

    var setCloudFromSrcInstance = new processor.Convert({
        fields: [
            {from: "googlecloud.vpcflow.source.instance.project_id", to: "cloud.project.id"},
            {from: "googlecloud.vpcflow.source.instance.vm_name", to: "cloud.instance.name"},
            {from: "googlecloud.vpcflow.source.instance.region", to: "cloud.region"},
            {from: "googlecloud.vpcflow.source.instance.zone", to: "cloud.availability_zone"},
            {from: "googlecloud.vpcflow.source.vpc.subnetwork_name", to: "network.name"},
        ],
        ignore_missing: true,
    });

    // Set the cloud metadata fields based on the instance that reported the
    // event.
    var setCloudMetadata = function(evt) {
        var reporter = evt.Get("googlecloud.vpcflow.reporter");

        if (reporter === "DEST") {
            setCloudFromDestInstance.Run(evt);
        } else if (reporter === "SRC") {
            setCloudFromSrcInstance.Run(evt);
        }
    };

    var communityId = new processor.CommunityID({
        fields: {
            transport: "network.iana_number",
        }
    });

    // VPC flows are unidirectional so we only have to worry about copy the
    // source.bytes/packets over to network.bytes/packets.
    var setNetworkBytesPackets = new processor.Convert({
        fields: [
            {from: "source.bytes", to: "network.bytes"},
            {from: "source.packets", to: "network.packets"},
        ],
        ignore_missing: true,
    });

    // VPC flow logs are reported for TCP and UDP traffic only so handle these
    // protocols' IANA numbers.
    var setNetworkTransport = function(event) {
        var ianaNumber = event.Get("network.iana_number");
        switch (ianaNumber) {
            case "6":
                event.Put("network.transport", "tcp");
                break;
            case "17":
                event.Put("network.transport", "udp");
                break;
        }
    };

    var setNetworkDirection = function(event) {
        var srcInstance = event.Get("googlecloud.vpcflow.source.instance");
        var destInstance = event.Get("googlecloud.vpcflow.destination.instance");
        var direction = "unknown";

        if (srcInstance && destInstance) {
            direction = "internal";
        } else if (srcInstance) {
            direction = "outbound";
        } else if (destInstance) {
            direction = "inbound";
        }
        event.Put("network.direction", direction);
    };

    var setNetworkType = function(event) {
        var ip = event.Get("source.ip");
        if (!ip) {
            return;
        }

        if (ip.indexOf(".") !== -1) {
            event.Put("network.type", "ipv4");
        } else {
            event.Put("network.type", "ipv6");
        }
    };

    var setRelatedIP = function(event) {
        event.AppendTo("related.ip", event.Get("source.ip"));
        event.AppendTo("related.ip", event.Get("destination.ip"));
    };

    var pipeline = new processor.Chain()
        .Add(decodeJson)
        .Add(dropPubSubFields)
        .Add(categorizeEvent)
        .Add(convertLogEntry)
        .Add(parseTimestamp)
        .Add(convertJsonPayload)
        .Add(dropEmptyObjects)
        .Add(copyAddressFields)
        .Add(setCloudMetadata)
        .Add(communityId)
        .Add(setNetworkBytesPackets)
        .Add(setNetworkTransport)
        .Add(setNetworkDirection)
        .Add(setNetworkType)
        .Add(setRelatedIP)
        .Build();

    return {
        process: pipeline.Run,
    };
})();

function process(evt) {
    return vpcflow.process(evt);
}

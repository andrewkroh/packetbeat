// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeFlowLogsRequest
type DescribeFlowLogsInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`

	// One or more filters.
	//
	//    * deliver-log-status - The status of the logs delivery (SUCCESS | FAILED).
	//
	//    * log-destination-type - The type of destination to which the flow log
	//    publishes data. Possible destination types include cloud-watch-logs and
	//    S3.
	//
	//    * flow-log-id - The ID of the flow log.
	//
	//    * log-group-name - The name of the log group.
	//
	//    * resource-id - The ID of the VPC, subnet, or network interface.
	//
	//    * traffic-type - The type of traffic (ACCEPT | REJECT | ALL).
	Filter []Filter `locationNameList:"Filter" type:"list"`

	// One or more flow log IDs.
	//
	// Constraint: Maximum of 1000 flow log IDs.
	FlowLogIds []string `locationName:"FlowLogId" locationNameList:"item" type:"list"`

	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	MaxResults *int64 `type:"integer"`

	// The token for the next page of results.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s DescribeFlowLogsInput) String() string {
	return awsutil.Prettify(s)
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeFlowLogsResult
type DescribeFlowLogsOutput struct {
	_ struct{} `type:"structure"`

	// Information about the flow logs.
	FlowLogs []FlowLog `locationName:"flowLogSet" locationNameList:"item" type:"list"`

	// The token to use to retrieve the next page of results. This value is null
	// when there are no more results to return.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s DescribeFlowLogsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeFlowLogs = "DescribeFlowLogs"

// DescribeFlowLogsRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Describes one or more flow logs. To view the information in your flow logs
// (the log streams for the network interfaces), you must use the CloudWatch
// Logs console or the CloudWatch Logs API.
//
//    // Example sending a request using DescribeFlowLogsRequest.
//    req := client.DescribeFlowLogsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeFlowLogs
func (c *Client) DescribeFlowLogsRequest(input *DescribeFlowLogsInput) DescribeFlowLogsRequest {
	op := &aws.Operation{
		Name:       opDescribeFlowLogs,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &DescribeFlowLogsInput{}
	}

	req := c.newRequest(op, input, &DescribeFlowLogsOutput{})
	return DescribeFlowLogsRequest{Request: req, Input: input, Copy: c.DescribeFlowLogsRequest}
}

// DescribeFlowLogsRequest is the request type for the
// DescribeFlowLogs API operation.
type DescribeFlowLogsRequest struct {
	*aws.Request
	Input *DescribeFlowLogsInput
	Copy  func(*DescribeFlowLogsInput) DescribeFlowLogsRequest
}

// Send marshals and sends the DescribeFlowLogs API request.
func (r DescribeFlowLogsRequest) Send(ctx context.Context) (*DescribeFlowLogsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeFlowLogsResponse{
		DescribeFlowLogsOutput: r.Request.Data.(*DescribeFlowLogsOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeFlowLogsRequestPaginator returns a paginator for DescribeFlowLogs.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeFlowLogsRequest(input)
//   p := ec2.NewDescribeFlowLogsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeFlowLogsPaginator(req DescribeFlowLogsRequest) DescribeFlowLogsPaginator {
	return DescribeFlowLogsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribeFlowLogsInput
				if req.Input != nil {
					tmp := *req.Input
					inCpy = &tmp
				}

				newReq := req.Copy(inCpy)
				newReq.SetContext(ctx)
				return newReq.Request, nil
			},
		},
	}
}

// DescribeFlowLogsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeFlowLogsPaginator struct {
	aws.Pager
}

func (p *DescribeFlowLogsPaginator) CurrentPage() *DescribeFlowLogsOutput {
	return p.Pager.CurrentPage().(*DescribeFlowLogsOutput)
}

// DescribeFlowLogsResponse is the response type for the
// DescribeFlowLogs API operation.
type DescribeFlowLogsResponse struct {
	*DescribeFlowLogsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeFlowLogs request.
func (r *DescribeFlowLogsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}

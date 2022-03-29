// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package http_endpoint

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/elastic/beats/v7/libbeat/beat"
	"github.com/elastic/beats/v7/libbeat/common"
	"github.com/elastic/beats/v7/libbeat/logp"

	"github.com/stretchr/testify/assert"
)

func Test_httpReadJSON(t *testing.T) {
	tests := []struct {
		name           string
		body           string
		wantObjs       []common.MapStr
		wantStatus     int
		wantErr        bool
		wantRawMessage []json.RawMessage
	}{
		{
			name:       "single object",
			body:       `{"a": 42, "b": "c"}`,
			wantObjs:   []common.MapStr{{"a": int64(42), "b": "c"}},
			wantStatus: http.StatusOK,
		},
		{
			name:       "array accepted",
			body:       `[{"a":"b"},{"c":"d"}]`,
			wantObjs:   []common.MapStr{{"a": "b"}, {"c": "d"}},
			wantStatus: http.StatusOK,
		},
		{
			name:       "not an object not accepted",
			body:       `42`,
			wantStatus: http.StatusBadRequest,
			wantErr:    true,
		},
		{
			name: "not an object mixed",
			body: `[{a:1},
								42,
							{a:2}]`,
			wantStatus: http.StatusBadRequest,
			wantErr:    true,
		},
		{
			name:       "sequence of objects accepted (CRLF)",
			body:       `{"a":1}` + "\r" + `{"a":2}`,
			wantObjs:   []common.MapStr{{"a": int64(1)}, {"a": int64(2)}},
			wantStatus: http.StatusOK,
		},
		{
			name: "sequence of objects accepted (LF)",
			body: `{"a":"1"}
									{"a":"2"}`,
			wantRawMessage: []json.RawMessage{
				[]byte(`{"a":"1"}`),
				[]byte(`{"a":"2"}`),
			},
			wantObjs:   []common.MapStr{{"a": "1"}, {"a": "2"}},
			wantStatus: http.StatusOK,
		},
		{
			name:       "sequence of objects accepted (SP)",
			body:       `{"a":"2"} {"a":"2"}`,
			wantObjs:   []common.MapStr{{"a": "2"}, {"a": "2"}},
			wantStatus: http.StatusOK,
		},
		{
			name:       "sequence of objects accepted (no separator)",
			body:       `{"a":"2"}{"a":"2"}`,
			wantObjs:   []common.MapStr{{"a": "2"}, {"a": "2"}},
			wantStatus: http.StatusOK,
		},
		{
			name: "not an object in sequence",
			body: `{"a":"2"}
									42
						 {"a":"2"}`,
			wantStatus: http.StatusBadRequest,
			wantErr:    true,
		},
		{
			name: "array of objects in stream",
			body: `{"a":"1"} [{"a":"2"},{"a":"3"}] {"a":"4"}`,
			wantRawMessage: []json.RawMessage{
				[]byte(`{"a":"1"}`),
				[]byte(`{"a":"2"}`),
				[]byte(`{"a":"3"}`),
				[]byte(`{"a":"4"}`),
			},
			wantObjs:   []common.MapStr{{"a": "1"}, {"a": "2"}, {"a": "3"}, {"a": "4"}},
			wantStatus: http.StatusOK,
		},
		{
			name: "numbers",
			body: `{"a":1} [{"a":false},{"a":3.14}] {"a":-4}`,
			wantRawMessage: []json.RawMessage{
				[]byte(`{"a":1}`),
				[]byte(`{"a":false}`),
				[]byte(`{"a":3.14}`),
				[]byte(`{"a":-4}`),
			},
			wantObjs: []common.MapStr{
				{"a": int64(1)},
				{"a": false},
				{"a": 3.14},
				{"a": int64(-4)},
			},
			wantStatus: http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotObjs, rawMessages, gotStatus, err := httpReadJSON(strings.NewReader(tt.body))
			if (err != nil) != tt.wantErr {
				t.Errorf("httpReadJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !assert.EqualValues(t, tt.wantObjs, gotObjs) {
				t.Errorf("httpReadJSON() gotObjs = %v, want %v", gotObjs, tt.wantObjs)
			}
			if gotStatus != tt.wantStatus {
				t.Errorf("httpReadJSON() gotStatus = %v, want %v", gotStatus, tt.wantStatus)
			}
			if tt.wantRawMessage != nil {
				assert.Equal(t, tt.wantRawMessage, rawMessages)
			}
			assert.Equal(t, len(gotObjs), len(rawMessages))
		})
	}
}

type publisher struct {
	events []beat.Event
}

func (p *publisher) Publish(event beat.Event) {
	p.events = append(p.events, event)
}

func Test_apiResponse(t *testing.T) {
	pub := new(publisher)
	conf := defaultConfig()
	apiHandler := newHandler(conf, pub, logp.NewLogger("http_endpoint.test"))

	const singleObjectJSON = `{"id":0}`

	t.Run("json object", func(t *testing.T) {
		pub.events = pub.events[:0]

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBufferString(singleObjectJSON))
		req.Header.Set("Content-Type", "application/json")
		respRec := httptest.NewRecorder()
		apiHandler(respRec, req)

		assert.Equal(t, http.StatusOK, respRec.Code)
		assert.Equal(t, conf.ResponseBody, respRec.Body.String())
		assert.Len(t, pub.events, 1)
	})

	t.Run("gzip json object", func(t *testing.T) {
		pub.events = pub.events[:0]

		buf := new(bytes.Buffer)
		b := gzip.NewWriter(buf)
		_, _ = io.WriteString(b, singleObjectJSON)
		b.Close()

		req := httptest.NewRequest(http.MethodPost, "/", buf)
		req.Header.Set("Content-Encoding", "gzip")
		req.Header.Set("Content-Type", "application/json")
		respRec := httptest.NewRecorder()
		apiHandler(respRec, req)

		assert.Equal(t, http.StatusOK, respRec.Code)
		assert.Equal(t, conf.ResponseBody, respRec.Body.String())
		assert.Len(t, pub.events, 1)
	})

	t.Run("gzip json objects", func(t *testing.T) {
		pub.events = pub.events[:0]

		const multipleObjectsJSON = `{"id":0}` + "\n" + `{"id":1}`
		buf := new(bytes.Buffer)
		b := gzip.NewWriter(buf)
		_, _ = io.WriteString(b, multipleObjectsJSON)
		b.Close()

		req := httptest.NewRequest(http.MethodPost, "/", buf)
		req.Header.Set("Content-Encoding", "gzip")
		req.Header.Set("Content-Type", "application/json")
		respRec := httptest.NewRecorder()
		apiHandler(respRec, req)

		assert.Equal(t, http.StatusOK, respRec.Code)
		assert.Equal(t, conf.ResponseBody, respRec.Body.String())
		assert.Len(t, pub.events, 2)

		for i, evt := range pub.events {
			id, _ := evt.GetValue("json.id")
			assert.EqualValues(t, i, id)
		}
	})
}

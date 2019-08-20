// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elasticloadbalancingv2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticloadbalancingv2-2015-12-01/DescribeRulesInput
type DescribeRulesInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the listener.
	ListenerArn *string `type:"string"`

	// The marker for the next set of results. (You received this marker from a
	// previous call.)
	Marker *string `type:"string"`

	// The maximum number of results to return with this call.
	PageSize *int64 `min:"1" type:"integer"`

	// The Amazon Resource Names (ARN) of the rules.
	RuleArns []string `type:"list"`
}

// String returns the string representation
func (s DescribeRulesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeRulesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeRulesInput"}
	if s.PageSize != nil && *s.PageSize < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("PageSize", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticloadbalancingv2-2015-12-01/DescribeRulesOutput
type DescribeRulesOutput struct {
	_ struct{} `type:"structure"`

	// If there are additional results, this is the marker for the next set of results.
	// Otherwise, this is null.
	NextMarker *string `type:"string"`

	// Information about the rules.
	Rules []Rule `type:"list"`
}

// String returns the string representation
func (s DescribeRulesOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeRules = "DescribeRules"

// DescribeRulesRequest returns a request value for making API operation for
// Elastic Load Balancing.
//
// Describes the specified rules or the rules for the specified listener. You
// must specify either a listener or one or more rules.
//
//    // Example sending a request using DescribeRulesRequest.
//    req := client.DescribeRulesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticloadbalancingv2-2015-12-01/DescribeRules
func (c *Client) DescribeRulesRequest(input *DescribeRulesInput) DescribeRulesRequest {
	op := &aws.Operation{
		Name:       opDescribeRules,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeRulesInput{}
	}

	req := c.newRequest(op, input, &DescribeRulesOutput{})
	return DescribeRulesRequest{Request: req, Input: input, Copy: c.DescribeRulesRequest}
}

// DescribeRulesRequest is the request type for the
// DescribeRules API operation.
type DescribeRulesRequest struct {
	*aws.Request
	Input *DescribeRulesInput
	Copy  func(*DescribeRulesInput) DescribeRulesRequest
}

// Send marshals and sends the DescribeRules API request.
func (r DescribeRulesRequest) Send(ctx context.Context) (*DescribeRulesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeRulesResponse{
		DescribeRulesOutput: r.Request.Data.(*DescribeRulesOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeRulesResponse is the response type for the
// DescribeRules API operation.
type DescribeRulesResponse struct {
	*DescribeRulesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeRules request.
func (r *DescribeRulesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}

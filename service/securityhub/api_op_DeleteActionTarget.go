// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package securityhub

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/securityhub-2018-10-26/DeleteActionTargetRequest
type DeleteActionTargetInput struct {
	_ struct{} `type:"structure"`

	// The ARN of the custom action target to delete.
	//
	// ActionTargetArn is a required field
	ActionTargetArn *string `location:"uri" locationName:"ActionTargetArn" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteActionTargetInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteActionTargetInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteActionTargetInput"}

	if s.ActionTargetArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("ActionTargetArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteActionTargetInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.ActionTargetArn != nil {
		v := *s.ActionTargetArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "ActionTargetArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/securityhub-2018-10-26/DeleteActionTargetResponse
type DeleteActionTargetOutput struct {
	_ struct{} `type:"structure"`

	// The ARN of the custom action target that was deleted.
	//
	// ActionTargetArn is a required field
	ActionTargetArn *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteActionTargetOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteActionTargetOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ActionTargetArn != nil {
		v := *s.ActionTargetArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "ActionTargetArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opDeleteActionTarget = "DeleteActionTarget"

// DeleteActionTargetRequest returns a request value for making API operation for
// AWS SecurityHub.
//
// Deletes a custom action target from Security Hub. Deleting a custom action
// target doesn't affect any findings or insights that were already sent to
// Amazon CloudWatch Events using the custom action.
//
//    // Example sending a request using DeleteActionTargetRequest.
//    req := client.DeleteActionTargetRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/securityhub-2018-10-26/DeleteActionTarget
func (c *Client) DeleteActionTargetRequest(input *DeleteActionTargetInput) DeleteActionTargetRequest {
	op := &aws.Operation{
		Name:       opDeleteActionTarget,
		HTTPMethod: "DELETE",
		HTTPPath:   "/actionTargets/{ActionTargetArn+}",
	}

	if input == nil {
		input = &DeleteActionTargetInput{}
	}

	req := c.newRequest(op, input, &DeleteActionTargetOutput{})
	return DeleteActionTargetRequest{Request: req, Input: input, Copy: c.DeleteActionTargetRequest}
}

// DeleteActionTargetRequest is the request type for the
// DeleteActionTarget API operation.
type DeleteActionTargetRequest struct {
	*aws.Request
	Input *DeleteActionTargetInput
	Copy  func(*DeleteActionTargetInput) DeleteActionTargetRequest
}

// Send marshals and sends the DeleteActionTarget API request.
func (r DeleteActionTargetRequest) Send(ctx context.Context) (*DeleteActionTargetResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteActionTargetResponse{
		DeleteActionTargetOutput: r.Request.Data.(*DeleteActionTargetOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteActionTargetResponse is the response type for the
// DeleteActionTarget API operation.
type DeleteActionTargetResponse struct {
	*DeleteActionTargetOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteActionTarget request.
func (r *DeleteActionTargetResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
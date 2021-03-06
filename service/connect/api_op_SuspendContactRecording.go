// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package connect

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type SuspendContactRecordingInput struct {
	_ struct{} `type:"structure"`

	// The identifier of the contact.
	//
	// ContactId is a required field
	ContactId *string `min:"1" type:"string" required:"true"`

	// The identifier of the contact. This is the identifier of the contact associated
	// with the first interaction with the contact center.
	//
	// InitialContactId is a required field
	InitialContactId *string `min:"1" type:"string" required:"true"`

	// The identifier of the Amazon Connect instance.
	//
	// InstanceId is a required field
	InstanceId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s SuspendContactRecordingInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *SuspendContactRecordingInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "SuspendContactRecordingInput"}

	if s.ContactId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ContactId"))
	}
	if s.ContactId != nil && len(*s.ContactId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ContactId", 1))
	}

	if s.InitialContactId == nil {
		invalidParams.Add(aws.NewErrParamRequired("InitialContactId"))
	}
	if s.InitialContactId != nil && len(*s.InitialContactId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("InitialContactId", 1))
	}

	if s.InstanceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("InstanceId"))
	}
	if s.InstanceId != nil && len(*s.InstanceId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("InstanceId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s SuspendContactRecordingInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.ContactId != nil {
		v := *s.ContactId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "ContactId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.InitialContactId != nil {
		v := *s.InitialContactId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "InitialContactId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.InstanceId != nil {
		v := *s.InstanceId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "InstanceId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type SuspendContactRecordingOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s SuspendContactRecordingOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s SuspendContactRecordingOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opSuspendContactRecording = "SuspendContactRecording"

// SuspendContactRecordingRequest returns a request value for making API operation for
// Amazon Connect Service.
//
// When a contact is being recorded, this API suspends recording the call. For
// example, you might suspend the call recording while collecting sensitive
// information, such as a credit card number. Then use ResumeContactRecording
// to restart recording.
//
// The period of time that the recording is suspended is filled with silence
// in the final recording.
//
// Only voice recordings are supported at this time.
//
//    // Example sending a request using SuspendContactRecordingRequest.
//    req := client.SuspendContactRecordingRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/connect-2017-08-08/SuspendContactRecording
func (c *Client) SuspendContactRecordingRequest(input *SuspendContactRecordingInput) SuspendContactRecordingRequest {
	op := &aws.Operation{
		Name:       opSuspendContactRecording,
		HTTPMethod: "POST",
		HTTPPath:   "/contact/suspend-recording",
	}

	if input == nil {
		input = &SuspendContactRecordingInput{}
	}

	req := c.newRequest(op, input, &SuspendContactRecordingOutput{})

	return SuspendContactRecordingRequest{Request: req, Input: input, Copy: c.SuspendContactRecordingRequest}
}

// SuspendContactRecordingRequest is the request type for the
// SuspendContactRecording API operation.
type SuspendContactRecordingRequest struct {
	*aws.Request
	Input *SuspendContactRecordingInput
	Copy  func(*SuspendContactRecordingInput) SuspendContactRecordingRequest
}

// Send marshals and sends the SuspendContactRecording API request.
func (r SuspendContactRecordingRequest) Send(ctx context.Context) (*SuspendContactRecordingResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &SuspendContactRecordingResponse{
		SuspendContactRecordingOutput: r.Request.Data.(*SuspendContactRecordingOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// SuspendContactRecordingResponse is the response type for the
// SuspendContactRecording API operation.
type SuspendContactRecordingResponse struct {
	*SuspendContactRecordingOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// SuspendContactRecording request.
func (r *SuspendContactRecordingResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}

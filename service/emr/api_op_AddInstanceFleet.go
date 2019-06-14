// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package emr

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticmapreduce-2009-03-31/AddInstanceFleetInput
type AddInstanceFleetInput struct {
	_ struct{} `type:"structure"`

	// The unique identifier of the cluster.
	//
	// ClusterId is a required field
	ClusterId *string `type:"string" required:"true"`

	// Specifies the configuration of the instance fleet.
	//
	// InstanceFleet is a required field
	InstanceFleet *InstanceFleetConfig `type:"structure" required:"true"`
}

// String returns the string representation
func (s AddInstanceFleetInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AddInstanceFleetInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AddInstanceFleetInput"}

	if s.ClusterId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ClusterId"))
	}

	if s.InstanceFleet == nil {
		invalidParams.Add(aws.NewErrParamRequired("InstanceFleet"))
	}
	if s.InstanceFleet != nil {
		if err := s.InstanceFleet.Validate(); err != nil {
			invalidParams.AddNested("InstanceFleet", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticmapreduce-2009-03-31/AddInstanceFleetOutput
type AddInstanceFleetOutput struct {
	_ struct{} `type:"structure"`

	// The unique identifier of the cluster.
	ClusterId *string `type:"string"`

	// The unique identifier of the instance fleet.
	InstanceFleetId *string `type:"string"`
}

// String returns the string representation
func (s AddInstanceFleetOutput) String() string {
	return awsutil.Prettify(s)
}

const opAddInstanceFleet = "AddInstanceFleet"

// AddInstanceFleetRequest returns a request value for making API operation for
// Amazon Elastic MapReduce.
//
// Adds an instance fleet to a running cluster.
//
// The instance fleet configuration is available only in Amazon EMR versions
// 4.8.0 and later, excluding 5.0.x.
//
//    // Example sending a request using AddInstanceFleetRequest.
//    req := client.AddInstanceFleetRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticmapreduce-2009-03-31/AddInstanceFleet
func (c *Client) AddInstanceFleetRequest(input *AddInstanceFleetInput) AddInstanceFleetRequest {
	op := &aws.Operation{
		Name:       opAddInstanceFleet,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AddInstanceFleetInput{}
	}

	req := c.newRequest(op, input, &AddInstanceFleetOutput{})
	return AddInstanceFleetRequest{Request: req, Input: input, Copy: c.AddInstanceFleetRequest}
}

// AddInstanceFleetRequest is the request type for the
// AddInstanceFleet API operation.
type AddInstanceFleetRequest struct {
	*aws.Request
	Input *AddInstanceFleetInput
	Copy  func(*AddInstanceFleetInput) AddInstanceFleetRequest
}

// Send marshals and sends the AddInstanceFleet API request.
func (r AddInstanceFleetRequest) Send(ctx context.Context) (*AddInstanceFleetResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AddInstanceFleetResponse{
		AddInstanceFleetOutput: r.Request.Data.(*AddInstanceFleetOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AddInstanceFleetResponse is the response type for the
// AddInstanceFleet API operation.
type AddInstanceFleetResponse struct {
	*AddInstanceFleetOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AddInstanceFleet request.
func (r *AddInstanceFleetResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
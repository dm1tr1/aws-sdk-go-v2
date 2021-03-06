// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package mediastore

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeContainerInput struct {
	_ struct{} `type:"structure"`

	// The name of the container to query.
	ContainerName *string `min:"1" type:"string"`
}

// String returns the string representation
func (s DescribeContainerInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeContainerInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeContainerInput"}
	if s.ContainerName != nil && len(*s.ContainerName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ContainerName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeContainerOutput struct {
	_ struct{} `type:"structure"`

	// The name of the queried container.
	Container *Container `type:"structure"`
}

// String returns the string representation
func (s DescribeContainerOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeContainer = "DescribeContainer"

// DescribeContainerRequest returns a request value for making API operation for
// AWS Elemental MediaStore.
//
// Retrieves the properties of the requested container. This request is commonly
// used to retrieve the endpoint of a container. An endpoint is a value assigned
// by the service when a new container is created. A container's endpoint does
// not change after it has been assigned. The DescribeContainer request returns
// a single Container object based on ContainerName. To return all Container
// objects that are associated with a specified AWS account, use ListContainers.
//
//    // Example sending a request using DescribeContainerRequest.
//    req := client.DescribeContainerRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediastore-2017-09-01/DescribeContainer
func (c *Client) DescribeContainerRequest(input *DescribeContainerInput) DescribeContainerRequest {
	op := &aws.Operation{
		Name:       opDescribeContainer,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeContainerInput{}
	}

	req := c.newRequest(op, input, &DescribeContainerOutput{})

	return DescribeContainerRequest{Request: req, Input: input, Copy: c.DescribeContainerRequest}
}

// DescribeContainerRequest is the request type for the
// DescribeContainer API operation.
type DescribeContainerRequest struct {
	*aws.Request
	Input *DescribeContainerInput
	Copy  func(*DescribeContainerInput) DescribeContainerRequest
}

// Send marshals and sends the DescribeContainer API request.
func (r DescribeContainerRequest) Send(ctx context.Context) (*DescribeContainerResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeContainerResponse{
		DescribeContainerOutput: r.Request.Data.(*DescribeContainerOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeContainerResponse is the response type for the
// DescribeContainer API operation.
type DescribeContainerResponse struct {
	*DescribeContainerOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeContainer request.
func (r *DescribeContainerResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}

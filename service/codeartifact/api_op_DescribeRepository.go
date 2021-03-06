// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codeartifact

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type DescribeRepositoryInput struct {
	_ struct{} `type:"structure"`

	// The name of the domain that contains the repository to describe.
	//
	// Domain is a required field
	Domain *string `location:"querystring" locationName:"domain" min:"2" type:"string" required:"true"`

	// The 12-digit account number of the AWS account that owns the domain. It does
	// not include dashes or spaces.
	DomainOwner *string `location:"querystring" locationName:"domain-owner" min:"12" type:"string"`

	// A string that specifies the name of the requested repository.
	//
	// Repository is a required field
	Repository *string `location:"querystring" locationName:"repository" min:"2" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeRepositoryInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeRepositoryInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeRepositoryInput"}

	if s.Domain == nil {
		invalidParams.Add(aws.NewErrParamRequired("Domain"))
	}
	if s.Domain != nil && len(*s.Domain) < 2 {
		invalidParams.Add(aws.NewErrParamMinLen("Domain", 2))
	}
	if s.DomainOwner != nil && len(*s.DomainOwner) < 12 {
		invalidParams.Add(aws.NewErrParamMinLen("DomainOwner", 12))
	}

	if s.Repository == nil {
		invalidParams.Add(aws.NewErrParamRequired("Repository"))
	}
	if s.Repository != nil && len(*s.Repository) < 2 {
		invalidParams.Add(aws.NewErrParamMinLen("Repository", 2))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeRepositoryInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Domain != nil {
		v := *s.Domain

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "domain", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.DomainOwner != nil {
		v := *s.DomainOwner

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "domain-owner", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Repository != nil {
		v := *s.Repository

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "repository", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type DescribeRepositoryOutput struct {
	_ struct{} `type:"structure"`

	// A RepositoryDescription object that contains the requested repository information.
	Repository *RepositoryDescription `locationName:"repository" type:"structure"`
}

// String returns the string representation
func (s DescribeRepositoryOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeRepositoryOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Repository != nil {
		v := s.Repository

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "repository", v, metadata)
	}
	return nil
}

const opDescribeRepository = "DescribeRepository"

// DescribeRepositoryRequest returns a request value for making API operation for
// CodeArtifact.
//
// Returns a RepositoryDescription object that contains detailed information
// about the requested repository.
//
//    // Example sending a request using DescribeRepositoryRequest.
//    req := client.DescribeRepositoryRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codeartifact-2018-09-22/DescribeRepository
func (c *Client) DescribeRepositoryRequest(input *DescribeRepositoryInput) DescribeRepositoryRequest {
	op := &aws.Operation{
		Name:       opDescribeRepository,
		HTTPMethod: "GET",
		HTTPPath:   "/v1/repository",
	}

	if input == nil {
		input = &DescribeRepositoryInput{}
	}

	req := c.newRequest(op, input, &DescribeRepositoryOutput{})

	return DescribeRepositoryRequest{Request: req, Input: input, Copy: c.DescribeRepositoryRequest}
}

// DescribeRepositoryRequest is the request type for the
// DescribeRepository API operation.
type DescribeRepositoryRequest struct {
	*aws.Request
	Input *DescribeRepositoryInput
	Copy  func(*DescribeRepositoryInput) DescribeRepositoryRequest
}

// Send marshals and sends the DescribeRepository API request.
func (r DescribeRepositoryRequest) Send(ctx context.Context) (*DescribeRepositoryResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeRepositoryResponse{
		DescribeRepositoryOutput: r.Request.Data.(*DescribeRepositoryOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeRepositoryResponse is the response type for the
// DescribeRepository API operation.
type DescribeRepositoryResponse struct {
	*DescribeRepositoryOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeRepository request.
func (r *DescribeRepositoryResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}

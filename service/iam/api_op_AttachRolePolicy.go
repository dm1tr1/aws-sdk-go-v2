// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iam

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
)

type AttachRolePolicyInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the IAM policy you want to attach.
	//
	// For more information about ARNs, see Amazon Resource Names (ARNs) and AWS
	// Service Namespaces (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	// in the AWS General Reference.
	//
	// PolicyArn is a required field
	PolicyArn *string `min:"20" type:"string" required:"true"`

	// The name (friendly name, not ARN) of the role to attach the policy to.
	//
	// This parameter allows (through its regex pattern (http://wikipedia.org/wiki/regex))
	// a string of characters consisting of upper and lowercase alphanumeric characters
	// with no spaces. You can also include any of the following characters: _+=,.@-
	//
	// RoleName is a required field
	RoleName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s AttachRolePolicyInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AttachRolePolicyInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AttachRolePolicyInput"}

	if s.PolicyArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("PolicyArn"))
	}
	if s.PolicyArn != nil && len(*s.PolicyArn) < 20 {
		invalidParams.Add(aws.NewErrParamMinLen("PolicyArn", 20))
	}

	if s.RoleName == nil {
		invalidParams.Add(aws.NewErrParamRequired("RoleName"))
	}
	if s.RoleName != nil && len(*s.RoleName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("RoleName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type AttachRolePolicyOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s AttachRolePolicyOutput) String() string {
	return awsutil.Prettify(s)
}

const opAttachRolePolicy = "AttachRolePolicy"

// AttachRolePolicyRequest returns a request value for making API operation for
// AWS Identity and Access Management.
//
// Attaches the specified managed policy to the specified IAM role. When you
// attach a managed policy to a role, the managed policy becomes part of the
// role's permission (access) policy.
//
// You cannot use a managed policy as the role's trust policy. The role's trust
// policy is created at the same time as the role, using CreateRole. You can
// update a role's trust policy using UpdateAssumeRolePolicy.
//
// Use this API to attach a managed policy to a role. To embed an inline policy
// in a role, use PutRolePolicy. For more information about policies, see Managed
// Policies and Inline Policies (https://docs.aws.amazon.com/IAM/latest/UserGuide/policies-managed-vs-inline.html)
// in the IAM User Guide.
//
//    // Example sending a request using AttachRolePolicyRequest.
//    req := client.AttachRolePolicyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/AttachRolePolicy
func (c *Client) AttachRolePolicyRequest(input *AttachRolePolicyInput) AttachRolePolicyRequest {
	op := &aws.Operation{
		Name:       opAttachRolePolicy,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AttachRolePolicyInput{}
	}

	req := c.newRequest(op, input, &AttachRolePolicyOutput{})
	req.Handlers.Unmarshal.Remove(query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	return AttachRolePolicyRequest{Request: req, Input: input, Copy: c.AttachRolePolicyRequest}
}

// AttachRolePolicyRequest is the request type for the
// AttachRolePolicy API operation.
type AttachRolePolicyRequest struct {
	*aws.Request
	Input *AttachRolePolicyInput
	Copy  func(*AttachRolePolicyInput) AttachRolePolicyRequest
}

// Send marshals and sends the AttachRolePolicy API request.
func (r AttachRolePolicyRequest) Send(ctx context.Context) (*AttachRolePolicyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AttachRolePolicyResponse{
		AttachRolePolicyOutput: r.Request.Data.(*AttachRolePolicyOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AttachRolePolicyResponse is the response type for the
// AttachRolePolicy API operation.
type AttachRolePolicyResponse struct {
	*AttachRolePolicyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AttachRolePolicy request.
func (r *AttachRolePolicyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}

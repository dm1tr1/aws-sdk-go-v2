// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iam

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type UploadSigningCertificateInput struct {
	_ struct{} `type:"structure"`

	// The contents of the signing certificate.
	//
	// The regex pattern (http://wikipedia.org/wiki/regex) used to validate this
	// parameter is a string of characters consisting of the following:
	//
	//    * Any printable ASCII character ranging from the space character (\u0020)
	//    through the end of the ASCII character range
	//
	//    * The printable characters in the Basic Latin and Latin-1 Supplement character
	//    set (through \u00FF)
	//
	//    * The special characters tab (\u0009), line feed (\u000A), and carriage
	//    return (\u000D)
	//
	// CertificateBody is a required field
	CertificateBody *string `min:"1" type:"string" required:"true"`

	// The name of the user the signing certificate is for.
	//
	// This parameter allows (through its regex pattern (http://wikipedia.org/wiki/regex))
	// a string of characters consisting of upper and lowercase alphanumeric characters
	// with no spaces. You can also include any of the following characters: _+=,.@-
	UserName *string `min:"1" type:"string"`
}

// String returns the string representation
func (s UploadSigningCertificateInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UploadSigningCertificateInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UploadSigningCertificateInput"}

	if s.CertificateBody == nil {
		invalidParams.Add(aws.NewErrParamRequired("CertificateBody"))
	}
	if s.CertificateBody != nil && len(*s.CertificateBody) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("CertificateBody", 1))
	}
	if s.UserName != nil && len(*s.UserName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("UserName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Contains the response to a successful UploadSigningCertificate request.
type UploadSigningCertificateOutput struct {
	_ struct{} `type:"structure"`

	// Information about the certificate.
	//
	// Certificate is a required field
	Certificate *SigningCertificate `type:"structure" required:"true"`
}

// String returns the string representation
func (s UploadSigningCertificateOutput) String() string {
	return awsutil.Prettify(s)
}

const opUploadSigningCertificate = "UploadSigningCertificate"

// UploadSigningCertificateRequest returns a request value for making API operation for
// AWS Identity and Access Management.
//
// Uploads an X.509 signing certificate and associates it with the specified
// IAM user. Some AWS services use X.509 signing certificates to validate requests
// that are signed with a corresponding private key. When you upload the certificate,
// its default status is Active.
//
// If the UserName is not specified, the IAM user name is determined implicitly
// based on the AWS access key ID used to sign the request. This operation works
// for access keys under the AWS account. Consequently, you can use this operation
// to manage AWS account root user credentials even if the AWS account has no
// associated users.
//
// Because the body of an X.509 certificate can be large, you should use POST
// rather than GET when calling UploadSigningCertificate. For information about
// setting up signatures and authorization through the API, go to Signing AWS
// API Requests (https://docs.aws.amazon.com/general/latest/gr/signing_aws_api_requests.html)
// in the AWS General Reference. For general information about using the Query
// API with IAM, go to Making Query Requests (https://docs.aws.amazon.com/IAM/latest/UserGuide/IAM_UsingQueryAPI.html)
// in the IAM User Guide.
//
//    // Example sending a request using UploadSigningCertificateRequest.
//    req := client.UploadSigningCertificateRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/UploadSigningCertificate
func (c *Client) UploadSigningCertificateRequest(input *UploadSigningCertificateInput) UploadSigningCertificateRequest {
	op := &aws.Operation{
		Name:       opUploadSigningCertificate,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UploadSigningCertificateInput{}
	}

	req := c.newRequest(op, input, &UploadSigningCertificateOutput{})

	return UploadSigningCertificateRequest{Request: req, Input: input, Copy: c.UploadSigningCertificateRequest}
}

// UploadSigningCertificateRequest is the request type for the
// UploadSigningCertificate API operation.
type UploadSigningCertificateRequest struct {
	*aws.Request
	Input *UploadSigningCertificateInput
	Copy  func(*UploadSigningCertificateInput) UploadSigningCertificateRequest
}

// Send marshals and sends the UploadSigningCertificate API request.
func (r UploadSigningCertificateRequest) Send(ctx context.Context) (*UploadSigningCertificateResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UploadSigningCertificateResponse{
		UploadSigningCertificateOutput: r.Request.Data.(*UploadSigningCertificateOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UploadSigningCertificateResponse is the response type for the
// UploadSigningCertificate API operation.
type UploadSigningCertificateResponse struct {
	*UploadSigningCertificateOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UploadSigningCertificate request.
func (r *UploadSigningCertificateResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}

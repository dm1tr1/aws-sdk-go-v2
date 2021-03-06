// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package route53

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// A complex type that contains information about the request to associate a
// VPC with a private hosted zone.
type AssociateVPCWithHostedZoneInput struct {
	_ struct{} `locationName:"AssociateVPCWithHostedZoneRequest" type:"structure" xmlURI:"https://route53.amazonaws.com/doc/2013-04-01/"`

	// Optional: A comment about the association request.
	Comment *string `type:"string"`

	// The ID of the private hosted zone that you want to associate an Amazon VPC
	// with.
	//
	// Note that you can't associate a VPC with a hosted zone that doesn't have
	// an existing VPC association.
	//
	// HostedZoneId is a required field
	HostedZoneId *string `location:"uri" locationName:"Id" type:"string" required:"true"`

	// A complex type that contains information about the VPC that you want to associate
	// with a private hosted zone.
	//
	// VPC is a required field
	VPC *VPC `type:"structure" required:"true"`
}

// String returns the string representation
func (s AssociateVPCWithHostedZoneInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AssociateVPCWithHostedZoneInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AssociateVPCWithHostedZoneInput"}

	if s.HostedZoneId == nil {
		invalidParams.Add(aws.NewErrParamRequired("HostedZoneId"))
	}

	if s.VPC == nil {
		invalidParams.Add(aws.NewErrParamRequired("VPC"))
	}
	if s.VPC != nil {
		if err := s.VPC.Validate(); err != nil {
			invalidParams.AddNested("VPC", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s AssociateVPCWithHostedZoneInput) MarshalFields(e protocol.FieldEncoder) error {

	e.SetFields(protocol.BodyTarget, "AssociateVPCWithHostedZoneRequest", protocol.FieldMarshalerFunc(func(e protocol.FieldEncoder) error {
		if s.Comment != nil {
			v := *s.Comment

			metadata := protocol.Metadata{}
			e.SetValue(protocol.BodyTarget, "Comment", protocol.StringValue(v), metadata)
		}
		if s.VPC != nil {
			v := s.VPC

			metadata := protocol.Metadata{}
			e.SetFields(protocol.BodyTarget, "VPC", v, metadata)
		}
		return nil
	}), protocol.Metadata{XMLNamespaceURI: "https://route53.amazonaws.com/doc/2013-04-01/"})
	if s.HostedZoneId != nil {
		v := *s.HostedZoneId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "Id", protocol.StringValue(v), metadata)
	}
	return nil
}

// A complex type that contains the response information for the AssociateVPCWithHostedZone
// request.
type AssociateVPCWithHostedZoneOutput struct {
	_ struct{} `type:"structure"`

	// A complex type that describes the changes made to your hosted zone.
	//
	// ChangeInfo is a required field
	ChangeInfo *ChangeInfo `type:"structure" required:"true"`
}

// String returns the string representation
func (s AssociateVPCWithHostedZoneOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s AssociateVPCWithHostedZoneOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ChangeInfo != nil {
		v := s.ChangeInfo

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "ChangeInfo", v, metadata)
	}
	return nil
}

const opAssociateVPCWithHostedZone = "AssociateVPCWithHostedZone"

// AssociateVPCWithHostedZoneRequest returns a request value for making API operation for
// Amazon Route 53.
//
// Associates an Amazon VPC with a private hosted zone.
//
// To perform the association, the VPC and the private hosted zone must already
// exist. Also, you can't convert a public hosted zone into a private hosted
// zone.
//
// If you want to associate a VPC that was created by one AWS account with a
// private hosted zone that was created by a different account, do one of the
// following:
//
//    * Use the AWS account that created the private hosted zone to submit a
//    CreateVPCAssociationAuthorization (https://docs.aws.amazon.com/Route53/latest/APIReference/API_CreateVPCAssociationAuthorization.html)
//    request. Then use the account that created the VPC to submit an AssociateVPCWithHostedZone
//    request.
//
//    * If a subnet in the VPC was shared with another account, you can use
//    the account that the subnet was shared with to submit an AssociateVPCWithHostedZone
//    request. For more information about sharing subnets, see Working with
//    Shared VPCs (https://docs.aws.amazon.com/vpc/latest/userguide/vpc-sharing.html).
//
//    // Example sending a request using AssociateVPCWithHostedZoneRequest.
//    req := client.AssociateVPCWithHostedZoneRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/route53-2013-04-01/AssociateVPCWithHostedZone
func (c *Client) AssociateVPCWithHostedZoneRequest(input *AssociateVPCWithHostedZoneInput) AssociateVPCWithHostedZoneRequest {
	op := &aws.Operation{
		Name:       opAssociateVPCWithHostedZone,
		HTTPMethod: "POST",
		HTTPPath:   "/2013-04-01/hostedzone/{Id}/associatevpc",
	}

	if input == nil {
		input = &AssociateVPCWithHostedZoneInput{}
	}

	req := c.newRequest(op, input, &AssociateVPCWithHostedZoneOutput{})

	return AssociateVPCWithHostedZoneRequest{Request: req, Input: input, Copy: c.AssociateVPCWithHostedZoneRequest}
}

// AssociateVPCWithHostedZoneRequest is the request type for the
// AssociateVPCWithHostedZone API operation.
type AssociateVPCWithHostedZoneRequest struct {
	*aws.Request
	Input *AssociateVPCWithHostedZoneInput
	Copy  func(*AssociateVPCWithHostedZoneInput) AssociateVPCWithHostedZoneRequest
}

// Send marshals and sends the AssociateVPCWithHostedZone API request.
func (r AssociateVPCWithHostedZoneRequest) Send(ctx context.Context) (*AssociateVPCWithHostedZoneResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AssociateVPCWithHostedZoneResponse{
		AssociateVPCWithHostedZoneOutput: r.Request.Data.(*AssociateVPCWithHostedZoneOutput),
		response:                         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AssociateVPCWithHostedZoneResponse is the response type for the
// AssociateVPCWithHostedZone API operation.
type AssociateVPCWithHostedZoneResponse struct {
	*AssociateVPCWithHostedZoneOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AssociateVPCWithHostedZone request.
func (r *AssociateVPCWithHostedZoneResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}

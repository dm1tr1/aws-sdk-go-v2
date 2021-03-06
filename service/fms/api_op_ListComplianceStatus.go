// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package fms

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListComplianceStatusInput struct {
	_ struct{} `type:"structure"`

	// Specifies the number of PolicyComplianceStatus objects that you want AWS
	// Firewall Manager to return for this request. If you have more PolicyComplianceStatus
	// objects than the number that you specify for MaxResults, the response includes
	// a NextToken value that you can use to get another batch of PolicyComplianceStatus
	// objects.
	MaxResults *int64 `min:"1" type:"integer"`

	// If you specify a value for MaxResults and you have more PolicyComplianceStatus
	// objects than the number that you specify for MaxResults, AWS Firewall Manager
	// returns a NextToken value in the response that allows you to list another
	// group of PolicyComplianceStatus objects. For the second and subsequent ListComplianceStatus
	// requests, specify the value of NextToken from the previous response to get
	// information about another batch of PolicyComplianceStatus objects.
	NextToken *string `min:"1" type:"string"`

	// The ID of the AWS Firewall Manager policy that you want the details for.
	//
	// PolicyId is a required field
	PolicyId *string `min:"36" type:"string" required:"true"`
}

// String returns the string representation
func (s ListComplianceStatusInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListComplianceStatusInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListComplianceStatusInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 1))
	}

	if s.PolicyId == nil {
		invalidParams.Add(aws.NewErrParamRequired("PolicyId"))
	}
	if s.PolicyId != nil && len(*s.PolicyId) < 36 {
		invalidParams.Add(aws.NewErrParamMinLen("PolicyId", 36))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListComplianceStatusOutput struct {
	_ struct{} `type:"structure"`

	// If you have more PolicyComplianceStatus objects than the number that you
	// specified for MaxResults in the request, the response includes a NextToken
	// value. To list more PolicyComplianceStatus objects, submit another ListComplianceStatus
	// request, and specify the NextToken value from the response in the NextToken
	// value in the next request.
	NextToken *string `min:"1" type:"string"`

	// An array of PolicyComplianceStatus objects.
	PolicyComplianceStatusList []PolicyComplianceStatus `type:"list"`
}

// String returns the string representation
func (s ListComplianceStatusOutput) String() string {
	return awsutil.Prettify(s)
}

const opListComplianceStatus = "ListComplianceStatus"

// ListComplianceStatusRequest returns a request value for making API operation for
// Firewall Management Service.
//
// Returns an array of PolicyComplianceStatus objects. Use PolicyComplianceStatus
// to get a summary of which member accounts are protected by the specified
// policy.
//
//    // Example sending a request using ListComplianceStatusRequest.
//    req := client.ListComplianceStatusRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/fms-2018-01-01/ListComplianceStatus
func (c *Client) ListComplianceStatusRequest(input *ListComplianceStatusInput) ListComplianceStatusRequest {
	op := &aws.Operation{
		Name:       opListComplianceStatus,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListComplianceStatusInput{}
	}

	req := c.newRequest(op, input, &ListComplianceStatusOutput{})

	return ListComplianceStatusRequest{Request: req, Input: input, Copy: c.ListComplianceStatusRequest}
}

// ListComplianceStatusRequest is the request type for the
// ListComplianceStatus API operation.
type ListComplianceStatusRequest struct {
	*aws.Request
	Input *ListComplianceStatusInput
	Copy  func(*ListComplianceStatusInput) ListComplianceStatusRequest
}

// Send marshals and sends the ListComplianceStatus API request.
func (r ListComplianceStatusRequest) Send(ctx context.Context) (*ListComplianceStatusResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListComplianceStatusResponse{
		ListComplianceStatusOutput: r.Request.Data.(*ListComplianceStatusOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListComplianceStatusRequestPaginator returns a paginator for ListComplianceStatus.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListComplianceStatusRequest(input)
//   p := fms.NewListComplianceStatusRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListComplianceStatusPaginator(req ListComplianceStatusRequest) ListComplianceStatusPaginator {
	return ListComplianceStatusPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListComplianceStatusInput
				if req.Input != nil {
					tmp := *req.Input
					inCpy = &tmp
				}

				newReq := req.Copy(inCpy)
				newReq.SetContext(ctx)
				return newReq.Request, nil
			},
		},
	}
}

// ListComplianceStatusPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListComplianceStatusPaginator struct {
	aws.Pager
}

func (p *ListComplianceStatusPaginator) CurrentPage() *ListComplianceStatusOutput {
	return p.Pager.CurrentPage().(*ListComplianceStatusOutput)
}

// ListComplianceStatusResponse is the response type for the
// ListComplianceStatus API operation.
type ListComplianceStatusResponse struct {
	*ListComplianceStatusOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListComplianceStatus request.
func (r *ListComplianceStatusResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}

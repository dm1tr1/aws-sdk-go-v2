// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package kinesisanalytics

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type AddApplicationOutputInput struct {
	_ struct{} `type:"structure"`

	// Name of the application to which you want to add the output configuration.
	//
	// ApplicationName is a required field
	ApplicationName *string `min:"1" type:"string" required:"true"`

	// Version of the application to which you want to add the output configuration.
	// You can use the DescribeApplication (https://docs.aws.amazon.com/kinesisanalytics/latest/dev/API_DescribeApplication.html)
	// operation to get the current application version. If the version specified
	// is not the current version, the ConcurrentModificationException is returned.
	//
	// CurrentApplicationVersionId is a required field
	CurrentApplicationVersionId *int64 `min:"1" type:"long" required:"true"`

	// An array of objects, each describing one output configuration. In the output
	// configuration, you specify the name of an in-application stream, a destination
	// (that is, an Amazon Kinesis stream, an Amazon Kinesis Firehose delivery stream,
	// or an AWS Lambda function), and record the formation to use when writing
	// to the destination.
	//
	// Output is a required field
	Output *Output `type:"structure" required:"true"`
}

// String returns the string representation
func (s AddApplicationOutputInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AddApplicationOutputInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AddApplicationOutputInput"}

	if s.ApplicationName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApplicationName"))
	}
	if s.ApplicationName != nil && len(*s.ApplicationName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ApplicationName", 1))
	}

	if s.CurrentApplicationVersionId == nil {
		invalidParams.Add(aws.NewErrParamRequired("CurrentApplicationVersionId"))
	}
	if s.CurrentApplicationVersionId != nil && *s.CurrentApplicationVersionId < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("CurrentApplicationVersionId", 1))
	}

	if s.Output == nil {
		invalidParams.Add(aws.NewErrParamRequired("Output"))
	}
	if s.Output != nil {
		if err := s.Output.Validate(); err != nil {
			invalidParams.AddNested("Output", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type AddApplicationOutputOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s AddApplicationOutputOutput) String() string {
	return awsutil.Prettify(s)
}

const opAddApplicationOutput = "AddApplicationOutput"

// AddApplicationOutputRequest returns a request value for making API operation for
// Amazon Kinesis Analytics.
//
//
// This documentation is for version 1 of the Amazon Kinesis Data Analytics
// API, which only supports SQL applications. Version 2 of the API supports
// SQL and Java applications. For more information about version 2, see Amazon
// Kinesis Data Analytics API V2 Documentation (/kinesisanalytics/latest/apiv2/Welcome.html).
//
// Adds an external destination to your Amazon Kinesis Analytics application.
//
// If you want Amazon Kinesis Analytics to deliver data from an in-application
// stream within your application to an external destination (such as an Amazon
// Kinesis stream, an Amazon Kinesis Firehose delivery stream, or an AWS Lambda
// function), you add the relevant configuration to your application using this
// operation. You can configure one or more outputs for your application. Each
// output configuration maps an in-application stream and an external destination.
//
// You can use one of the output configurations to deliver data from your in-application
// error stream to an external destination so that you can analyze the errors.
// For more information, see Understanding Application Output (Destination)
// (https://docs.aws.amazon.com/kinesisanalytics/latest/dev/how-it-works-output.html).
//
// Any configuration update, including adding a streaming source using this
// operation, results in a new version of the application. You can use the DescribeApplication
// (https://docs.aws.amazon.com/kinesisanalytics/latest/dev/API_DescribeApplication.html)
// operation to find the current application version.
//
// For the limits on the number of application inputs and outputs you can configure,
// see Limits (https://docs.aws.amazon.com/kinesisanalytics/latest/dev/limits.html).
//
// This operation requires permissions to perform the kinesisanalytics:AddApplicationOutput
// action.
//
//    // Example sending a request using AddApplicationOutputRequest.
//    req := client.AddApplicationOutputRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/kinesisanalytics-2015-08-14/AddApplicationOutput
func (c *Client) AddApplicationOutputRequest(input *AddApplicationOutputInput) AddApplicationOutputRequest {
	op := &aws.Operation{
		Name:       opAddApplicationOutput,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AddApplicationOutputInput{}
	}

	req := c.newRequest(op, input, &AddApplicationOutputOutput{})

	return AddApplicationOutputRequest{Request: req, Input: input, Copy: c.AddApplicationOutputRequest}
}

// AddApplicationOutputRequest is the request type for the
// AddApplicationOutput API operation.
type AddApplicationOutputRequest struct {
	*aws.Request
	Input *AddApplicationOutputInput
	Copy  func(*AddApplicationOutputInput) AddApplicationOutputRequest
}

// Send marshals and sends the AddApplicationOutput API request.
func (r AddApplicationOutputRequest) Send(ctx context.Context) (*AddApplicationOutputResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AddApplicationOutputResponse{
		AddApplicationOutputOutput: r.Request.Data.(*AddApplicationOutputOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AddApplicationOutputResponse is the response type for the
// AddApplicationOutput API operation.
type AddApplicationOutputResponse struct {
	*AddApplicationOutputOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AddApplicationOutput request.
func (r *AddApplicationOutputResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}

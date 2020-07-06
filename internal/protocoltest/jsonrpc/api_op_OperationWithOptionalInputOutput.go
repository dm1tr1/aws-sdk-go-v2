// Code generated by smithy-go-codegen DO NOT EDIT.
package jsonrpc

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

func (c *Client) OperationWithOptionalInputOutput(ctx context.Context, params *OperationWithOptionalInputOutputInput, optFns ...func(*Options)) (*OperationWithOptionalInputOutputOutput, error) {
	stack := middleware.NewStack("OperationWithOptionalInputOutput", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	awsmiddleware.AddResolveServiceEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	retry.AddRetryMiddlewares(stack, options)
	v4.AddHTTPSignerMiddleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opOperationWithOptionalInputOutput(options.Region), middleware.Before)
	addawsAwsjson11_serdeOpOperationWithOptionalInputOutputMiddlewares(stack)

	for _, fn := range options.APIOptions {
		if err := fn(stack); err != nil {
			return nil, err
		}
	}
	handler := middleware.DecorateHandler(smithyhttp.NewClientHandler(options.HTTPClient), stack)
	result, metadata, err := handler.Handle(ctx, params)
	if err != nil {
		return nil, &smithy.OperationError{
			ServiceID:     c.ServiceID(),
			OperationName: "OperationWithOptionalInputOutput",
			Err:           err,
		}
	}
	out := result.(*OperationWithOptionalInputOutputOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type OperationWithOptionalInputOutputInput struct {
	Value *string
}

type OperationWithOptionalInputOutputOutput struct {
	Value *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpOperationWithOptionalInputOutputMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpOperationWithOptionalInputOutput{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpOperationWithOptionalInputOutput{}, middleware.After)
}

func newServiceMetadataMiddleware_opOperationWithOptionalInputOutput(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:         region,
		ServiceName:    "Json Protocol",
		ServiceID:      "jsonprotocol",
		EndpointPrefix: "jsonprotocol",
		SigningName:    "foo",
		OperationName:  "OperationWithOptionalInputOutput",
	}
}
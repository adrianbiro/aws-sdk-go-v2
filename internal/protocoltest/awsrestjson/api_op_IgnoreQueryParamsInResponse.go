// Code generated by smithy-go-codegen DO NOT EDIT.

package awsrestjson

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// This example ensures that query string bound request parameters are serialized
// in the body of responses if the structure is used in both the request and
// response.
func (c *Client) IgnoreQueryParamsInResponse(ctx context.Context, params *IgnoreQueryParamsInResponseInput, optFns ...func(*Options)) (*IgnoreQueryParamsInResponseOutput, error) {
	stack := middleware.NewStack("IgnoreQueryParamsInResponse", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpIgnoreQueryParamsInResponseMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	retry.AddRetryMiddlewares(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opIgnoreQueryParamsInResponse(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)

	for _, fn := range options.APIOptions {
		if err := fn(stack); err != nil {
			return nil, err
		}
	}
	handler := middleware.DecorateHandler(smithyhttp.NewClientHandler(options.HTTPClient), stack)
	result, metadata, err := handler.Handle(ctx, params)
	if err != nil {
		return nil, &smithy.OperationError{
			ServiceID:     ServiceID,
			OperationName: "IgnoreQueryParamsInResponse",
			Err:           err,
		}
	}
	out := result.(*IgnoreQueryParamsInResponseOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type IgnoreQueryParamsInResponseInput struct {
}

type IgnoreQueryParamsInResponseOutput struct {
	Baz *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpIgnoreQueryParamsInResponseMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpIgnoreQueryParamsInResponse{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpIgnoreQueryParamsInResponse{}, middleware.After)
}

func newServiceMetadataMiddleware_opIgnoreQueryParamsInResponse(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "IgnoreQueryParamsInResponse",
	}
}
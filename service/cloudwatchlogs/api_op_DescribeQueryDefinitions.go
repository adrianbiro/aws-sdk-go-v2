// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudwatchlogs

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

func (c *Client) DescribeQueryDefinitions(ctx context.Context, params *DescribeQueryDefinitionsInput, optFns ...func(*Options)) (*DescribeQueryDefinitionsOutput, error) {
	stack := middleware.NewStack("DescribeQueryDefinitions", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribeQueryDefinitionsMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	retry.AddRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeQueryDefinitions(options.Region), middleware.Before)
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
			OperationName: "DescribeQueryDefinitions",
			Err:           err,
		}
	}
	out := result.(*DescribeQueryDefinitionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeQueryDefinitionsInput struct {
	QueryDefinitionNamePrefix *string
	MaxResults                *int32
	// The token for the next set of items to return. The token expires after 24 hours.
	NextToken *string
}

type DescribeQueryDefinitionsOutput struct {
	// The token for the next set of items to return. The token expires after 24 hours.
	NextToken        *string
	QueryDefinitions []*types.QueryDefinition

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribeQueryDefinitionsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeQueryDefinitions{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeQueryDefinitions{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeQueryDefinitions(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "logs",
		OperationName: "DescribeQueryDefinitions",
	}
}
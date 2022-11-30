// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Describes the app.
func (c *Client) DescribeApp(ctx context.Context, params *DescribeAppInput, optFns ...func(*Options)) (*DescribeAppOutput, error) {
	if params == nil {
		params = &DescribeAppInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeApp", params, optFns, c.addOperationDescribeAppMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeAppOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeAppInput struct {

	// The name of the app.
	//
	// This member is required.
	AppName *string

	// The type of app.
	//
	// This member is required.
	AppType types.AppType

	// The domain ID.
	//
	// This member is required.
	DomainId *string

	// The name of the space.
	SpaceName *string

	// The user profile name.
	UserProfileName *string

	noSmithyDocumentSerde
}

type DescribeAppOutput struct {

	// The Amazon Resource Name (ARN) of the app.
	AppArn *string

	// The name of the app.
	AppName *string

	// The type of app.
	AppType types.AppType

	// The creation time.
	CreationTime *time.Time

	// The domain ID.
	DomainId *string

	// The failure reason.
	FailureReason *string

	// The timestamp of the last health check.
	LastHealthCheckTimestamp *time.Time

	// The timestamp of the last user's activity. LastUserActivityTimestamp is also
	// updated when SageMaker performs health checks without user activity. As a
	// result, this value is set to the same value as LastHealthCheckTimestamp.
	LastUserActivityTimestamp *time.Time

	// The instance type and the Amazon Resource Name (ARN) of the SageMaker image
	// created on the instance.
	ResourceSpec *types.ResourceSpec

	// The name of the space.
	SpaceName *string

	// The status.
	Status types.AppStatus

	// The user profile name.
	UserProfileName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeAppMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeApp{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeApp{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpDescribeAppValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeApp(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opDescribeApp(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "DescribeApp",
	}
}

// Code generated by smithy-go-codegen DO NOT EDIT.

package securityhub

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/securityhub/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This method is deprecated. Instead, use GetAdministratorAccount. The Security
// Hub console continues to use GetMasterAccount. It will eventually change to use
// GetAdministratorAccount. Any IAM policies that specifically control access to
// this function must continue to use GetMasterAccount. You should also add
// GetAdministratorAccount to your policies to ensure that the correct permissions
// are in place after the console begins to use GetAdministratorAccount. Provides
// the details for the Security Hub administrator account for the current member
// account. Can be used by both member accounts that are managed using
// Organizations and accounts that were invited manually.
//
// Deprecated: This API has been deprecated, use GetAdministratorAccount API
// instead.
func (c *Client) GetMasterAccount(ctx context.Context, params *GetMasterAccountInput, optFns ...func(*Options)) (*GetMasterAccountOutput, error) {
	if params == nil {
		params = &GetMasterAccountInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetMasterAccount", params, optFns, addOperationGetMasterAccountMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetMasterAccountOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetMasterAccountInput struct {
}

type GetMasterAccountOutput struct {

	// A list of details about the Security Hub administrator account for the current
	// member account.
	Master *types.Invitation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetMasterAccountMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetMasterAccount{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetMasterAccount{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetMasterAccount(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetMasterAccount(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "securityhub",
		OperationName: "GetMasterAccount",
	}
}

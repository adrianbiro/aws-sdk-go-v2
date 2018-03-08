// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package translate

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
)

// Translate provides the API operation methods for making requests to
// Amazon Translate. See this package's package overview docs
// for details on the service.
//
// Translate methods are safe to use concurrently. It is not safe to
// modify mutate any of the struct's properties though.
type Translate struct {
	*aws.Client
}

// Used for custom client initialization logic
var initClient func(*Translate)

// Used for custom request initialization logic
var initRequest func(*Translate, *aws.Request)

// Service information constants
const (
	ServiceName = "translate" // Service endpoint prefix API calls made to.
	EndpointsID = ServiceName // Service ID for Regions and Endpoints metadata.
)

// New creates a new instance of the Translate client with a config.
//
// Example:
//     // Create a Translate client from just a config.
//     svc := translate.New(myConfig)
func New(config aws.Config) *Translate {
	var signingName string
	signingName = "translate"
	signingRegion := config.Region

	svc := &Translate{
		Client: aws.NewClient(
			config,
			aws.Metadata{
				ServiceName:   ServiceName,
				SigningName:   signingName,
				SigningRegion: signingRegion,
				APIVersion:    "2017-07-01",
				JSONVersion:   "1.1",
				TargetPrefix:  "AWSShineFrontendService_20170701",
			},
		),
	}

	// Handlers
	svc.Handlers.Sign.PushBackNamed(v4.SignRequestHandler)
	svc.Handlers.Build.PushBackNamed(jsonrpc.BuildHandler)
	svc.Handlers.Unmarshal.PushBackNamed(jsonrpc.UnmarshalHandler)
	svc.Handlers.UnmarshalMeta.PushBackNamed(jsonrpc.UnmarshalMetaHandler)
	svc.Handlers.UnmarshalError.PushBackNamed(jsonrpc.UnmarshalErrorHandler)

	// Run custom client initialization if present
	if initClient != nil {
		initClient(svc)
	}

	return svc
}

// newRequest creates a new request for a Translate operation and runs any
// custom request initialization.
func (c *Translate) newRequest(op *aws.Operation, params, data interface{}) *aws.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(c, req)
	}

	return req
}
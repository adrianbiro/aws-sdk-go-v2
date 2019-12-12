// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package chime

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type ListAttendeesInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of results to return in a single call.
	MaxResults *int64 `location:"querystring" locationName:"max-results" min:"1" type:"integer"`

	// The Amazon Chime SDK meeting ID.
	//
	// MeetingId is a required field
	MeetingId *string `location:"uri" locationName:"meetingId" type:"string" required:"true"`

	// The token to use to retrieve the next page of results.
	NextToken *string `location:"querystring" locationName:"next-token" type:"string"`
}

// String returns the string representation
func (s ListAttendeesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListAttendeesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListAttendeesInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if s.MeetingId == nil {
		invalidParams.Add(aws.NewErrParamRequired("MeetingId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListAttendeesInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.MeetingId != nil {
		v := *s.MeetingId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "meetingId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.MaxResults != nil {
		v := *s.MaxResults

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "max-results", protocol.Int64Value(v), metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "next-token", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type ListAttendeesOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Chime SDK attendee information.
	Attendees []Attendee `type:"list"`

	// The token to use to retrieve the next page of results.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListAttendeesOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListAttendeesOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Attendees != nil {
		v := s.Attendees

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "Attendees", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "NextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opListAttendees = "ListAttendees"

// ListAttendeesRequest returns a request value for making API operation for
// Amazon Chime.
//
// Lists the attendees for the specified Amazon Chime SDK meeting. For more
// information about the Amazon Chime SDK, see Using the Amazon Chime SDK (https://docs.aws.amazon.com/chime/latest/dg/meetings-sdk.html)
// in the Amazon Chime Developer Guide.
//
//    // Example sending a request using ListAttendeesRequest.
//    req := client.ListAttendeesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/chime-2018-05-01/ListAttendees
func (c *Client) ListAttendeesRequest(input *ListAttendeesInput) ListAttendeesRequest {
	op := &aws.Operation{
		Name:       opListAttendees,
		HTTPMethod: "GET",
		HTTPPath:   "/meetings/{meetingId}/attendees",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListAttendeesInput{}
	}

	req := c.newRequest(op, input, &ListAttendeesOutput{})
	return ListAttendeesRequest{Request: req, Input: input, Copy: c.ListAttendeesRequest}
}

// ListAttendeesRequest is the request type for the
// ListAttendees API operation.
type ListAttendeesRequest struct {
	*aws.Request
	Input *ListAttendeesInput
	Copy  func(*ListAttendeesInput) ListAttendeesRequest
}

// Send marshals and sends the ListAttendees API request.
func (r ListAttendeesRequest) Send(ctx context.Context) (*ListAttendeesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListAttendeesResponse{
		ListAttendeesOutput: r.Request.Data.(*ListAttendeesOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListAttendeesRequestPaginator returns a paginator for ListAttendees.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListAttendeesRequest(input)
//   p := chime.NewListAttendeesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListAttendeesPaginator(req ListAttendeesRequest) ListAttendeesPaginator {
	return ListAttendeesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListAttendeesInput
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

// ListAttendeesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListAttendeesPaginator struct {
	aws.Pager
}

func (p *ListAttendeesPaginator) CurrentPage() *ListAttendeesOutput {
	return p.Pager.CurrentPage().(*ListAttendeesOutput)
}

// ListAttendeesResponse is the response type for the
// ListAttendees API operation.
type ListAttendeesResponse struct {
	*ListAttendeesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListAttendees request.
func (r *ListAttendeesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
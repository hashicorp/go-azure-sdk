package eventexceptionoccurrence

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ForwardEventExceptionOccurrenceOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type ForwardEventExceptionOccurrenceOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultForwardEventExceptionOccurrenceOperationOptions() ForwardEventExceptionOccurrenceOperationOptions {
	return ForwardEventExceptionOccurrenceOperationOptions{}
}

func (o ForwardEventExceptionOccurrenceOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ForwardEventExceptionOccurrenceOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o ForwardEventExceptionOccurrenceOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// ForwardEventExceptionOccurrence - Invoke action forward. This action allows the organizer or attendee of a meeting
// event to forward the meeting request to a new recipient. If the meeting event is forwarded from an attendee's
// Microsoft 365 mailbox to another recipient, this action also sends a message to notify the organizer of the
// forwarding, and adds the recipient to the organizer's copy of the meeting event. This convenience is not available
// when forwarding from an Outlook.com account.
func (c EventExceptionOccurrenceClient) ForwardEventExceptionOccurrence(ctx context.Context, id beta.UserIdEventIdExceptionOccurrenceId, input ForwardEventExceptionOccurrenceRequest, options ForwardEventExceptionOccurrenceOperationOptions) (result ForwardEventExceptionOccurrenceOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/forward", id.ID()),
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	if err = req.Marshal(input); err != nil {
		return
	}

	var resp *client.Response
	resp, err = req.Execute(ctx)
	if resp != nil {
		result.OData = resp.OData
		result.HttpResponse = resp.Response
	}
	if err != nil {
		return
	}

	return
}

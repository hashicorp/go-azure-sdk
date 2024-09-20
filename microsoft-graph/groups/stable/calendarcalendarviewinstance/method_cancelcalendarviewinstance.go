package calendarcalendarviewinstance

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CancelCalendarViewInstanceOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type CancelCalendarViewInstanceOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultCancelCalendarViewInstanceOperationOptions() CancelCalendarViewInstanceOperationOptions {
	return CancelCalendarViewInstanceOperationOptions{}
}

func (o CancelCalendarViewInstanceOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CancelCalendarViewInstanceOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CancelCalendarViewInstanceOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CancelCalendarViewInstance - Invoke action cancel. This action allows the organizer of a meeting to send a
// cancellation message and cancel the event. The action moves the event to the Deleted Items folder. The organizer can
// also cancel an occurrence of a recurring meeting by providing the occurrence event ID. An attendee calling this
// action gets an error (HTTP 400 Bad Request), with the following error message: 'Your request can't be completed. You
// need to be an organizer to cancel a meeting.' This action differs from Delete in that Cancel is available to only the
// organizer, and lets the organizer send a custom message to the attendees about the cancellation.
func (c CalendarCalendarViewInstanceClient) CancelCalendarViewInstance(ctx context.Context, id stable.GroupIdCalendarCalendarViewIdInstanceId, input CancelCalendarViewInstanceRequest, options CancelCalendarViewInstanceOperationOptions) (result CancelCalendarViewInstanceOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/cancel", id.ID()),
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

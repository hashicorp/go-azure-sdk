package calendareventinstanceexceptionoccurrence

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

type SnoozeCalendarEventInstanceExceptionOccurrenceReminderOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type SnoozeCalendarEventInstanceExceptionOccurrenceReminderOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultSnoozeCalendarEventInstanceExceptionOccurrenceReminderOperationOptions() SnoozeCalendarEventInstanceExceptionOccurrenceReminderOperationOptions {
	return SnoozeCalendarEventInstanceExceptionOccurrenceReminderOperationOptions{}
}

func (o SnoozeCalendarEventInstanceExceptionOccurrenceReminderOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o SnoozeCalendarEventInstanceExceptionOccurrenceReminderOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o SnoozeCalendarEventInstanceExceptionOccurrenceReminderOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// SnoozeCalendarEventInstanceExceptionOccurrenceReminder - Invoke action snoozeReminder. Postpone a reminder for an
// event in a user calendar until a new time.
func (c CalendarEventInstanceExceptionOccurrenceClient) SnoozeCalendarEventInstanceExceptionOccurrenceReminder(ctx context.Context, id beta.GroupIdCalendarEventIdInstanceIdExceptionOccurrenceId, input SnoozeCalendarEventInstanceExceptionOccurrenceReminderRequest, options SnoozeCalendarEventInstanceExceptionOccurrenceReminderOperationOptions) (result SnoozeCalendarEventInstanceExceptionOccurrenceReminderOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusCreated,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/snoozeReminder", id.ID()),
		RetryFunc:     options.RetryFunc,
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

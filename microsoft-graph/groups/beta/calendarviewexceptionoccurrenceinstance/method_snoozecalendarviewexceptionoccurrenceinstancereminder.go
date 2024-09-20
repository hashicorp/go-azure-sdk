package calendarviewexceptionoccurrenceinstance

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

type SnoozeCalendarViewExceptionOccurrenceInstanceReminderOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type SnoozeCalendarViewExceptionOccurrenceInstanceReminderOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultSnoozeCalendarViewExceptionOccurrenceInstanceReminderOperationOptions() SnoozeCalendarViewExceptionOccurrenceInstanceReminderOperationOptions {
	return SnoozeCalendarViewExceptionOccurrenceInstanceReminderOperationOptions{}
}

func (o SnoozeCalendarViewExceptionOccurrenceInstanceReminderOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o SnoozeCalendarViewExceptionOccurrenceInstanceReminderOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o SnoozeCalendarViewExceptionOccurrenceInstanceReminderOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// SnoozeCalendarViewExceptionOccurrenceInstanceReminder - Invoke action snoozeReminder. Postpone a reminder for an
// event in a user calendar until a new time.
func (c CalendarViewExceptionOccurrenceInstanceClient) SnoozeCalendarViewExceptionOccurrenceInstanceReminder(ctx context.Context, id beta.GroupIdCalendarViewIdExceptionOccurrenceIdInstanceId, input SnoozeCalendarViewExceptionOccurrenceInstanceReminderRequest, options SnoozeCalendarViewExceptionOccurrenceInstanceReminderOperationOptions) (result SnoozeCalendarViewExceptionOccurrenceInstanceReminderOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/snoozeReminder", id.ID()),
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

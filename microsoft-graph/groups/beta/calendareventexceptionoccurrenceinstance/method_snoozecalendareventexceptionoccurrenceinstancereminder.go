package calendareventexceptionoccurrenceinstance

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

type SnoozeCalendarEventExceptionOccurrenceInstanceReminderOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type SnoozeCalendarEventExceptionOccurrenceInstanceReminderOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultSnoozeCalendarEventExceptionOccurrenceInstanceReminderOperationOptions() SnoozeCalendarEventExceptionOccurrenceInstanceReminderOperationOptions {
	return SnoozeCalendarEventExceptionOccurrenceInstanceReminderOperationOptions{}
}

func (o SnoozeCalendarEventExceptionOccurrenceInstanceReminderOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o SnoozeCalendarEventExceptionOccurrenceInstanceReminderOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o SnoozeCalendarEventExceptionOccurrenceInstanceReminderOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// SnoozeCalendarEventExceptionOccurrenceInstanceReminder - Invoke action snoozeReminder. Postpone a reminder for an
// event in a user calendar until a new time.
func (c CalendarEventExceptionOccurrenceInstanceClient) SnoozeCalendarEventExceptionOccurrenceInstanceReminder(ctx context.Context, id beta.GroupIdCalendarEventIdExceptionOccurrenceIdInstanceId, input SnoozeCalendarEventExceptionOccurrenceInstanceReminderRequest, options SnoozeCalendarEventExceptionOccurrenceInstanceReminderOperationOptions) (result SnoozeCalendarEventExceptionOccurrenceInstanceReminderOperationResponse, err error) {
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

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

type DismissCalendarEventExceptionOccurrenceInstanceReminderOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type DismissCalendarEventExceptionOccurrenceInstanceReminderOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultDismissCalendarEventExceptionOccurrenceInstanceReminderOperationOptions() DismissCalendarEventExceptionOccurrenceInstanceReminderOperationOptions {
	return DismissCalendarEventExceptionOccurrenceInstanceReminderOperationOptions{}
}

func (o DismissCalendarEventExceptionOccurrenceInstanceReminderOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o DismissCalendarEventExceptionOccurrenceInstanceReminderOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o DismissCalendarEventExceptionOccurrenceInstanceReminderOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// DismissCalendarEventExceptionOccurrenceInstanceReminder - Invoke action dismissReminder. Dismiss a reminder that has
// been triggered for an event in a user calendar.
func (c CalendarEventExceptionOccurrenceInstanceClient) DismissCalendarEventExceptionOccurrenceInstanceReminder(ctx context.Context, id beta.GroupIdCalendarEventIdExceptionOccurrenceIdInstanceId, options DismissCalendarEventExceptionOccurrenceInstanceReminderOperationOptions) (result DismissCalendarEventExceptionOccurrenceInstanceReminderOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/dismissReminder", id.ID()),
		RetryFunc:     options.RetryFunc,
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
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

package calendareventinstance

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

type SnoozeCalendarEventInstanceReminderOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type SnoozeCalendarEventInstanceReminderOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultSnoozeCalendarEventInstanceReminderOperationOptions() SnoozeCalendarEventInstanceReminderOperationOptions {
	return SnoozeCalendarEventInstanceReminderOperationOptions{}
}

func (o SnoozeCalendarEventInstanceReminderOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o SnoozeCalendarEventInstanceReminderOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o SnoozeCalendarEventInstanceReminderOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// SnoozeCalendarEventInstanceReminder - Invoke action snoozeReminder. Postpone a reminder for an event in a user
// calendar until a new time.
func (c CalendarEventInstanceClient) SnoozeCalendarEventInstanceReminder(ctx context.Context, id stable.GroupIdCalendarEventIdInstanceId, input SnoozeCalendarEventInstanceReminderRequest, options SnoozeCalendarEventInstanceReminderOperationOptions) (result SnoozeCalendarEventInstanceReminderOperationResponse, err error) {
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

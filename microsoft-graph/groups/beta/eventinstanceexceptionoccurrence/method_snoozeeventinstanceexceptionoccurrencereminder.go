package eventinstanceexceptionoccurrence

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

type SnoozeEventInstanceExceptionOccurrenceReminderOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type SnoozeEventInstanceExceptionOccurrenceReminderOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultSnoozeEventInstanceExceptionOccurrenceReminderOperationOptions() SnoozeEventInstanceExceptionOccurrenceReminderOperationOptions {
	return SnoozeEventInstanceExceptionOccurrenceReminderOperationOptions{}
}

func (o SnoozeEventInstanceExceptionOccurrenceReminderOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o SnoozeEventInstanceExceptionOccurrenceReminderOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o SnoozeEventInstanceExceptionOccurrenceReminderOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// SnoozeEventInstanceExceptionOccurrenceReminder - Invoke action snoozeReminder. Postpone a reminder for an event in a
// user calendar until a new time.
func (c EventInstanceExceptionOccurrenceClient) SnoozeEventInstanceExceptionOccurrenceReminder(ctx context.Context, id beta.GroupIdEventIdInstanceIdExceptionOccurrenceId, input SnoozeEventInstanceExceptionOccurrenceReminderRequest, options SnoozeEventInstanceExceptionOccurrenceReminderOperationOptions) (result SnoozeEventInstanceExceptionOccurrenceReminderOperationResponse, err error) {
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

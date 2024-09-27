package eventexceptionoccurrenceinstance

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

type SnoozeEventExceptionOccurrenceInstanceReminderOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type SnoozeEventExceptionOccurrenceInstanceReminderOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultSnoozeEventExceptionOccurrenceInstanceReminderOperationOptions() SnoozeEventExceptionOccurrenceInstanceReminderOperationOptions {
	return SnoozeEventExceptionOccurrenceInstanceReminderOperationOptions{}
}

func (o SnoozeEventExceptionOccurrenceInstanceReminderOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o SnoozeEventExceptionOccurrenceInstanceReminderOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o SnoozeEventExceptionOccurrenceInstanceReminderOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// SnoozeEventExceptionOccurrenceInstanceReminder - Invoke action snoozeReminder. Postpone a reminder for an event in a
// user calendar until a new time.
func (c EventExceptionOccurrenceInstanceClient) SnoozeEventExceptionOccurrenceInstanceReminder(ctx context.Context, id beta.UserIdEventIdExceptionOccurrenceIdInstanceId, input SnoozeEventExceptionOccurrenceInstanceReminderRequest, options SnoozeEventExceptionOccurrenceInstanceReminderOperationOptions) (result SnoozeEventExceptionOccurrenceInstanceReminderOperationResponse, err error) {
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

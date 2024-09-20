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

type DismissEventInstanceExceptionOccurrenceReminderOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type DismissEventInstanceExceptionOccurrenceReminderOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultDismissEventInstanceExceptionOccurrenceReminderOperationOptions() DismissEventInstanceExceptionOccurrenceReminderOperationOptions {
	return DismissEventInstanceExceptionOccurrenceReminderOperationOptions{}
}

func (o DismissEventInstanceExceptionOccurrenceReminderOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o DismissEventInstanceExceptionOccurrenceReminderOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o DismissEventInstanceExceptionOccurrenceReminderOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// DismissEventInstanceExceptionOccurrenceReminder - Invoke action dismissReminder. Dismiss a reminder that has been
// triggered for an event in a user calendar.
func (c EventInstanceExceptionOccurrenceClient) DismissEventInstanceExceptionOccurrenceReminder(ctx context.Context, id beta.GroupIdEventIdInstanceIdExceptionOccurrenceId, options DismissEventInstanceExceptionOccurrenceReminderOperationOptions) (result DismissEventInstanceExceptionOccurrenceReminderOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/dismissReminder", id.ID()),
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

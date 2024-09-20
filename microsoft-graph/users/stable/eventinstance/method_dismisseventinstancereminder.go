package eventinstance

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

type DismissEventInstanceReminderOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type DismissEventInstanceReminderOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultDismissEventInstanceReminderOperationOptions() DismissEventInstanceReminderOperationOptions {
	return DismissEventInstanceReminderOperationOptions{}
}

func (o DismissEventInstanceReminderOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o DismissEventInstanceReminderOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o DismissEventInstanceReminderOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// DismissEventInstanceReminder - Invoke action dismissReminder. Dismiss a reminder that has been triggered for an event
// in a user calendar.
func (c EventInstanceClient) DismissEventInstanceReminder(ctx context.Context, id stable.UserIdEventIdInstanceId, options DismissEventInstanceReminderOperationOptions) (result DismissEventInstanceReminderOperationResponse, err error) {
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

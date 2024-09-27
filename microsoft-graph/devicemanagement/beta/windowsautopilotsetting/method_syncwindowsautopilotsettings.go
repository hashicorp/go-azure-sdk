package windowsautopilotsetting

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SyncWindowsAutopilotSettingsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type SyncWindowsAutopilotSettingsOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultSyncWindowsAutopilotSettingsOperationOptions() SyncWindowsAutopilotSettingsOperationOptions {
	return SyncWindowsAutopilotSettingsOperationOptions{}
}

func (o SyncWindowsAutopilotSettingsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o SyncWindowsAutopilotSettingsOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o SyncWindowsAutopilotSettingsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// SyncWindowsAutopilotSettings - Invoke action sync. Initiates a sync of all AutoPilot registered devices from Store
// for Business and other portals. If the sync successful, this action returns a 204 No Content response code. If a sync
// is already in progress, the action returns a 409 Conflict response code. If this sync action is called within 10
// minutes of the previous sync, the action returns a 429 Too Many Requests response code.
func (c WindowsAutopilotSettingClient) SyncWindowsAutopilotSettings(ctx context.Context, options SyncWindowsAutopilotSettingsOperationOptions) (result SyncWindowsAutopilotSettingsOperationResponse, err error) {
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
		Path:          "/deviceManagement/windowsAutopilotSettings/sync",
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

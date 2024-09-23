package windowsdriverupdateprofile

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

type SyncWindowsDriverUpdateProfileInventoryOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type SyncWindowsDriverUpdateProfileInventoryOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultSyncWindowsDriverUpdateProfileInventoryOperationOptions() SyncWindowsDriverUpdateProfileInventoryOperationOptions {
	return SyncWindowsDriverUpdateProfileInventoryOperationOptions{}
}

func (o SyncWindowsDriverUpdateProfileInventoryOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o SyncWindowsDriverUpdateProfileInventoryOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o SyncWindowsDriverUpdateProfileInventoryOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// SyncWindowsDriverUpdateProfileInventory - Invoke action syncInventory. Sync the driver inventory of a
// WindowsDriverUpdateProfile.
func (c WindowsDriverUpdateProfileClient) SyncWindowsDriverUpdateProfileInventory(ctx context.Context, id beta.DeviceManagementWindowsDriverUpdateProfileId, options SyncWindowsDriverUpdateProfileInventoryOperationOptions) (result SyncWindowsDriverUpdateProfileInventoryOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/syncInventory", id.ID()),
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

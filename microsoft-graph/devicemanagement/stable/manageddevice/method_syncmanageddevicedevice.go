package manageddevice

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

type SyncManagedDeviceDeviceOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type SyncManagedDeviceDeviceOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultSyncManagedDeviceDeviceOperationOptions() SyncManagedDeviceDeviceOperationOptions {
	return SyncManagedDeviceDeviceOperationOptions{}
}

func (o SyncManagedDeviceDeviceOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o SyncManagedDeviceDeviceOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o SyncManagedDeviceDeviceOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// SyncManagedDeviceDevice - Invoke action syncDevice. Not yet documented
func (c ManagedDeviceClient) SyncManagedDeviceDevice(ctx context.Context, id stable.DeviceManagementManagedDeviceId, options SyncManagedDeviceDeviceOperationOptions) (result SyncManagedDeviceDeviceOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/syncDevice", id.ID()),
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

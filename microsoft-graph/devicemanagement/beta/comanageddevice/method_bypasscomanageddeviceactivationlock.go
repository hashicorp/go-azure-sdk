package comanageddevice

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

type BypassComanagedDeviceActivationLockOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type BypassComanagedDeviceActivationLockOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultBypassComanagedDeviceActivationLockOperationOptions() BypassComanagedDeviceActivationLockOperationOptions {
	return BypassComanagedDeviceActivationLockOperationOptions{}
}

func (o BypassComanagedDeviceActivationLockOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o BypassComanagedDeviceActivationLockOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o BypassComanagedDeviceActivationLockOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// BypassComanagedDeviceActivationLock - Invoke action bypassActivationLock. Bypass activation lock
func (c ComanagedDeviceClient) BypassComanagedDeviceActivationLock(ctx context.Context, id beta.DeviceManagementComanagedDeviceId, options BypassComanagedDeviceActivationLockOperationOptions) (result BypassComanagedDeviceActivationLockOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/bypassActivationLock", id.ID()),
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

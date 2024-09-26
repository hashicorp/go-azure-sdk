package manageddevice

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

type ResizeManagedDeviceCloudPCOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type ResizeManagedDeviceCloudPCOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultResizeManagedDeviceCloudPCOperationOptions() ResizeManagedDeviceCloudPCOperationOptions {
	return ResizeManagedDeviceCloudPCOperationOptions{}
}

func (o ResizeManagedDeviceCloudPCOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ResizeManagedDeviceCloudPCOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o ResizeManagedDeviceCloudPCOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// ResizeManagedDeviceCloudPC - Invoke action resizeCloudPc. Upgrade or downgrade an existing Cloud PC to another
// configuration with a new virtual CPU (vCPU) and storage size.
func (c ManagedDeviceClient) ResizeManagedDeviceCloudPC(ctx context.Context, id beta.DeviceManagementManagedDeviceId, input ResizeManagedDeviceCloudPCRequest, options ResizeManagedDeviceCloudPCOperationOptions) (result ResizeManagedDeviceCloudPCOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/resizeCloudPc", id.ID()),
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

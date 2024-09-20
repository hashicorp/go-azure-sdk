package manageddevice

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateManagedDeviceBulkRestoreCloudPCOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.CloudPCBulkRemoteActionResult
}

type CreateManagedDeviceBulkRestoreCloudPCOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultCreateManagedDeviceBulkRestoreCloudPCOperationOptions() CreateManagedDeviceBulkRestoreCloudPCOperationOptions {
	return CreateManagedDeviceBulkRestoreCloudPCOperationOptions{}
}

func (o CreateManagedDeviceBulkRestoreCloudPCOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateManagedDeviceBulkRestoreCloudPCOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateManagedDeviceBulkRestoreCloudPCOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateManagedDeviceBulkRestoreCloudPC - Invoke action bulkRestoreCloudPc. Restore multiple Cloud PC devices with a
// single request that includes the IDs of Intune managed devices and a restore point date and time.
func (c ManagedDeviceClient) CreateManagedDeviceBulkRestoreCloudPC(ctx context.Context, input CreateManagedDeviceBulkRestoreCloudPCRequest, options CreateManagedDeviceBulkRestoreCloudPCOperationOptions) (result CreateManagedDeviceBulkRestoreCloudPCOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          "/me/managedDevices/bulkRestoreCloudPc",
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

	var model beta.CloudPCBulkRemoteActionResult
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}

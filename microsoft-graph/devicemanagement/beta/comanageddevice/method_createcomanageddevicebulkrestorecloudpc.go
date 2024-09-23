package comanageddevice

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateComanagedDeviceBulkRestoreCloudPCOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.CloudPCBulkRemoteActionResult
}

type CreateComanagedDeviceBulkRestoreCloudPCOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateComanagedDeviceBulkRestoreCloudPCOperationOptions() CreateComanagedDeviceBulkRestoreCloudPCOperationOptions {
	return CreateComanagedDeviceBulkRestoreCloudPCOperationOptions{}
}

func (o CreateComanagedDeviceBulkRestoreCloudPCOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateComanagedDeviceBulkRestoreCloudPCOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateComanagedDeviceBulkRestoreCloudPCOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateComanagedDeviceBulkRestoreCloudPC - Invoke action bulkRestoreCloudPc. Restore multiple Cloud PC devices with a
// single request that includes the IDs of Intune managed devices and a restore point date and time.
func (c ComanagedDeviceClient) CreateComanagedDeviceBulkRestoreCloudPC(ctx context.Context, input CreateComanagedDeviceBulkRestoreCloudPCRequest, options CreateComanagedDeviceBulkRestoreCloudPCOperationOptions) (result CreateComanagedDeviceBulkRestoreCloudPCOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          "/deviceManagement/comanagedDevices/bulkRestoreCloudPc",
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

	var model beta.CloudPCBulkRemoteActionResult
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}

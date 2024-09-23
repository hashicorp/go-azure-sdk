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

type CreateComanagedDeviceBulkReprovisionCloudPCOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.CloudPCBulkRemoteActionResult
}

type CreateComanagedDeviceBulkReprovisionCloudPCOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateComanagedDeviceBulkReprovisionCloudPCOperationOptions() CreateComanagedDeviceBulkReprovisionCloudPCOperationOptions {
	return CreateComanagedDeviceBulkReprovisionCloudPCOperationOptions{}
}

func (o CreateComanagedDeviceBulkReprovisionCloudPCOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateComanagedDeviceBulkReprovisionCloudPCOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateComanagedDeviceBulkReprovisionCloudPCOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateComanagedDeviceBulkReprovisionCloudPC - Invoke action bulkReprovisionCloudPc. Bulk reprovision a set of Cloud
// PC devices with Intune managed device IDs.
func (c ComanagedDeviceClient) CreateComanagedDeviceBulkReprovisionCloudPC(ctx context.Context, input CreateComanagedDeviceBulkReprovisionCloudPCRequest, options CreateComanagedDeviceBulkReprovisionCloudPCOperationOptions) (result CreateComanagedDeviceBulkReprovisionCloudPCOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          "/deviceManagement/comanagedDevices/bulkReprovisionCloudPc",
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

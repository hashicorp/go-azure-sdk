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

type CreateComanagedDeviceBulkSetCloudPCReviewStatusOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.CloudPCBulkRemoteActionResult
}

type CreateComanagedDeviceBulkSetCloudPCReviewStatusOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateComanagedDeviceBulkSetCloudPCReviewStatusOperationOptions() CreateComanagedDeviceBulkSetCloudPCReviewStatusOperationOptions {
	return CreateComanagedDeviceBulkSetCloudPCReviewStatusOperationOptions{}
}

func (o CreateComanagedDeviceBulkSetCloudPCReviewStatusOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateComanagedDeviceBulkSetCloudPCReviewStatusOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateComanagedDeviceBulkSetCloudPCReviewStatusOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateComanagedDeviceBulkSetCloudPCReviewStatus - Invoke action bulkSetCloudPcReviewStatus. Set the review status of
// multiple Cloud PC devices with a single request that includes the IDs of Intune managed devices.
func (c ComanagedDeviceClient) CreateComanagedDeviceBulkSetCloudPCReviewStatus(ctx context.Context, input CreateComanagedDeviceBulkSetCloudPCReviewStatusRequest, options CreateComanagedDeviceBulkSetCloudPCReviewStatusOperationOptions) (result CreateComanagedDeviceBulkSetCloudPCReviewStatusOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          "/deviceManagement/comanagedDevices/bulkSetCloudPcReviewStatus",
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

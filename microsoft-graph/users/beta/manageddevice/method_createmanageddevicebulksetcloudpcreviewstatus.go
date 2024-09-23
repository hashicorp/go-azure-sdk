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

type CreateManagedDeviceBulkSetCloudPCReviewStatusOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.CloudPCBulkRemoteActionResult
}

type CreateManagedDeviceBulkSetCloudPCReviewStatusOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateManagedDeviceBulkSetCloudPCReviewStatusOperationOptions() CreateManagedDeviceBulkSetCloudPCReviewStatusOperationOptions {
	return CreateManagedDeviceBulkSetCloudPCReviewStatusOperationOptions{}
}

func (o CreateManagedDeviceBulkSetCloudPCReviewStatusOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateManagedDeviceBulkSetCloudPCReviewStatusOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateManagedDeviceBulkSetCloudPCReviewStatusOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateManagedDeviceBulkSetCloudPCReviewStatus - Invoke action bulkSetCloudPcReviewStatus. Set the review status of
// multiple Cloud PC devices with a single request that includes the IDs of Intune managed devices.
func (c ManagedDeviceClient) CreateManagedDeviceBulkSetCloudPCReviewStatus(ctx context.Context, id beta.UserId, input CreateManagedDeviceBulkSetCloudPCReviewStatusRequest, options CreateManagedDeviceBulkSetCloudPCReviewStatusOperationOptions) (result CreateManagedDeviceBulkSetCloudPCReviewStatusOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/managedDevices/bulkSetCloudPcReviewStatus", id.ID()),
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

package comanageddevicedevicecategory

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

type RemoveComanagedDeviceCategoryRefOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type RemoveComanagedDeviceCategoryRefOperationOptions struct {
	IfMatch   *string
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultRemoveComanagedDeviceCategoryRefOperationOptions() RemoveComanagedDeviceCategoryRefOperationOptions {
	return RemoveComanagedDeviceCategoryRefOperationOptions{}
}

func (o RemoveComanagedDeviceCategoryRefOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o RemoveComanagedDeviceCategoryRefOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o RemoveComanagedDeviceCategoryRefOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// RemoveComanagedDeviceCategoryRef - Delete ref of navigation property deviceCategory for deviceManagement
func (c ComanagedDeviceDeviceCategoryClient) RemoveComanagedDeviceCategoryRef(ctx context.Context, id beta.DeviceManagementComanagedDeviceId, options RemoveComanagedDeviceCategoryRefOperationOptions) (result RemoveComanagedDeviceCategoryRefOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodDelete,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/deviceCategory/$ref", id.ID()),
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

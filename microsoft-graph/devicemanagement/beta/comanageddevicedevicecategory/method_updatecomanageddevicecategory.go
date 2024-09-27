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

type UpdateComanagedDeviceCategoryOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateComanagedDeviceCategoryOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdateComanagedDeviceCategoryOperationOptions() UpdateComanagedDeviceCategoryOperationOptions {
	return UpdateComanagedDeviceCategoryOperationOptions{}
}

func (o UpdateComanagedDeviceCategoryOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateComanagedDeviceCategoryOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateComanagedDeviceCategoryOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateComanagedDeviceCategory - Update the navigation property deviceCategory in deviceManagement
func (c ComanagedDeviceDeviceCategoryClient) UpdateComanagedDeviceCategory(ctx context.Context, id beta.DeviceManagementComanagedDeviceId, input beta.DeviceCategory, options UpdateComanagedDeviceCategoryOperationOptions) (result UpdateComanagedDeviceCategoryOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/deviceCategory", id.ID()),
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

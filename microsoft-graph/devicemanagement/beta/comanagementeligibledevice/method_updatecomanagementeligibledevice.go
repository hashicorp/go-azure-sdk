package comanagementeligibledevice

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateComanagementEligibleDeviceOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateComanagementEligibleDeviceOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdateComanagementEligibleDeviceOperationOptions() UpdateComanagementEligibleDeviceOperationOptions {
	return UpdateComanagementEligibleDeviceOperationOptions{}
}

func (o UpdateComanagementEligibleDeviceOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateComanagementEligibleDeviceOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateComanagementEligibleDeviceOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateComanagementEligibleDevice - Update the navigation property comanagementEligibleDevices in deviceManagement
func (c ComanagementEligibleDeviceClient) UpdateComanagementEligibleDevice(ctx context.Context, id beta.DeviceManagementComanagementEligibleDeviceId, input beta.ComanagementEligibleDevice, options UpdateComanagementEligibleDeviceOperationOptions) (result UpdateComanagementEligibleDeviceOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          id.ID(),
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

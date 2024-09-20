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

type CreateComanagementEligibleDeviceOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.ComanagementEligibleDevice
}

type CreateComanagementEligibleDeviceOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultCreateComanagementEligibleDeviceOperationOptions() CreateComanagementEligibleDeviceOperationOptions {
	return CreateComanagementEligibleDeviceOperationOptions{}
}

func (o CreateComanagementEligibleDeviceOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateComanagementEligibleDeviceOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateComanagementEligibleDeviceOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateComanagementEligibleDevice - Create new navigation property to comanagementEligibleDevices for deviceManagement
func (c ComanagementEligibleDeviceClient) CreateComanagementEligibleDevice(ctx context.Context, input beta.ComanagementEligibleDevice, options CreateComanagementEligibleDeviceOperationOptions) (result CreateComanagementEligibleDeviceOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          "/deviceManagement/comanagementEligibleDevices",
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

	var model beta.ComanagementEligibleDevice
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}

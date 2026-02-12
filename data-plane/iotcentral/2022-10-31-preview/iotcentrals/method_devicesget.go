package iotcentrals

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DevicesGetOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *Device
}

type DevicesGetOperationOptions struct {
	Expand *string
}

func DefaultDevicesGetOperationOptions() DevicesGetOperationOptions {
	return DevicesGetOperationOptions{}
}

func (o DevicesGetOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o DevicesGetOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o DevicesGetOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Expand != nil {
		out.Append("expand", fmt.Sprintf("%v", *o.Expand))
	}
	return &out
}

// DevicesGet ...
func (c IotcentralsClient) DevicesGet(ctx context.Context, id DeviceId, options DevicesGetOperationOptions) (result DevicesGetOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          id.Path(),
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

	var model Device
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}

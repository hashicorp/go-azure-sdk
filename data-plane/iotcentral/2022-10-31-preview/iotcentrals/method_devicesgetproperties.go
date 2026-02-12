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

type DevicesGetPropertiesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *interface{}
}

type DevicesGetPropertiesOperationOptions struct {
	Unmodeled *bool
}

func DefaultDevicesGetPropertiesOperationOptions() DevicesGetPropertiesOperationOptions {
	return DevicesGetPropertiesOperationOptions{}
}

func (o DevicesGetPropertiesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o DevicesGetPropertiesOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o DevicesGetPropertiesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Unmodeled != nil {
		out.Append("unmodeled", fmt.Sprintf("%v", *o.Unmodeled))
	}
	return &out
}

// DevicesGetProperties ...
func (c IotcentralsClient) DevicesGetProperties(ctx context.Context, id DeviceId, options DevicesGetPropertiesOperationOptions) (result DevicesGetPropertiesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/properties", id.Path()),
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

	var model interface{}
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}

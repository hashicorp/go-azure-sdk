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

type DeviceTemplatesUpdateOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *DeviceTemplate
}

type DeviceTemplatesUpdateOperationOptions struct {
	IfMatch *string
}

func DefaultDeviceTemplatesUpdateOperationOptions() DeviceTemplatesUpdateOperationOptions {
	return DeviceTemplatesUpdateOperationOptions{}
}

func (o DeviceTemplatesUpdateOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o DeviceTemplatesUpdateOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o DeviceTemplatesUpdateOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// DeviceTemplatesUpdate ...
func (c IotcentralsClient) DeviceTemplatesUpdate(ctx context.Context, id DeviceTemplateId, input interface{}, options DeviceTemplatesUpdateOperationOptions) (result DeviceTemplatesUpdateOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          id.Path(),
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

	var model DeviceTemplate
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}

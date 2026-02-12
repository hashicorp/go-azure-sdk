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

type DevicesRemoveOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type DevicesRemoveOperationOptions struct {
	Expand *string
}

func DefaultDevicesRemoveOperationOptions() DevicesRemoveOperationOptions {
	return DevicesRemoveOperationOptions{}
}

func (o DevicesRemoveOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o DevicesRemoveOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o DevicesRemoveOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Expand != nil {
		out.Append("expand", fmt.Sprintf("%v", *o.Expand))
	}
	return &out
}

// DevicesRemove ...
func (c IotcentralsClient) DevicesRemove(ctx context.Context, id DeviceId, options DevicesRemoveOperationOptions) (result DevicesRemoveOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodDelete,
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

	return
}

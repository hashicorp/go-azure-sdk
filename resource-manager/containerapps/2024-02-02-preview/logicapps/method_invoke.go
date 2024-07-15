package logicapps

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InvokeOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *interface{}
}

type InvokeOperationOptions struct {
	XMsLogicAppsProxyMethod *LogicAppsProxyMethod
	XMsLogicAppsProxyPath   *string
}

func DefaultInvokeOperationOptions() InvokeOperationOptions {
	return InvokeOperationOptions{}
}

func (o InvokeOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.XMsLogicAppsProxyMethod != nil {
		out.Append("x-ms-logicApps-proxy-method", fmt.Sprintf("%v", *o.XMsLogicAppsProxyMethod))
	}
	if o.XMsLogicAppsProxyPath != nil {
		out.Append("x-ms-logicApps-proxy-path", fmt.Sprintf("%v", *o.XMsLogicAppsProxyPath))
	}
	return &out
}

func (o InvokeOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	return &out
}

func (o InvokeOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// Invoke ...
func (c LogicAppsClient) Invoke(ctx context.Context, id LogicAppId, options InvokeOperationOptions) (result InvokeOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/invoke", id.ID()),
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

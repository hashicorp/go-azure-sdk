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

type DashboardsUpdateOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *Dashboard
}

type DashboardsUpdateOperationOptions struct {
	IfMatch *string
}

func DefaultDashboardsUpdateOperationOptions() DashboardsUpdateOperationOptions {
	return DashboardsUpdateOperationOptions{}
}

func (o DashboardsUpdateOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o DashboardsUpdateOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o DashboardsUpdateOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// DashboardsUpdate ...
func (c IotcentralsClient) DashboardsUpdate(ctx context.Context, id DashboardId, input interface{}, options DashboardsUpdateOperationOptions) (result DashboardsUpdateOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/merge-patch+json",
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

	var model Dashboard
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}

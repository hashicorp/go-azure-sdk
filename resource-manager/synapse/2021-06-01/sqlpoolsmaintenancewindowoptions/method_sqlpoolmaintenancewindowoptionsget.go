package sqlpoolsmaintenancewindowoptions

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SqlPoolMaintenanceWindowOptionsGetOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *MaintenanceWindowOptions
}

type SqlPoolMaintenanceWindowOptionsGetOperationOptions struct {
	MaintenanceWindowOptionsName *string
}

func DefaultSqlPoolMaintenanceWindowOptionsGetOperationOptions() SqlPoolMaintenanceWindowOptionsGetOperationOptions {
	return SqlPoolMaintenanceWindowOptionsGetOperationOptions{}
}

func (o SqlPoolMaintenanceWindowOptionsGetOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o SqlPoolMaintenanceWindowOptionsGetOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	return &out
}

func (o SqlPoolMaintenanceWindowOptionsGetOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.MaintenanceWindowOptionsName != nil {
		out.Append("maintenanceWindowOptionsName", fmt.Sprintf("%v", *o.MaintenanceWindowOptionsName))
	}
	return &out
}

// SqlPoolMaintenanceWindowOptionsGet ...
func (c SqlPoolsMaintenanceWindowOptionsClient) SqlPoolMaintenanceWindowOptionsGet(ctx context.Context, id SqlPoolId, options SqlPoolMaintenanceWindowOptionsGetOperationOptions) (result SqlPoolMaintenanceWindowOptionsGetOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		Path:          fmt.Sprintf("%s/maintenanceWindowOptions/current", id.ID()),
		OptionsObject: options,
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

	var model MaintenanceWindowOptions
	result.Model = &model

	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}

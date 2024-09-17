package sqlpoolsmaintenancewindows

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SqlPoolMaintenanceWindowsGetOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *MaintenanceWindows
}

type SqlPoolMaintenanceWindowsGetOperationOptions struct {
	MaintenanceWindowName *string
}

func DefaultSqlPoolMaintenanceWindowsGetOperationOptions() SqlPoolMaintenanceWindowsGetOperationOptions {
	return SqlPoolMaintenanceWindowsGetOperationOptions{}
}

func (o SqlPoolMaintenanceWindowsGetOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o SqlPoolMaintenanceWindowsGetOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o SqlPoolMaintenanceWindowsGetOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.MaintenanceWindowName != nil {
		out.Append("maintenanceWindowName", fmt.Sprintf("%v", *o.MaintenanceWindowName))
	}
	return &out
}

// SqlPoolMaintenanceWindowsGet ...
func (c SqlPoolsMaintenanceWindowsClient) SqlPoolMaintenanceWindowsGet(ctx context.Context, id SqlPoolId, options SqlPoolMaintenanceWindowsGetOperationOptions) (result SqlPoolMaintenanceWindowsGetOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/maintenancewindows/current", id.ID()),
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

	var model MaintenanceWindows
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}

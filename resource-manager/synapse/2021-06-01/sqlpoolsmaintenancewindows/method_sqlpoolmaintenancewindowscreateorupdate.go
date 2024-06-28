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

type SqlPoolMaintenanceWindowsCreateOrUpdateOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type SqlPoolMaintenanceWindowsCreateOrUpdateOperationOptions struct {
	MaintenanceWindowName *string
}

func DefaultSqlPoolMaintenanceWindowsCreateOrUpdateOperationOptions() SqlPoolMaintenanceWindowsCreateOrUpdateOperationOptions {
	return SqlPoolMaintenanceWindowsCreateOrUpdateOperationOptions{}
}

func (o SqlPoolMaintenanceWindowsCreateOrUpdateOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o SqlPoolMaintenanceWindowsCreateOrUpdateOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	return &out
}

func (o SqlPoolMaintenanceWindowsCreateOrUpdateOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.MaintenanceWindowName != nil {
		out.Append("maintenanceWindowName", fmt.Sprintf("%v", *o.MaintenanceWindowName))
	}
	return &out
}

// SqlPoolMaintenanceWindowsCreateOrUpdate ...
func (c SqlPoolsMaintenanceWindowsClient) SqlPoolMaintenanceWindowsCreateOrUpdate(ctx context.Context, id SqlPoolId, input MaintenanceWindows, options SqlPoolMaintenanceWindowsCreateOrUpdateOperationOptions) (result SqlPoolMaintenanceWindowsCreateOrUpdateOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodPut,

		Path:          fmt.Sprintf("%s/maintenancewindows/current", id.ID()),
		OptionsObject: options,
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

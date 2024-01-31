package maintenancewindows

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateOrUpdateOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type CreateOrUpdateOperationOptions struct {
	MaintenanceWindowName *string
}

func DefaultCreateOrUpdateOperationOptions() CreateOrUpdateOperationOptions {
	return CreateOrUpdateOperationOptions{}
}

func (o CreateOrUpdateOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateOrUpdateOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	return &out
}

func (o CreateOrUpdateOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.MaintenanceWindowName != nil {
		out.Append("maintenanceWindowName", fmt.Sprintf("%v", *o.MaintenanceWindowName))
	}
	return &out
}

// CreateOrUpdate ...
func (c MaintenanceWindowsClient) CreateOrUpdate(ctx context.Context, id commonids.SqlDatabaseId, input MaintenanceWindows, options CreateOrUpdateOperationOptions) (result CreateOrUpdateOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPut,
		Path:          fmt.Sprintf("%s/maintenanceWindows/current", id.ID()),
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

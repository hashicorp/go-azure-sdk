package windowsdriverupdateprofiledriverinventory

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeleteWindowsDriverUpdateProfileDriverInventoryOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type DeleteWindowsDriverUpdateProfileDriverInventoryOperationOptions struct {
	IfMatch  *string
	Metadata *odata.Metadata
}

func DefaultDeleteWindowsDriverUpdateProfileDriverInventoryOperationOptions() DeleteWindowsDriverUpdateProfileDriverInventoryOperationOptions {
	return DeleteWindowsDriverUpdateProfileDriverInventoryOperationOptions{}
}

func (o DeleteWindowsDriverUpdateProfileDriverInventoryOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o DeleteWindowsDriverUpdateProfileDriverInventoryOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o DeleteWindowsDriverUpdateProfileDriverInventoryOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// DeleteWindowsDriverUpdateProfileDriverInventory - Delete navigation property driverInventories for deviceManagement
func (c WindowsDriverUpdateProfileDriverInventoryClient) DeleteWindowsDriverUpdateProfileDriverInventory(ctx context.Context, id beta.DeviceManagementWindowsDriverUpdateProfileIdDriverInventoryId, options DeleteWindowsDriverUpdateProfileDriverInventoryOperationOptions) (result DeleteWindowsDriverUpdateProfileDriverInventoryOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodDelete,
		OptionsObject: options,
		Path:          id.ID(),
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

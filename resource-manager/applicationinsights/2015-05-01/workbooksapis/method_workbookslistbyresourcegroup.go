package workbooksapis

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

type WorkbooksListByResourceGroupOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *WorkbooksListResult
}

type WorkbooksListByResourceGroupOperationOptions struct {
	CanFetchContent *bool
	Category        *CategoryType
	Tags            *string
}

func DefaultWorkbooksListByResourceGroupOperationOptions() WorkbooksListByResourceGroupOperationOptions {
	return WorkbooksListByResourceGroupOperationOptions{}
}

func (o WorkbooksListByResourceGroupOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o WorkbooksListByResourceGroupOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	return &out
}

func (o WorkbooksListByResourceGroupOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.CanFetchContent != nil {
		out.Append("canFetchContent", fmt.Sprintf("%v", *o.CanFetchContent))
	}
	if o.Category != nil {
		out.Append("category", fmt.Sprintf("%v", *o.Category))
	}
	if o.Tags != nil {
		out.Append("tags", fmt.Sprintf("%v", *o.Tags))
	}
	return &out
}

// WorkbooksListByResourceGroup ...
func (c WorkbooksAPIsClient) WorkbooksListByResourceGroup(ctx context.Context, id commonids.ResourceGroupId, options WorkbooksListByResourceGroupOperationOptions) (result WorkbooksListByResourceGroupOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,

		Path:          fmt.Sprintf("%s/providers/Microsoft.Insights/workbooks", id.ID()),
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

	var model WorkbooksListResult
	result.Model = &model

	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}

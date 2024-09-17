package myworkbooksapis

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

type MyWorkbooksListByResourceGroupOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *MyWorkbooksListResult
}

type MyWorkbooksListByResourceGroupOperationOptions struct {
	CanFetchContent *bool
	Category        *CategoryType
	Tags            *string
}

func DefaultMyWorkbooksListByResourceGroupOperationOptions() MyWorkbooksListByResourceGroupOperationOptions {
	return MyWorkbooksListByResourceGroupOperationOptions{}
}

func (o MyWorkbooksListByResourceGroupOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o MyWorkbooksListByResourceGroupOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o MyWorkbooksListByResourceGroupOperationOptions) ToQuery() *client.QueryParams {
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

// MyWorkbooksListByResourceGroup ...
func (c MyworkbooksAPIsClient) MyWorkbooksListByResourceGroup(ctx context.Context, id commonids.ResourceGroupId, options MyWorkbooksListByResourceGroupOperationOptions) (result MyWorkbooksListByResourceGroupOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/providers/Microsoft.Insights/myWorkbooks", id.ID()),
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

	var model MyWorkbooksListResult
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}

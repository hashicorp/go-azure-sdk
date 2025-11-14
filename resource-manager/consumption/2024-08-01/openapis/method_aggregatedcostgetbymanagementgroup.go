package openapis

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

type AggregatedCostGetByManagementGroupOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *ManagementGroupAggregatedCostResult
}

type AggregatedCostGetByManagementGroupOperationOptions struct {
	Filter *string
}

func DefaultAggregatedCostGetByManagementGroupOperationOptions() AggregatedCostGetByManagementGroupOperationOptions {
	return AggregatedCostGetByManagementGroupOperationOptions{}
}

func (o AggregatedCostGetByManagementGroupOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o AggregatedCostGetByManagementGroupOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o AggregatedCostGetByManagementGroupOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Filter != nil {
		out.Append("$filter", fmt.Sprintf("%v", *o.Filter))
	}
	return &out
}

// AggregatedCostGetByManagementGroup ...
func (c OpenapisClient) AggregatedCostGetByManagementGroup(ctx context.Context, id commonids.ManagementGroupId, options AggregatedCostGetByManagementGroupOperationOptions) (result AggregatedCostGetByManagementGroupOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/providers/Microsoft.Consumption/aggregatedcost", id.ID()),
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

	var model ManagementGroupAggregatedCostResult
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}

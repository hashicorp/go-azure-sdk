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

type ForecastUsageOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *ForecastResult
}

type ForecastUsageOperationOptions struct {
	Filter *string
}

func DefaultForecastUsageOperationOptions() ForecastUsageOperationOptions {
	return ForecastUsageOperationOptions{}
}

func (o ForecastUsageOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ForecastUsageOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o ForecastUsageOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Filter != nil {
		out.Append("$filter", fmt.Sprintf("%v", *o.Filter))
	}
	return &out
}

// ForecastUsage ...
func (c OpenapisClient) ForecastUsage(ctx context.Context, id commonids.ScopeId, input ForecastDefinition, options ForecastUsageOperationOptions) (result ForecastUsageOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/providers/Microsoft.CostManagement/forecast", id.ID()),
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

	var model ForecastResult
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}

package insighttrendingresource

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetInsightTrendingResourceOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        stable.Entity
}

type GetInsightTrendingResourceOperationOptions struct {
	Expand *odata.Expand
	Select *[]string
}

func DefaultGetInsightTrendingResourceOperationOptions() GetInsightTrendingResourceOperationOptions {
	return GetInsightTrendingResourceOperationOptions{}
}

func (o GetInsightTrendingResourceOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetInsightTrendingResourceOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetInsightTrendingResourceOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetInsightTrendingResource - Get resource from users. Used for navigating to the trending document.
func (c InsightTrendingResourceClient) GetInsightTrendingResource(ctx context.Context, id stable.UserIdInsightTrendingId, options GetInsightTrendingResourceOperationOptions) (result GetInsightTrendingResourceOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/resource", id.ID()),
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

	var respObj json.RawMessage
	if err = resp.Unmarshal(&respObj); err != nil {
		return
	}
	model, err := stable.UnmarshalEntityImplementation(respObj)
	if err != nil {
		return
	}
	result.Model = model

	return
}

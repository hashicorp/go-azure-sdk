package contentproductpackages

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProductPackagesListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]ProductPackageModel
}

type ProductPackagesListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []ProductPackageModel
}

type ProductPackagesListOperationOptions struct {
	Filter  *string
	Orderby *string
	Top     *int64
}

func DefaultProductPackagesListOperationOptions() ProductPackagesListOperationOptions {
	return ProductPackagesListOperationOptions{}
}

func (o ProductPackagesListOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ProductPackagesListOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	return &out
}

func (o ProductPackagesListOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Filter != nil {
		out.Append("$filter", fmt.Sprintf("%v", *o.Filter))
	}
	if o.Orderby != nil {
		out.Append("$orderby", fmt.Sprintf("%v", *o.Orderby))
	}
	if o.Top != nil {
		out.Append("$top", fmt.Sprintf("%v", *o.Top))
	}
	return &out
}

// ProductPackagesList ...
func (c ContentProductPackagesClient) ProductPackagesList(ctx context.Context, id WorkspaceId, options ProductPackagesListOperationOptions) (result ProductPackagesListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		Path:          fmt.Sprintf("%s/providers/Microsoft.SecurityInsights/contentProductPackages", id.ID()),
		OptionsObject: options,
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	var resp *client.Response
	resp, err = req.ExecutePaged(ctx)
	if resp != nil {
		result.OData = resp.OData
		result.HttpResponse = resp.Response
	}
	if err != nil {
		return
	}

	var values struct {
		Values *[]ProductPackageModel `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ProductPackagesListComplete retrieves all the results into a single object
func (c ContentProductPackagesClient) ProductPackagesListComplete(ctx context.Context, id WorkspaceId, options ProductPackagesListOperationOptions) (ProductPackagesListCompleteResult, error) {
	return c.ProductPackagesListCompleteMatchingPredicate(ctx, id, options, ProductPackageModelOperationPredicate{})
}

// ProductPackagesListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ContentProductPackagesClient) ProductPackagesListCompleteMatchingPredicate(ctx context.Context, id WorkspaceId, options ProductPackagesListOperationOptions, predicate ProductPackageModelOperationPredicate) (result ProductPackagesListCompleteResult, err error) {
	items := make([]ProductPackageModel, 0)

	resp, err := c.ProductPackagesList(ctx, id, options)
	if err != nil {
		err = fmt.Errorf("loading results: %+v", err)
		return
	}
	if resp.Model != nil {
		for _, v := range *resp.Model {
			if predicate.Matches(v) {
				items = append(items, v)
			}
		}
	}

	result = ProductPackagesListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}

package contentproducttemplates

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProductTemplatesListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]ProductTemplateModel
}

type ProductTemplatesListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []ProductTemplateModel
}

type ProductTemplatesListOperationOptions struct {
	Count   *bool
	Filter  *string
	Orderby *string
	Search  *string
	Skip    *int64
	Top     *int64
}

func DefaultProductTemplatesListOperationOptions() ProductTemplatesListOperationOptions {
	return ProductTemplatesListOperationOptions{}
}

func (o ProductTemplatesListOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ProductTemplatesListOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	return &out
}

func (o ProductTemplatesListOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Count != nil {
		out.Append("$count", fmt.Sprintf("%v", *o.Count))
	}
	if o.Filter != nil {
		out.Append("$filter", fmt.Sprintf("%v", *o.Filter))
	}
	if o.Orderby != nil {
		out.Append("$orderby", fmt.Sprintf("%v", *o.Orderby))
	}
	if o.Search != nil {
		out.Append("$search", fmt.Sprintf("%v", *o.Search))
	}
	if o.Skip != nil {
		out.Append("$skip", fmt.Sprintf("%v", *o.Skip))
	}
	if o.Top != nil {
		out.Append("$top", fmt.Sprintf("%v", *o.Top))
	}
	return &out
}

// ProductTemplatesList ...
func (c ContentProductTemplatesClient) ProductTemplatesList(ctx context.Context, id WorkspaceId, options ProductTemplatesListOperationOptions) (result ProductTemplatesListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		Path:          fmt.Sprintf("%s/providers/Microsoft.SecurityInsights/contentProductTemplates", id.ID()),
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
		Values *[]ProductTemplateModel `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ProductTemplatesListComplete retrieves all the results into a single object
func (c ContentProductTemplatesClient) ProductTemplatesListComplete(ctx context.Context, id WorkspaceId, options ProductTemplatesListOperationOptions) (ProductTemplatesListCompleteResult, error) {
	return c.ProductTemplatesListCompleteMatchingPredicate(ctx, id, options, ProductTemplateModelOperationPredicate{})
}

// ProductTemplatesListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ContentProductTemplatesClient) ProductTemplatesListCompleteMatchingPredicate(ctx context.Context, id WorkspaceId, options ProductTemplatesListOperationOptions, predicate ProductTemplateModelOperationPredicate) (result ProductTemplatesListCompleteResult, err error) {
	items := make([]ProductTemplateModel, 0)

	resp, err := c.ProductTemplatesList(ctx, id, options)
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

	result = ProductTemplatesListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}

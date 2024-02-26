package products

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByInvoiceSectionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]Product
}

type ListByInvoiceSectionCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []Product
}

type ListByInvoiceSectionOperationOptions struct {
	Filter *string
}

func DefaultListByInvoiceSectionOperationOptions() ListByInvoiceSectionOperationOptions {
	return ListByInvoiceSectionOperationOptions{}
}

func (o ListByInvoiceSectionOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListByInvoiceSectionOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	return &out
}

func (o ListByInvoiceSectionOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Filter != nil {
		out.Append("$filter", fmt.Sprintf("%v", *o.Filter))
	}
	return &out
}

// ListByInvoiceSection ...
func (c ProductsClient) ListByInvoiceSection(ctx context.Context, id InvoiceSectionId, options ListByInvoiceSectionOperationOptions) (result ListByInvoiceSectionOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		Path:          fmt.Sprintf("%s/products", id.ID()),
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
		Values *[]Product `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListByInvoiceSectionComplete retrieves all the results into a single object
func (c ProductsClient) ListByInvoiceSectionComplete(ctx context.Context, id InvoiceSectionId, options ListByInvoiceSectionOperationOptions) (ListByInvoiceSectionCompleteResult, error) {
	return c.ListByInvoiceSectionCompleteMatchingPredicate(ctx, id, options, ProductOperationPredicate{})
}

// ListByInvoiceSectionCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ProductsClient) ListByInvoiceSectionCompleteMatchingPredicate(ctx context.Context, id InvoiceSectionId, options ListByInvoiceSectionOperationOptions, predicate ProductOperationPredicate) (result ListByInvoiceSectionCompleteResult, err error) {
	items := make([]Product, 0)

	resp, err := c.ListByInvoiceSection(ctx, id, options)
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

	result = ListByInvoiceSectionCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}

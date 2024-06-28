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

type ListByCustomerOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]Product
}

type ListByCustomerCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []Product
}

type ListByCustomerCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ListByCustomerCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListByCustomer ...
func (c ProductsClient) ListByCustomer(ctx context.Context, id CustomerId) (result ListByCustomerOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListByCustomerCustomPager{},
		Path:       fmt.Sprintf("%s/products", id.ID()),
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

// ListByCustomerComplete retrieves all the results into a single object
func (c ProductsClient) ListByCustomerComplete(ctx context.Context, id CustomerId) (ListByCustomerCompleteResult, error) {
	return c.ListByCustomerCompleteMatchingPredicate(ctx, id, ProductOperationPredicate{})
}

// ListByCustomerCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ProductsClient) ListByCustomerCompleteMatchingPredicate(ctx context.Context, id CustomerId, predicate ProductOperationPredicate) (result ListByCustomerCompleteResult, err error) {
	items := make([]Product, 0)

	resp, err := c.ListByCustomer(ctx, id)
	if err != nil {
		result.LatestHttpResponse = resp.HttpResponse
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

	result = ListByCustomerCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}

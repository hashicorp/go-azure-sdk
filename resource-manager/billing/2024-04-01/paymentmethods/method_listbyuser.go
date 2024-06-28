package paymentmethods

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByUserOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]PaymentMethod
}

type ListByUserCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []PaymentMethod
}

type ListByUserCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ListByUserCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListByUser ...
func (c PaymentMethodsClient) ListByUser(ctx context.Context) (result ListByUserOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListByUserCustomPager{},
		Path:       "/providers/Microsoft.Billing/paymentMethods",
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
		Values *[]PaymentMethod `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListByUserComplete retrieves all the results into a single object
func (c PaymentMethodsClient) ListByUserComplete(ctx context.Context) (ListByUserCompleteResult, error) {
	return c.ListByUserCompleteMatchingPredicate(ctx, PaymentMethodOperationPredicate{})
}

// ListByUserCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c PaymentMethodsClient) ListByUserCompleteMatchingPredicate(ctx context.Context, predicate PaymentMethodOperationPredicate) (result ListByUserCompleteResult, err error) {
	items := make([]PaymentMethod, 0)

	resp, err := c.ListByUser(ctx)
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

	result = ListByUserCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}

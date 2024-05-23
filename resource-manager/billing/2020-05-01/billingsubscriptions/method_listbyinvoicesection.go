package billingsubscriptions

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
	Model        *[]BillingSubscription
}

type ListByInvoiceSectionCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []BillingSubscription
}

// ListByInvoiceSection ...
func (c BillingSubscriptionsClient) ListByInvoiceSection(ctx context.Context, id InvoiceSectionId) (result ListByInvoiceSectionOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/billingSubscriptions", id.ID()),
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
		Values *[]BillingSubscription `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListByInvoiceSectionComplete retrieves all the results into a single object
func (c BillingSubscriptionsClient) ListByInvoiceSectionComplete(ctx context.Context, id InvoiceSectionId) (ListByInvoiceSectionCompleteResult, error) {
	return c.ListByInvoiceSectionCompleteMatchingPredicate(ctx, id, BillingSubscriptionOperationPredicate{})
}

// ListByInvoiceSectionCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c BillingSubscriptionsClient) ListByInvoiceSectionCompleteMatchingPredicate(ctx context.Context, id InvoiceSectionId, predicate BillingSubscriptionOperationPredicate) (result ListByInvoiceSectionCompleteResult, err error) {
	items := make([]BillingSubscription, 0)

	resp, err := c.ListByInvoiceSection(ctx, id)
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

	result = ListByInvoiceSectionCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}

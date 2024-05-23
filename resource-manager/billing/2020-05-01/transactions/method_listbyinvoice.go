package transactions

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByInvoiceOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]Transaction
}

type ListByInvoiceCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []Transaction
}

// ListByInvoice ...
func (c TransactionsClient) ListByInvoice(ctx context.Context, id BillingAccountInvoiceId) (result ListByInvoiceOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/transactions", id.ID()),
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
		Values *[]Transaction `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListByInvoiceComplete retrieves all the results into a single object
func (c TransactionsClient) ListByInvoiceComplete(ctx context.Context, id BillingAccountInvoiceId) (ListByInvoiceCompleteResult, error) {
	return c.ListByInvoiceCompleteMatchingPredicate(ctx, id, TransactionOperationPredicate{})
}

// ListByInvoiceCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c TransactionsClient) ListByInvoiceCompleteMatchingPredicate(ctx context.Context, id BillingAccountInvoiceId, predicate TransactionOperationPredicate) (result ListByInvoiceCompleteResult, err error) {
	items := make([]Transaction, 0)

	resp, err := c.ListByInvoice(ctx, id)
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

	result = ListByInvoiceCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}

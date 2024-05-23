package invoices

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByBillingProfileOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]Invoice
}

type ListByBillingProfileCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []Invoice
}

type ListByBillingProfileOperationOptions struct {
	PeriodEndDate   *string
	PeriodStartDate *string
}

func DefaultListByBillingProfileOperationOptions() ListByBillingProfileOperationOptions {
	return ListByBillingProfileOperationOptions{}
}

func (o ListByBillingProfileOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListByBillingProfileOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	return &out
}

func (o ListByBillingProfileOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.PeriodEndDate != nil {
		out.Append("periodEndDate", fmt.Sprintf("%v", *o.PeriodEndDate))
	}
	if o.PeriodStartDate != nil {
		out.Append("periodStartDate", fmt.Sprintf("%v", *o.PeriodStartDate))
	}
	return &out
}

// ListByBillingProfile ...
func (c InvoicesClient) ListByBillingProfile(ctx context.Context, id BillingProfileId, options ListByBillingProfileOperationOptions) (result ListByBillingProfileOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		Path:          fmt.Sprintf("%s/invoices", id.ID()),
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
		Values *[]Invoice `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListByBillingProfileComplete retrieves all the results into a single object
func (c InvoicesClient) ListByBillingProfileComplete(ctx context.Context, id BillingProfileId, options ListByBillingProfileOperationOptions) (ListByBillingProfileCompleteResult, error) {
	return c.ListByBillingProfileCompleteMatchingPredicate(ctx, id, options, InvoiceOperationPredicate{})
}

// ListByBillingProfileCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c InvoicesClient) ListByBillingProfileCompleteMatchingPredicate(ctx context.Context, id BillingProfileId, options ListByBillingProfileOperationOptions, predicate InvoiceOperationPredicate) (result ListByBillingProfileCompleteResult, err error) {
	items := make([]Invoice, 0)

	resp, err := c.ListByBillingProfile(ctx, id, options)
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

	result = ListByBillingProfileCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}

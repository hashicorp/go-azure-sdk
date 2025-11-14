package openapis

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ReservationTransactionsListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]ReservationTransaction
}

type ReservationTransactionsListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []ReservationTransaction
}

type ReservationTransactionsListOperationOptions struct {
	Filter                  *string
	PreviewMarkupPercentage *float64
	UseMarkupIfPartner      *bool
}

func DefaultReservationTransactionsListOperationOptions() ReservationTransactionsListOperationOptions {
	return ReservationTransactionsListOperationOptions{}
}

func (o ReservationTransactionsListOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ReservationTransactionsListOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o ReservationTransactionsListOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Filter != nil {
		out.Append("$filter", fmt.Sprintf("%v", *o.Filter))
	}
	if o.PreviewMarkupPercentage != nil {
		out.Append("previewMarkupPercentage", fmt.Sprintf("%v", *o.PreviewMarkupPercentage))
	}
	if o.UseMarkupIfPartner != nil {
		out.Append("useMarkupIfPartner", fmt.Sprintf("%v", *o.UseMarkupIfPartner))
	}
	return &out
}

type ReservationTransactionsListCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ReservationTransactionsListCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ReservationTransactionsList ...
func (c OpenapisClient) ReservationTransactionsList(ctx context.Context, id BillingAccountId, options ReservationTransactionsListOperationOptions) (result ReservationTransactionsListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ReservationTransactionsListCustomPager{},
		Path:          fmt.Sprintf("%s/providers/Microsoft.Consumption/reservationTransactions", id.ID()),
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
		Values *[]ReservationTransaction `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ReservationTransactionsListComplete retrieves all the results into a single object
func (c OpenapisClient) ReservationTransactionsListComplete(ctx context.Context, id BillingAccountId, options ReservationTransactionsListOperationOptions) (ReservationTransactionsListCompleteResult, error) {
	return c.ReservationTransactionsListCompleteMatchingPredicate(ctx, id, options, ReservationTransactionOperationPredicate{})
}

// ReservationTransactionsListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c OpenapisClient) ReservationTransactionsListCompleteMatchingPredicate(ctx context.Context, id BillingAccountId, options ReservationTransactionsListOperationOptions, predicate ReservationTransactionOperationPredicate) (result ReservationTransactionsListCompleteResult, err error) {
	items := make([]ReservationTransaction, 0)

	resp, err := c.ReservationTransactionsList(ctx, id, options)
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

	result = ReservationTransactionsListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}

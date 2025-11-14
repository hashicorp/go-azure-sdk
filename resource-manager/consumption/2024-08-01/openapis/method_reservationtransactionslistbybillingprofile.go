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

type ReservationTransactionsListByBillingProfileOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]ModernReservationTransaction
}

type ReservationTransactionsListByBillingProfileCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []ModernReservationTransaction
}

type ReservationTransactionsListByBillingProfileOperationOptions struct {
	Filter *string
}

func DefaultReservationTransactionsListByBillingProfileOperationOptions() ReservationTransactionsListByBillingProfileOperationOptions {
	return ReservationTransactionsListByBillingProfileOperationOptions{}
}

func (o ReservationTransactionsListByBillingProfileOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ReservationTransactionsListByBillingProfileOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o ReservationTransactionsListByBillingProfileOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Filter != nil {
		out.Append("$filter", fmt.Sprintf("%v", *o.Filter))
	}
	return &out
}

type ReservationTransactionsListByBillingProfileCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ReservationTransactionsListByBillingProfileCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ReservationTransactionsListByBillingProfile ...
func (c OpenapisClient) ReservationTransactionsListByBillingProfile(ctx context.Context, id BillingProfileId, options ReservationTransactionsListByBillingProfileOperationOptions) (result ReservationTransactionsListByBillingProfileOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ReservationTransactionsListByBillingProfileCustomPager{},
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
		Values *[]ModernReservationTransaction `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ReservationTransactionsListByBillingProfileComplete retrieves all the results into a single object
func (c OpenapisClient) ReservationTransactionsListByBillingProfileComplete(ctx context.Context, id BillingProfileId, options ReservationTransactionsListByBillingProfileOperationOptions) (ReservationTransactionsListByBillingProfileCompleteResult, error) {
	return c.ReservationTransactionsListByBillingProfileCompleteMatchingPredicate(ctx, id, options, ModernReservationTransactionOperationPredicate{})
}

// ReservationTransactionsListByBillingProfileCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c OpenapisClient) ReservationTransactionsListByBillingProfileCompleteMatchingPredicate(ctx context.Context, id BillingProfileId, options ReservationTransactionsListByBillingProfileOperationOptions, predicate ModernReservationTransactionOperationPredicate) (result ReservationTransactionsListByBillingProfileCompleteResult, err error) {
	items := make([]ModernReservationTransaction, 0)

	resp, err := c.ReservationTransactionsListByBillingProfile(ctx, id, options)
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

	result = ReservationTransactionsListByBillingProfileCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}

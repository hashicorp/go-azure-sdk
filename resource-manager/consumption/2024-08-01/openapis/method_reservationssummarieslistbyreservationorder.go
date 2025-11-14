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

type ReservationsSummariesListByReservationOrderOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]ReservationSummary
}

type ReservationsSummariesListByReservationOrderCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []ReservationSummary
}

type ReservationsSummariesListByReservationOrderOperationOptions struct {
	Filter *string
	Grain  *Datagrain
}

func DefaultReservationsSummariesListByReservationOrderOperationOptions() ReservationsSummariesListByReservationOrderOperationOptions {
	return ReservationsSummariesListByReservationOrderOperationOptions{}
}

func (o ReservationsSummariesListByReservationOrderOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ReservationsSummariesListByReservationOrderOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o ReservationsSummariesListByReservationOrderOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Filter != nil {
		out.Append("$filter", fmt.Sprintf("%v", *o.Filter))
	}
	if o.Grain != nil {
		out.Append("grain", fmt.Sprintf("%v", *o.Grain))
	}
	return &out
}

type ReservationsSummariesListByReservationOrderCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ReservationsSummariesListByReservationOrderCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ReservationsSummariesListByReservationOrder ...
func (c OpenapisClient) ReservationsSummariesListByReservationOrder(ctx context.Context, id ReservationOrderId, options ReservationsSummariesListByReservationOrderOperationOptions) (result ReservationsSummariesListByReservationOrderOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ReservationsSummariesListByReservationOrderCustomPager{},
		Path:          fmt.Sprintf("%s/providers/Microsoft.Consumption/reservationSummaries", id.ID()),
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
		Values *[]ReservationSummary `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ReservationsSummariesListByReservationOrderComplete retrieves all the results into a single object
func (c OpenapisClient) ReservationsSummariesListByReservationOrderComplete(ctx context.Context, id ReservationOrderId, options ReservationsSummariesListByReservationOrderOperationOptions) (ReservationsSummariesListByReservationOrderCompleteResult, error) {
	return c.ReservationsSummariesListByReservationOrderCompleteMatchingPredicate(ctx, id, options, ReservationSummaryOperationPredicate{})
}

// ReservationsSummariesListByReservationOrderCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c OpenapisClient) ReservationsSummariesListByReservationOrderCompleteMatchingPredicate(ctx context.Context, id ReservationOrderId, options ReservationsSummariesListByReservationOrderOperationOptions, predicate ReservationSummaryOperationPredicate) (result ReservationsSummariesListByReservationOrderCompleteResult, err error) {
	items := make([]ReservationSummary, 0)

	resp, err := c.ReservationsSummariesListByReservationOrder(ctx, id, options)
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

	result = ReservationsSummariesListByReservationOrderCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}

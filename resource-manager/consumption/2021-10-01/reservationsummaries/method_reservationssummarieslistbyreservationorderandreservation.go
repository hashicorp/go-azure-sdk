package reservationsummaries

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ReservationsSummariesListByReservationOrderAndReservationOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]ReservationSummary
}

type ReservationsSummariesListByReservationOrderAndReservationCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []ReservationSummary
}

type ReservationsSummariesListByReservationOrderAndReservationOperationOptions struct {
	Filter *string
	Grain  *Datagrain
}

func DefaultReservationsSummariesListByReservationOrderAndReservationOperationOptions() ReservationsSummariesListByReservationOrderAndReservationOperationOptions {
	return ReservationsSummariesListByReservationOrderAndReservationOperationOptions{}
}

func (o ReservationsSummariesListByReservationOrderAndReservationOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ReservationsSummariesListByReservationOrderAndReservationOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	return &out
}

func (o ReservationsSummariesListByReservationOrderAndReservationOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Filter != nil {
		out.Append("$filter", fmt.Sprintf("%v", *o.Filter))
	}
	if o.Grain != nil {
		out.Append("grain", fmt.Sprintf("%v", *o.Grain))
	}
	return &out
}

// ReservationsSummariesListByReservationOrderAndReservation ...
func (c ReservationSummariesClient) ReservationsSummariesListByReservationOrderAndReservation(ctx context.Context, id ReservationId, options ReservationsSummariesListByReservationOrderAndReservationOperationOptions) (result ReservationsSummariesListByReservationOrderAndReservationOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		Path:          fmt.Sprintf("%s/providers/Microsoft.Consumption/reservationSummaries", id.ID()),
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
		Values *[]ReservationSummary `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ReservationsSummariesListByReservationOrderAndReservationComplete retrieves all the results into a single object
func (c ReservationSummariesClient) ReservationsSummariesListByReservationOrderAndReservationComplete(ctx context.Context, id ReservationId, options ReservationsSummariesListByReservationOrderAndReservationOperationOptions) (ReservationsSummariesListByReservationOrderAndReservationCompleteResult, error) {
	return c.ReservationsSummariesListByReservationOrderAndReservationCompleteMatchingPredicate(ctx, id, options, ReservationSummaryOperationPredicate{})
}

// ReservationsSummariesListByReservationOrderAndReservationCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ReservationSummariesClient) ReservationsSummariesListByReservationOrderAndReservationCompleteMatchingPredicate(ctx context.Context, id ReservationId, options ReservationsSummariesListByReservationOrderAndReservationOperationOptions, predicate ReservationSummaryOperationPredicate) (result ReservationsSummariesListByReservationOrderAndReservationCompleteResult, err error) {
	items := make([]ReservationSummary, 0)

	resp, err := c.ReservationsSummariesListByReservationOrderAndReservation(ctx, id, options)
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

	result = ReservationsSummariesListByReservationOrderAndReservationCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}

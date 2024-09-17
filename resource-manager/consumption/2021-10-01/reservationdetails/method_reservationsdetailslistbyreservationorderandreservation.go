package reservationdetails

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ReservationsDetailsListByReservationOrderAndReservationOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]ReservationDetail
}

type ReservationsDetailsListByReservationOrderAndReservationCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []ReservationDetail
}

type ReservationsDetailsListByReservationOrderAndReservationOperationOptions struct {
	Filter *string
}

func DefaultReservationsDetailsListByReservationOrderAndReservationOperationOptions() ReservationsDetailsListByReservationOrderAndReservationOperationOptions {
	return ReservationsDetailsListByReservationOrderAndReservationOperationOptions{}
}

func (o ReservationsDetailsListByReservationOrderAndReservationOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ReservationsDetailsListByReservationOrderAndReservationOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o ReservationsDetailsListByReservationOrderAndReservationOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Filter != nil {
		out.Append("$filter", fmt.Sprintf("%v", *o.Filter))
	}
	return &out
}

type ReservationsDetailsListByReservationOrderAndReservationCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ReservationsDetailsListByReservationOrderAndReservationCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ReservationsDetailsListByReservationOrderAndReservation ...
func (c ReservationDetailsClient) ReservationsDetailsListByReservationOrderAndReservation(ctx context.Context, id ReservationId, options ReservationsDetailsListByReservationOrderAndReservationOperationOptions) (result ReservationsDetailsListByReservationOrderAndReservationOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ReservationsDetailsListByReservationOrderAndReservationCustomPager{},
		Path:          fmt.Sprintf("%s/providers/Microsoft.Consumption/reservationDetails", id.ID()),
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
		Values *[]ReservationDetail `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ReservationsDetailsListByReservationOrderAndReservationComplete retrieves all the results into a single object
func (c ReservationDetailsClient) ReservationsDetailsListByReservationOrderAndReservationComplete(ctx context.Context, id ReservationId, options ReservationsDetailsListByReservationOrderAndReservationOperationOptions) (ReservationsDetailsListByReservationOrderAndReservationCompleteResult, error) {
	return c.ReservationsDetailsListByReservationOrderAndReservationCompleteMatchingPredicate(ctx, id, options, ReservationDetailOperationPredicate{})
}

// ReservationsDetailsListByReservationOrderAndReservationCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ReservationDetailsClient) ReservationsDetailsListByReservationOrderAndReservationCompleteMatchingPredicate(ctx context.Context, id ReservationId, options ReservationsDetailsListByReservationOrderAndReservationOperationOptions, predicate ReservationDetailOperationPredicate) (result ReservationsDetailsListByReservationOrderAndReservationCompleteResult, err error) {
	items := make([]ReservationDetail, 0)

	resp, err := c.ReservationsDetailsListByReservationOrderAndReservation(ctx, id, options)
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

	result = ReservationsDetailsListByReservationOrderAndReservationCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}

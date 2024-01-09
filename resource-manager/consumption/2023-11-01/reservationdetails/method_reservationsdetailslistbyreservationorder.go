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

type ReservationsDetailsListByReservationOrderOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]ReservationDetail
}

type ReservationsDetailsListByReservationOrderCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []ReservationDetail
}

type ReservationsDetailsListByReservationOrderOperationOptions struct {
	Filter *string
}

func DefaultReservationsDetailsListByReservationOrderOperationOptions() ReservationsDetailsListByReservationOrderOperationOptions {
	return ReservationsDetailsListByReservationOrderOperationOptions{}
}

func (o ReservationsDetailsListByReservationOrderOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ReservationsDetailsListByReservationOrderOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	return &out
}

func (o ReservationsDetailsListByReservationOrderOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Filter != nil {
		out.Append("$filter", fmt.Sprintf("%v", *o.Filter))
	}
	return &out
}

// ReservationsDetailsListByReservationOrder ...
func (c ReservationDetailsClient) ReservationsDetailsListByReservationOrder(ctx context.Context, id ReservationOrderId, options ReservationsDetailsListByReservationOrderOperationOptions) (result ReservationsDetailsListByReservationOrderOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		Path:          fmt.Sprintf("%s/providers/Microsoft.Consumption/reservationDetails", id.ID()),
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
		Values *[]ReservationDetail `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ReservationsDetailsListByReservationOrderComplete retrieves all the results into a single object
func (c ReservationDetailsClient) ReservationsDetailsListByReservationOrderComplete(ctx context.Context, id ReservationOrderId, options ReservationsDetailsListByReservationOrderOperationOptions) (ReservationsDetailsListByReservationOrderCompleteResult, error) {
	return c.ReservationsDetailsListByReservationOrderCompleteMatchingPredicate(ctx, id, options, ReservationDetailOperationPredicate{})
}

// ReservationsDetailsListByReservationOrderCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ReservationDetailsClient) ReservationsDetailsListByReservationOrderCompleteMatchingPredicate(ctx context.Context, id ReservationOrderId, options ReservationsDetailsListByReservationOrderOperationOptions, predicate ReservationDetailOperationPredicate) (result ReservationsDetailsListByReservationOrderCompleteResult, err error) {
	items := make([]ReservationDetail, 0)

	resp, err := c.ReservationsDetailsListByReservationOrder(ctx, id, options)
	if err != nil {
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

	result = ReservationsDetailsListByReservationOrderCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}

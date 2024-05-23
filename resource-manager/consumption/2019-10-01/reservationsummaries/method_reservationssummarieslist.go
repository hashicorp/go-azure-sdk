package reservationsummaries

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ReservationsSummariesListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]ReservationSummary
}

type ReservationsSummariesListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []ReservationSummary
}

type ReservationsSummariesListOperationOptions struct {
	EndDate            *string
	Filter             *string
	Grain              *Datagrain
	ReservationId      *string
	ReservationOrderId *string
	StartDate          *string
}

func DefaultReservationsSummariesListOperationOptions() ReservationsSummariesListOperationOptions {
	return ReservationsSummariesListOperationOptions{}
}

func (o ReservationsSummariesListOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ReservationsSummariesListOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	return &out
}

func (o ReservationsSummariesListOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.EndDate != nil {
		out.Append("endDate", fmt.Sprintf("%v", *o.EndDate))
	}
	if o.Filter != nil {
		out.Append("$filter", fmt.Sprintf("%v", *o.Filter))
	}
	if o.Grain != nil {
		out.Append("grain", fmt.Sprintf("%v", *o.Grain))
	}
	if o.ReservationId != nil {
		out.Append("reservationId", fmt.Sprintf("%v", *o.ReservationId))
	}
	if o.ReservationOrderId != nil {
		out.Append("reservationOrderId", fmt.Sprintf("%v", *o.ReservationOrderId))
	}
	if o.StartDate != nil {
		out.Append("startDate", fmt.Sprintf("%v", *o.StartDate))
	}
	return &out
}

// ReservationsSummariesList ...
func (c ReservationSummariesClient) ReservationsSummariesList(ctx context.Context, id commonids.ScopeId, options ReservationsSummariesListOperationOptions) (result ReservationsSummariesListOperationResponse, err error) {
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

// ReservationsSummariesListComplete retrieves all the results into a single object
func (c ReservationSummariesClient) ReservationsSummariesListComplete(ctx context.Context, id commonids.ScopeId, options ReservationsSummariesListOperationOptions) (ReservationsSummariesListCompleteResult, error) {
	return c.ReservationsSummariesListCompleteMatchingPredicate(ctx, id, options, ReservationSummaryOperationPredicate{})
}

// ReservationsSummariesListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ReservationSummariesClient) ReservationsSummariesListCompleteMatchingPredicate(ctx context.Context, id commonids.ScopeId, options ReservationsSummariesListOperationOptions, predicate ReservationSummaryOperationPredicate) (result ReservationsSummariesListCompleteResult, err error) {
	items := make([]ReservationSummary, 0)

	resp, err := c.ReservationsSummariesList(ctx, id, options)
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

	result = ReservationsSummariesListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}

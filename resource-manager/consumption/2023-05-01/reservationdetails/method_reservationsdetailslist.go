package reservationdetails

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

type ReservationsDetailsListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]ReservationDetail
}

type ReservationsDetailsListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []ReservationDetail
}

type ReservationsDetailsListOperationOptions struct {
	EndDate            *string
	Filter             *string
	ReservationId      *string
	ReservationOrderId *string
	StartDate          *string
}

func DefaultReservationsDetailsListOperationOptions() ReservationsDetailsListOperationOptions {
	return ReservationsDetailsListOperationOptions{}
}

func (o ReservationsDetailsListOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ReservationsDetailsListOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	return &out
}

func (o ReservationsDetailsListOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.EndDate != nil {
		out.Append("endDate", fmt.Sprintf("%v", *o.EndDate))
	}
	if o.Filter != nil {
		out.Append("$filter", fmt.Sprintf("%v", *o.Filter))
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

type ReservationsDetailsListCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ReservationsDetailsListCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ReservationsDetailsList ...
func (c ReservationDetailsClient) ReservationsDetailsList(ctx context.Context, id commonids.ScopeId, options ReservationsDetailsListOperationOptions) (result ReservationsDetailsListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		Pager:         &ReservationsDetailsListCustomPager{},
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

// ReservationsDetailsListComplete retrieves all the results into a single object
func (c ReservationDetailsClient) ReservationsDetailsListComplete(ctx context.Context, id commonids.ScopeId, options ReservationsDetailsListOperationOptions) (ReservationsDetailsListCompleteResult, error) {
	return c.ReservationsDetailsListCompleteMatchingPredicate(ctx, id, options, ReservationDetailOperationPredicate{})
}

// ReservationsDetailsListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ReservationDetailsClient) ReservationsDetailsListCompleteMatchingPredicate(ctx context.Context, id commonids.ScopeId, options ReservationsDetailsListOperationOptions, predicate ReservationDetailOperationPredicate) (result ReservationsDetailsListCompleteResult, err error) {
	items := make([]ReservationDetail, 0)

	resp, err := c.ReservationsDetailsList(ctx, id, options)
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

	result = ReservationsDetailsListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}

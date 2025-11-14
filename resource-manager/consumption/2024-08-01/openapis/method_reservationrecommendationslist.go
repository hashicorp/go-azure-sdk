package openapis

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ReservationRecommendationsListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]ReservationRecommendation
}

type ReservationRecommendationsListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []ReservationRecommendation
}

type ReservationRecommendationsListOperationOptions struct {
	Filter *string
}

func DefaultReservationRecommendationsListOperationOptions() ReservationRecommendationsListOperationOptions {
	return ReservationRecommendationsListOperationOptions{}
}

func (o ReservationRecommendationsListOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ReservationRecommendationsListOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o ReservationRecommendationsListOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Filter != nil {
		out.Append("$filter", fmt.Sprintf("%v", *o.Filter))
	}
	return &out
}

type ReservationRecommendationsListCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ReservationRecommendationsListCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ReservationRecommendationsList ...
func (c OpenapisClient) ReservationRecommendationsList(ctx context.Context, id commonids.ScopeId, options ReservationRecommendationsListOperationOptions) (result ReservationRecommendationsListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ReservationRecommendationsListCustomPager{},
		Path:          fmt.Sprintf("%s/providers/Microsoft.Consumption/reservationRecommendations", id.ID()),
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
		Values *[]json.RawMessage `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	temp := make([]ReservationRecommendation, 0)
	if values.Values != nil {
		for i, v := range *values.Values {
			val, err := UnmarshalReservationRecommendationImplementation(v)
			if err != nil {
				err = fmt.Errorf("unmarshalling item %d for ReservationRecommendation (%q): %+v", i, v, err)
				return result, err
			}
			temp = append(temp, val)
		}
	}
	result.Model = &temp

	return
}

// ReservationRecommendationsListComplete retrieves all the results into a single object
func (c OpenapisClient) ReservationRecommendationsListComplete(ctx context.Context, id commonids.ScopeId, options ReservationRecommendationsListOperationOptions) (ReservationRecommendationsListCompleteResult, error) {
	return c.ReservationRecommendationsListCompleteMatchingPredicate(ctx, id, options, ReservationRecommendationOperationPredicate{})
}

// ReservationRecommendationsListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c OpenapisClient) ReservationRecommendationsListCompleteMatchingPredicate(ctx context.Context, id commonids.ScopeId, options ReservationRecommendationsListOperationOptions, predicate ReservationRecommendationOperationPredicate) (result ReservationRecommendationsListCompleteResult, err error) {
	items := make([]ReservationRecommendation, 0)

	resp, err := c.ReservationRecommendationsList(ctx, id, options)
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

	result = ReservationRecommendationsListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}

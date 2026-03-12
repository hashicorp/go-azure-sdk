package openapis

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

type RecommendationsListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]Recommendation
}

type RecommendationsListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []Recommendation
}

type RecommendationsListOperationOptions struct {
	Featured *bool
	Filter   *string
}

func DefaultRecommendationsListOperationOptions() RecommendationsListOperationOptions {
	return RecommendationsListOperationOptions{}
}

func (o RecommendationsListOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o RecommendationsListOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o RecommendationsListOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Featured != nil {
		out.Append("featured", fmt.Sprintf("%v", *o.Featured))
	}
	if o.Filter != nil {
		out.Append("$filter", fmt.Sprintf("%v", *o.Filter))
	}
	return &out
}

type RecommendationsListCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *RecommendationsListCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// RecommendationsList ...
func (c OpenapisClient) RecommendationsList(ctx context.Context, id commonids.SubscriptionId, options RecommendationsListOperationOptions) (result RecommendationsListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &RecommendationsListCustomPager{},
		Path:          fmt.Sprintf("%s/providers/Microsoft.Web/recommendations", id.ID()),
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
		Values *[]Recommendation `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// RecommendationsListComplete retrieves all the results into a single object
func (c OpenapisClient) RecommendationsListComplete(ctx context.Context, id commonids.SubscriptionId, options RecommendationsListOperationOptions) (RecommendationsListCompleteResult, error) {
	return c.RecommendationsListCompleteMatchingPredicate(ctx, id, options, RecommendationOperationPredicate{})
}

// RecommendationsListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c OpenapisClient) RecommendationsListCompleteMatchingPredicate(ctx context.Context, id commonids.SubscriptionId, options RecommendationsListOperationOptions, predicate RecommendationOperationPredicate) (result RecommendationsListCompleteResult, err error) {
	items := make([]Recommendation, 0)

	resp, err := c.RecommendationsList(ctx, id, options)
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

	result = RecommendationsListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}

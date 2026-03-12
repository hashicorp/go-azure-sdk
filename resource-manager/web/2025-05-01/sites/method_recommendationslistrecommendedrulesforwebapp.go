package sites

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

type RecommendationsListRecommendedRulesForWebAppOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]Recommendation
}

type RecommendationsListRecommendedRulesForWebAppCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []Recommendation
}

type RecommendationsListRecommendedRulesForWebAppOperationOptions struct {
	Featured *bool
	Filter   *string
}

func DefaultRecommendationsListRecommendedRulesForWebAppOperationOptions() RecommendationsListRecommendedRulesForWebAppOperationOptions {
	return RecommendationsListRecommendedRulesForWebAppOperationOptions{}
}

func (o RecommendationsListRecommendedRulesForWebAppOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o RecommendationsListRecommendedRulesForWebAppOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o RecommendationsListRecommendedRulesForWebAppOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Featured != nil {
		out.Append("featured", fmt.Sprintf("%v", *o.Featured))
	}
	if o.Filter != nil {
		out.Append("$filter", fmt.Sprintf("%v", *o.Filter))
	}
	return &out
}

type RecommendationsListRecommendedRulesForWebAppCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *RecommendationsListRecommendedRulesForWebAppCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// RecommendationsListRecommendedRulesForWebApp ...
func (c SitesClient) RecommendationsListRecommendedRulesForWebApp(ctx context.Context, id commonids.AppServiceId, options RecommendationsListRecommendedRulesForWebAppOperationOptions) (result RecommendationsListRecommendedRulesForWebAppOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &RecommendationsListRecommendedRulesForWebAppCustomPager{},
		Path:          fmt.Sprintf("%s/recommendations", id.ID()),
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

// RecommendationsListRecommendedRulesForWebAppComplete retrieves all the results into a single object
func (c SitesClient) RecommendationsListRecommendedRulesForWebAppComplete(ctx context.Context, id commonids.AppServiceId, options RecommendationsListRecommendedRulesForWebAppOperationOptions) (RecommendationsListRecommendedRulesForWebAppCompleteResult, error) {
	return c.RecommendationsListRecommendedRulesForWebAppCompleteMatchingPredicate(ctx, id, options, RecommendationOperationPredicate{})
}

// RecommendationsListRecommendedRulesForWebAppCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c SitesClient) RecommendationsListRecommendedRulesForWebAppCompleteMatchingPredicate(ctx context.Context, id commonids.AppServiceId, options RecommendationsListRecommendedRulesForWebAppOperationOptions, predicate RecommendationOperationPredicate) (result RecommendationsListRecommendedRulesForWebAppCompleteResult, err error) {
	items := make([]Recommendation, 0)

	resp, err := c.RecommendationsListRecommendedRulesForWebApp(ctx, id, options)
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

	result = RecommendationsListRecommendedRulesForWebAppCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}

package appserviceenvironmentresources

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

type RecommendationsListRecommendedRulesForHostingEnvironmentOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]Recommendation
}

type RecommendationsListRecommendedRulesForHostingEnvironmentCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []Recommendation
}

type RecommendationsListRecommendedRulesForHostingEnvironmentOperationOptions struct {
	Featured *bool
	Filter   *string
}

func DefaultRecommendationsListRecommendedRulesForHostingEnvironmentOperationOptions() RecommendationsListRecommendedRulesForHostingEnvironmentOperationOptions {
	return RecommendationsListRecommendedRulesForHostingEnvironmentOperationOptions{}
}

func (o RecommendationsListRecommendedRulesForHostingEnvironmentOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o RecommendationsListRecommendedRulesForHostingEnvironmentOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o RecommendationsListRecommendedRulesForHostingEnvironmentOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Featured != nil {
		out.Append("featured", fmt.Sprintf("%v", *o.Featured))
	}
	if o.Filter != nil {
		out.Append("$filter", fmt.Sprintf("%v", *o.Filter))
	}
	return &out
}

type RecommendationsListRecommendedRulesForHostingEnvironmentCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *RecommendationsListRecommendedRulesForHostingEnvironmentCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// RecommendationsListRecommendedRulesForHostingEnvironment ...
func (c AppServiceEnvironmentResourcesClient) RecommendationsListRecommendedRulesForHostingEnvironment(ctx context.Context, id commonids.AppServiceEnvironmentId, options RecommendationsListRecommendedRulesForHostingEnvironmentOperationOptions) (result RecommendationsListRecommendedRulesForHostingEnvironmentOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &RecommendationsListRecommendedRulesForHostingEnvironmentCustomPager{},
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

// RecommendationsListRecommendedRulesForHostingEnvironmentComplete retrieves all the results into a single object
func (c AppServiceEnvironmentResourcesClient) RecommendationsListRecommendedRulesForHostingEnvironmentComplete(ctx context.Context, id commonids.AppServiceEnvironmentId, options RecommendationsListRecommendedRulesForHostingEnvironmentOperationOptions) (RecommendationsListRecommendedRulesForHostingEnvironmentCompleteResult, error) {
	return c.RecommendationsListRecommendedRulesForHostingEnvironmentCompleteMatchingPredicate(ctx, id, options, RecommendationOperationPredicate{})
}

// RecommendationsListRecommendedRulesForHostingEnvironmentCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c AppServiceEnvironmentResourcesClient) RecommendationsListRecommendedRulesForHostingEnvironmentCompleteMatchingPredicate(ctx context.Context, id commonids.AppServiceEnvironmentId, options RecommendationsListRecommendedRulesForHostingEnvironmentOperationOptions, predicate RecommendationOperationPredicate) (result RecommendationsListRecommendedRulesForHostingEnvironmentCompleteResult, err error) {
	items := make([]Recommendation, 0)

	resp, err := c.RecommendationsListRecommendedRulesForHostingEnvironment(ctx, id, options)
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

	result = RecommendationsListRecommendedRulesForHostingEnvironmentCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}

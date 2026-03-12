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

type RecommendationsListHistoryForHostingEnvironmentOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]Recommendation
}

type RecommendationsListHistoryForHostingEnvironmentCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []Recommendation
}

type RecommendationsListHistoryForHostingEnvironmentOperationOptions struct {
	ExpiredOnly *bool
	Filter      *string
}

func DefaultRecommendationsListHistoryForHostingEnvironmentOperationOptions() RecommendationsListHistoryForHostingEnvironmentOperationOptions {
	return RecommendationsListHistoryForHostingEnvironmentOperationOptions{}
}

func (o RecommendationsListHistoryForHostingEnvironmentOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o RecommendationsListHistoryForHostingEnvironmentOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o RecommendationsListHistoryForHostingEnvironmentOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.ExpiredOnly != nil {
		out.Append("expiredOnly", fmt.Sprintf("%v", *o.ExpiredOnly))
	}
	if o.Filter != nil {
		out.Append("$filter", fmt.Sprintf("%v", *o.Filter))
	}
	return &out
}

type RecommendationsListHistoryForHostingEnvironmentCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *RecommendationsListHistoryForHostingEnvironmentCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// RecommendationsListHistoryForHostingEnvironment ...
func (c AppServiceEnvironmentResourcesClient) RecommendationsListHistoryForHostingEnvironment(ctx context.Context, id commonids.AppServiceEnvironmentId, options RecommendationsListHistoryForHostingEnvironmentOperationOptions) (result RecommendationsListHistoryForHostingEnvironmentOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &RecommendationsListHistoryForHostingEnvironmentCustomPager{},
		Path:          fmt.Sprintf("%s/recommendationHistory", id.ID()),
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

// RecommendationsListHistoryForHostingEnvironmentComplete retrieves all the results into a single object
func (c AppServiceEnvironmentResourcesClient) RecommendationsListHistoryForHostingEnvironmentComplete(ctx context.Context, id commonids.AppServiceEnvironmentId, options RecommendationsListHistoryForHostingEnvironmentOperationOptions) (RecommendationsListHistoryForHostingEnvironmentCompleteResult, error) {
	return c.RecommendationsListHistoryForHostingEnvironmentCompleteMatchingPredicate(ctx, id, options, RecommendationOperationPredicate{})
}

// RecommendationsListHistoryForHostingEnvironmentCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c AppServiceEnvironmentResourcesClient) RecommendationsListHistoryForHostingEnvironmentCompleteMatchingPredicate(ctx context.Context, id commonids.AppServiceEnvironmentId, options RecommendationsListHistoryForHostingEnvironmentOperationOptions, predicate RecommendationOperationPredicate) (result RecommendationsListHistoryForHostingEnvironmentCompleteResult, err error) {
	items := make([]Recommendation, 0)

	resp, err := c.RecommendationsListHistoryForHostingEnvironment(ctx, id, options)
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

	result = RecommendationsListHistoryForHostingEnvironmentCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}

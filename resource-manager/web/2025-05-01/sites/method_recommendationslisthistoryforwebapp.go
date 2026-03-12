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

type RecommendationsListHistoryForWebAppOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]Recommendation
}

type RecommendationsListHistoryForWebAppCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []Recommendation
}

type RecommendationsListHistoryForWebAppOperationOptions struct {
	ExpiredOnly *bool
	Filter      *string
}

func DefaultRecommendationsListHistoryForWebAppOperationOptions() RecommendationsListHistoryForWebAppOperationOptions {
	return RecommendationsListHistoryForWebAppOperationOptions{}
}

func (o RecommendationsListHistoryForWebAppOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o RecommendationsListHistoryForWebAppOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o RecommendationsListHistoryForWebAppOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.ExpiredOnly != nil {
		out.Append("expiredOnly", fmt.Sprintf("%v", *o.ExpiredOnly))
	}
	if o.Filter != nil {
		out.Append("$filter", fmt.Sprintf("%v", *o.Filter))
	}
	return &out
}

type RecommendationsListHistoryForWebAppCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *RecommendationsListHistoryForWebAppCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// RecommendationsListHistoryForWebApp ...
func (c SitesClient) RecommendationsListHistoryForWebApp(ctx context.Context, id commonids.AppServiceId, options RecommendationsListHistoryForWebAppOperationOptions) (result RecommendationsListHistoryForWebAppOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &RecommendationsListHistoryForWebAppCustomPager{},
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

// RecommendationsListHistoryForWebAppComplete retrieves all the results into a single object
func (c SitesClient) RecommendationsListHistoryForWebAppComplete(ctx context.Context, id commonids.AppServiceId, options RecommendationsListHistoryForWebAppOperationOptions) (RecommendationsListHistoryForWebAppCompleteResult, error) {
	return c.RecommendationsListHistoryForWebAppCompleteMatchingPredicate(ctx, id, options, RecommendationOperationPredicate{})
}

// RecommendationsListHistoryForWebAppCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c SitesClient) RecommendationsListHistoryForWebAppCompleteMatchingPredicate(ctx context.Context, id commonids.AppServiceId, options RecommendationsListHistoryForWebAppOperationOptions, predicate RecommendationOperationPredicate) (result RecommendationsListHistoryForWebAppCompleteResult, err error) {
	items := make([]Recommendation, 0)

	resp, err := c.RecommendationsListHistoryForWebApp(ctx, id, options)
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

	result = RecommendationsListHistoryForWebAppCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}

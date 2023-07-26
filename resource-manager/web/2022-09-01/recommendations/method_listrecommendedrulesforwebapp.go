package recommendations

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListRecommendedRulesForWebAppOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]Recommendation
}

type ListRecommendedRulesForWebAppCompleteResult struct {
	Items []Recommendation
}

type ListRecommendedRulesForWebAppOperationOptions struct {
	Featured *bool
	Filter   *string
}

func DefaultListRecommendedRulesForWebAppOperationOptions() ListRecommendedRulesForWebAppOperationOptions {
	return ListRecommendedRulesForWebAppOperationOptions{}
}

func (o ListRecommendedRulesForWebAppOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListRecommendedRulesForWebAppOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	return &out
}

func (o ListRecommendedRulesForWebAppOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Featured != nil {
		out.Append("featured", fmt.Sprintf("%v", *o.Featured))
	}
	if o.Filter != nil {
		out.Append("$filter", fmt.Sprintf("%v", *o.Filter))
	}
	return &out
}

// ListRecommendedRulesForWebApp ...
func (c RecommendationsClient) ListRecommendedRulesForWebApp(ctx context.Context, id SiteId, options ListRecommendedRulesForWebAppOperationOptions) (result ListRecommendedRulesForWebAppOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		Path:          fmt.Sprintf("%s/recommendations", id.ID()),
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
		Values *[]Recommendation `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListRecommendedRulesForWebAppComplete retrieves all the results into a single object
func (c RecommendationsClient) ListRecommendedRulesForWebAppComplete(ctx context.Context, id SiteId, options ListRecommendedRulesForWebAppOperationOptions) (ListRecommendedRulesForWebAppCompleteResult, error) {
	return c.ListRecommendedRulesForWebAppCompleteMatchingPredicate(ctx, id, options, RecommendationOperationPredicate{})
}

// ListRecommendedRulesForWebAppCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c RecommendationsClient) ListRecommendedRulesForWebAppCompleteMatchingPredicate(ctx context.Context, id SiteId, options ListRecommendedRulesForWebAppOperationOptions, predicate RecommendationOperationPredicate) (result ListRecommendedRulesForWebAppCompleteResult, err error) {
	items := make([]Recommendation, 0)

	resp, err := c.ListRecommendedRulesForWebApp(ctx, id, options)
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

	result = ListRecommendedRulesForWebAppCompleteResult{
		Items: items,
	}
	return
}

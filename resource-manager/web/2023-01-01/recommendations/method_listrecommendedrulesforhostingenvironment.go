package recommendations

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

type ListRecommendedRulesForHostingEnvironmentOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]Recommendation
}

type ListRecommendedRulesForHostingEnvironmentCompleteResult struct {
	Items []Recommendation
}

type ListRecommendedRulesForHostingEnvironmentOperationOptions struct {
	Featured *bool
	Filter   *string
}

func DefaultListRecommendedRulesForHostingEnvironmentOperationOptions() ListRecommendedRulesForHostingEnvironmentOperationOptions {
	return ListRecommendedRulesForHostingEnvironmentOperationOptions{}
}

func (o ListRecommendedRulesForHostingEnvironmentOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListRecommendedRulesForHostingEnvironmentOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	return &out
}

func (o ListRecommendedRulesForHostingEnvironmentOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Featured != nil {
		out.Append("featured", fmt.Sprintf("%v", *o.Featured))
	}
	if o.Filter != nil {
		out.Append("$filter", fmt.Sprintf("%v", *o.Filter))
	}
	return &out
}

// ListRecommendedRulesForHostingEnvironment ...
func (c RecommendationsClient) ListRecommendedRulesForHostingEnvironment(ctx context.Context, id commonids.AppServiceEnvironmentId, options ListRecommendedRulesForHostingEnvironmentOperationOptions) (result ListRecommendedRulesForHostingEnvironmentOperationResponse, err error) {
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

// ListRecommendedRulesForHostingEnvironmentComplete retrieves all the results into a single object
func (c RecommendationsClient) ListRecommendedRulesForHostingEnvironmentComplete(ctx context.Context, id commonids.AppServiceEnvironmentId, options ListRecommendedRulesForHostingEnvironmentOperationOptions) (ListRecommendedRulesForHostingEnvironmentCompleteResult, error) {
	return c.ListRecommendedRulesForHostingEnvironmentCompleteMatchingPredicate(ctx, id, options, RecommendationOperationPredicate{})
}

// ListRecommendedRulesForHostingEnvironmentCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c RecommendationsClient) ListRecommendedRulesForHostingEnvironmentCompleteMatchingPredicate(ctx context.Context, id commonids.AppServiceEnvironmentId, options ListRecommendedRulesForHostingEnvironmentOperationOptions, predicate RecommendationOperationPredicate) (result ListRecommendedRulesForHostingEnvironmentCompleteResult, err error) {
	items := make([]Recommendation, 0)

	resp, err := c.ListRecommendedRulesForHostingEnvironment(ctx, id, options)
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

	result = ListRecommendedRulesForHostingEnvironmentCompleteResult{
		Items: items,
	}
	return
}

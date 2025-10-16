package tuningoptions

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListRecommendationsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]ObjectRecommendation
}

type ListRecommendationsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []ObjectRecommendation
}

type ListRecommendationsOperationOptions struct {
	RecommendationType *RecommendationType
}

func DefaultListRecommendationsOperationOptions() ListRecommendationsOperationOptions {
	return ListRecommendationsOperationOptions{}
}

func (o ListRecommendationsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListRecommendationsOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o ListRecommendationsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.RecommendationType != nil {
		out.Append("recommendationType", fmt.Sprintf("%v", *o.RecommendationType))
	}
	return &out
}

type ListRecommendationsCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ListRecommendationsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListRecommendations ...
func (c TuningOptionsClient) ListRecommendations(ctx context.Context, id TuningOptionId, options ListRecommendationsOperationOptions) (result ListRecommendationsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListRecommendationsCustomPager{},
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
		Values *[]ObjectRecommendation `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListRecommendationsComplete retrieves all the results into a single object
func (c TuningOptionsClient) ListRecommendationsComplete(ctx context.Context, id TuningOptionId, options ListRecommendationsOperationOptions) (ListRecommendationsCompleteResult, error) {
	return c.ListRecommendationsCompleteMatchingPredicate(ctx, id, options, ObjectRecommendationOperationPredicate{})
}

// ListRecommendationsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c TuningOptionsClient) ListRecommendationsCompleteMatchingPredicate(ctx context.Context, id TuningOptionId, options ListRecommendationsOperationOptions, predicate ObjectRecommendationOperationPredicate) (result ListRecommendationsCompleteResult, err error) {
	items := make([]ObjectRecommendation, 0)

	resp, err := c.ListRecommendations(ctx, id, options)
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

	result = ListRecommendationsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}

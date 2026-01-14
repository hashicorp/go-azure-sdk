package tuningoptionsoperationgroup

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TuningOptionsListRecommendationsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]ObjectRecommendation
}

type TuningOptionsListRecommendationsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []ObjectRecommendation
}

type TuningOptionsListRecommendationsOperationOptions struct {
	RecommendationType *RecommendationTypeParameterEnum
}

func DefaultTuningOptionsListRecommendationsOperationOptions() TuningOptionsListRecommendationsOperationOptions {
	return TuningOptionsListRecommendationsOperationOptions{}
}

func (o TuningOptionsListRecommendationsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o TuningOptionsListRecommendationsOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o TuningOptionsListRecommendationsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.RecommendationType != nil {
		out.Append("recommendationType", fmt.Sprintf("%v", *o.RecommendationType))
	}
	return &out
}

type TuningOptionsListRecommendationsCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *TuningOptionsListRecommendationsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// TuningOptionsListRecommendations ...
func (c TuningOptionsOperationGroupClient) TuningOptionsListRecommendations(ctx context.Context, id TuningOptionId, options TuningOptionsListRecommendationsOperationOptions) (result TuningOptionsListRecommendationsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &TuningOptionsListRecommendationsCustomPager{},
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

// TuningOptionsListRecommendationsComplete retrieves all the results into a single object
func (c TuningOptionsOperationGroupClient) TuningOptionsListRecommendationsComplete(ctx context.Context, id TuningOptionId, options TuningOptionsListRecommendationsOperationOptions) (TuningOptionsListRecommendationsCompleteResult, error) {
	return c.TuningOptionsListRecommendationsCompleteMatchingPredicate(ctx, id, options, ObjectRecommendationOperationPredicate{})
}

// TuningOptionsListRecommendationsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c TuningOptionsOperationGroupClient) TuningOptionsListRecommendationsCompleteMatchingPredicate(ctx context.Context, id TuningOptionId, options TuningOptionsListRecommendationsOperationOptions, predicate ObjectRecommendationOperationPredicate) (result TuningOptionsListRecommendationsCompleteResult, err error) {
	items := make([]ObjectRecommendation, 0)

	resp, err := c.TuningOptionsListRecommendations(ctx, id, options)
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

	result = TuningOptionsListRecommendationsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}

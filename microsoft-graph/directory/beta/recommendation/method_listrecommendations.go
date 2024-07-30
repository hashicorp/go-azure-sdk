package recommendation

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListRecommendationsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.Recommendation
}

type ListRecommendationsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.Recommendation
}

type ListRecommendationsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListRecommendationsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListRecommendations ...
func (c RecommendationClient) ListRecommendations(ctx context.Context) (result ListRecommendationsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListRecommendationsCustomPager{},
		Path:       "/directory/recommendations",
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
		Values *[]beta.Recommendation `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListRecommendationsComplete retrieves all the results into a single object
func (c RecommendationClient) ListRecommendationsComplete(ctx context.Context) (ListRecommendationsCompleteResult, error) {
	return c.ListRecommendationsCompleteMatchingPredicate(ctx, RecommendationOperationPredicate{})
}

// ListRecommendationsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c RecommendationClient) ListRecommendationsCompleteMatchingPredicate(ctx context.Context, predicate RecommendationOperationPredicate) (result ListRecommendationsCompleteResult, err error) {
	items := make([]beta.Recommendation, 0)

	resp, err := c.ListRecommendations(ctx)
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

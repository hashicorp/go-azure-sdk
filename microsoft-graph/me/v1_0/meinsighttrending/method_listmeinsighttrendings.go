package meinsighttrending

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common/v1_0/models"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListMeInsightTrendingsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.TrendingCollectionResponse
}

type ListMeInsightTrendingsCompleteResult struct {
	Items []models.TrendingCollectionResponse
}

// ListMeInsightTrendings ...
func (c MeInsightTrendingClient) ListMeInsightTrendings(ctx context.Context) (result ListMeInsightTrendingsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       "/me/insights/trending",
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
		Values *[]models.TrendingCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListMeInsightTrendingsComplete retrieves all the results into a single object
func (c MeInsightTrendingClient) ListMeInsightTrendingsComplete(ctx context.Context) (ListMeInsightTrendingsCompleteResult, error) {
	return c.ListMeInsightTrendingsCompleteMatchingPredicate(ctx, models.TrendingCollectionResponseOperationPredicate{})
}

// ListMeInsightTrendingsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeInsightTrendingClient) ListMeInsightTrendingsCompleteMatchingPredicate(ctx context.Context, predicate models.TrendingCollectionResponseOperationPredicate) (result ListMeInsightTrendingsCompleteResult, err error) {
	items := make([]models.TrendingCollectionResponse, 0)

	resp, err := c.ListMeInsightTrendings(ctx)
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

	result = ListMeInsightTrendingsCompleteResult{
		Items: items,
	}
	return
}

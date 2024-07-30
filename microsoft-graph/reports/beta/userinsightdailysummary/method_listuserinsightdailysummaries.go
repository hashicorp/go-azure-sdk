package userinsightdailysummary

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

type ListUserInsightDailySummariesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.InsightSummary
}

type ListUserInsightDailySummariesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.InsightSummary
}

type ListUserInsightDailySummariesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListUserInsightDailySummariesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListUserInsightDailySummaries ...
func (c UserInsightDailySummaryClient) ListUserInsightDailySummaries(ctx context.Context) (result ListUserInsightDailySummariesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListUserInsightDailySummariesCustomPager{},
		Path:       "/reports/userInsights/daily/summary",
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
		Values *[]beta.InsightSummary `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserInsightDailySummariesComplete retrieves all the results into a single object
func (c UserInsightDailySummaryClient) ListUserInsightDailySummariesComplete(ctx context.Context) (ListUserInsightDailySummariesCompleteResult, error) {
	return c.ListUserInsightDailySummariesCompleteMatchingPredicate(ctx, InsightSummaryOperationPredicate{})
}

// ListUserInsightDailySummariesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserInsightDailySummaryClient) ListUserInsightDailySummariesCompleteMatchingPredicate(ctx context.Context, predicate InsightSummaryOperationPredicate) (result ListUserInsightDailySummariesCompleteResult, err error) {
	items := make([]beta.InsightSummary, 0)

	resp, err := c.ListUserInsightDailySummaries(ctx)
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

	result = ListUserInsightDailySummariesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}

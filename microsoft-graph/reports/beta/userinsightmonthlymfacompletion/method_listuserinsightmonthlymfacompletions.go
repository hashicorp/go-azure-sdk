package userinsightmonthlymfacompletion

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

type ListUserInsightMonthlyMfaCompletionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.MfaCompletionMetric
}

type ListUserInsightMonthlyMfaCompletionsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.MfaCompletionMetric
}

type ListUserInsightMonthlyMfaCompletionsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListUserInsightMonthlyMfaCompletionsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListUserInsightMonthlyMfaCompletions ...
func (c UserInsightMonthlyMfaCompletionClient) ListUserInsightMonthlyMfaCompletions(ctx context.Context) (result ListUserInsightMonthlyMfaCompletionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListUserInsightMonthlyMfaCompletionsCustomPager{},
		Path:       "/reports/userInsights/monthly/mfaCompletions",
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
		Values *[]beta.MfaCompletionMetric `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserInsightMonthlyMfaCompletionsComplete retrieves all the results into a single object
func (c UserInsightMonthlyMfaCompletionClient) ListUserInsightMonthlyMfaCompletionsComplete(ctx context.Context) (ListUserInsightMonthlyMfaCompletionsCompleteResult, error) {
	return c.ListUserInsightMonthlyMfaCompletionsCompleteMatchingPredicate(ctx, MfaCompletionMetricOperationPredicate{})
}

// ListUserInsightMonthlyMfaCompletionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserInsightMonthlyMfaCompletionClient) ListUserInsightMonthlyMfaCompletionsCompleteMatchingPredicate(ctx context.Context, predicate MfaCompletionMetricOperationPredicate) (result ListUserInsightMonthlyMfaCompletionsCompleteResult, err error) {
	items := make([]beta.MfaCompletionMetric, 0)

	resp, err := c.ListUserInsightMonthlyMfaCompletions(ctx)
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

	result = ListUserInsightMonthlyMfaCompletionsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}

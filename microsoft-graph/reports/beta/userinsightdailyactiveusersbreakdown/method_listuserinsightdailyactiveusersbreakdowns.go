package userinsightdailyactiveusersbreakdown

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

type ListUserInsightDailyActiveUsersBreakdownsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.ActiveUsersBreakdownMetric
}

type ListUserInsightDailyActiveUsersBreakdownsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.ActiveUsersBreakdownMetric
}

type ListUserInsightDailyActiveUsersBreakdownsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListUserInsightDailyActiveUsersBreakdownsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListUserInsightDailyActiveUsersBreakdowns ...
func (c UserInsightDailyActiveUsersBreakdownClient) ListUserInsightDailyActiveUsersBreakdowns(ctx context.Context) (result ListUserInsightDailyActiveUsersBreakdownsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListUserInsightDailyActiveUsersBreakdownsCustomPager{},
		Path:       "/reports/userInsights/daily/activeUsersBreakdown",
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
		Values *[]beta.ActiveUsersBreakdownMetric `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserInsightDailyActiveUsersBreakdownsComplete retrieves all the results into a single object
func (c UserInsightDailyActiveUsersBreakdownClient) ListUserInsightDailyActiveUsersBreakdownsComplete(ctx context.Context) (ListUserInsightDailyActiveUsersBreakdownsCompleteResult, error) {
	return c.ListUserInsightDailyActiveUsersBreakdownsCompleteMatchingPredicate(ctx, ActiveUsersBreakdownMetricOperationPredicate{})
}

// ListUserInsightDailyActiveUsersBreakdownsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserInsightDailyActiveUsersBreakdownClient) ListUserInsightDailyActiveUsersBreakdownsCompleteMatchingPredicate(ctx context.Context, predicate ActiveUsersBreakdownMetricOperationPredicate) (result ListUserInsightDailyActiveUsersBreakdownsCompleteResult, err error) {
	items := make([]beta.ActiveUsersBreakdownMetric, 0)

	resp, err := c.ListUserInsightDailyActiveUsersBreakdowns(ctx)
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

	result = ListUserInsightDailyActiveUsersBreakdownsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}

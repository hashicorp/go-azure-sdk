package userinsightmonthlyactiveusersbreakdown

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

type ListUserInsightMonthlyActiveUsersBreakdownsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.ActiveUsersBreakdownMetric
}

type ListUserInsightMonthlyActiveUsersBreakdownsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.ActiveUsersBreakdownMetric
}

type ListUserInsightMonthlyActiveUsersBreakdownsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListUserInsightMonthlyActiveUsersBreakdownsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListUserInsightMonthlyActiveUsersBreakdowns ...
func (c UserInsightMonthlyActiveUsersBreakdownClient) ListUserInsightMonthlyActiveUsersBreakdowns(ctx context.Context) (result ListUserInsightMonthlyActiveUsersBreakdownsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListUserInsightMonthlyActiveUsersBreakdownsCustomPager{},
		Path:       "/reports/userInsights/monthly/activeUsersBreakdown",
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

// ListUserInsightMonthlyActiveUsersBreakdownsComplete retrieves all the results into a single object
func (c UserInsightMonthlyActiveUsersBreakdownClient) ListUserInsightMonthlyActiveUsersBreakdownsComplete(ctx context.Context) (ListUserInsightMonthlyActiveUsersBreakdownsCompleteResult, error) {
	return c.ListUserInsightMonthlyActiveUsersBreakdownsCompleteMatchingPredicate(ctx, ActiveUsersBreakdownMetricOperationPredicate{})
}

// ListUserInsightMonthlyActiveUsersBreakdownsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserInsightMonthlyActiveUsersBreakdownClient) ListUserInsightMonthlyActiveUsersBreakdownsCompleteMatchingPredicate(ctx context.Context, predicate ActiveUsersBreakdownMetricOperationPredicate) (result ListUserInsightMonthlyActiveUsersBreakdownsCompleteResult, err error) {
	items := make([]beta.ActiveUsersBreakdownMetric, 0)

	resp, err := c.ListUserInsightMonthlyActiveUsersBreakdowns(ctx)
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

	result = ListUserInsightMonthlyActiveUsersBreakdownsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}

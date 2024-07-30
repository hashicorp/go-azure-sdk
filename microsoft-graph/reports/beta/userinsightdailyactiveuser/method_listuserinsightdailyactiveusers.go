package userinsightdailyactiveuser

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

type ListUserInsightDailyActiveUsersOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.ActiveUsersMetric
}

type ListUserInsightDailyActiveUsersCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.ActiveUsersMetric
}

type ListUserInsightDailyActiveUsersCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListUserInsightDailyActiveUsersCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListUserInsightDailyActiveUsers ...
func (c UserInsightDailyActiveUserClient) ListUserInsightDailyActiveUsers(ctx context.Context) (result ListUserInsightDailyActiveUsersOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListUserInsightDailyActiveUsersCustomPager{},
		Path:       "/reports/userInsights/daily/activeUsers",
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
		Values *[]beta.ActiveUsersMetric `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserInsightDailyActiveUsersComplete retrieves all the results into a single object
func (c UserInsightDailyActiveUserClient) ListUserInsightDailyActiveUsersComplete(ctx context.Context) (ListUserInsightDailyActiveUsersCompleteResult, error) {
	return c.ListUserInsightDailyActiveUsersCompleteMatchingPredicate(ctx, ActiveUsersMetricOperationPredicate{})
}

// ListUserInsightDailyActiveUsersCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserInsightDailyActiveUserClient) ListUserInsightDailyActiveUsersCompleteMatchingPredicate(ctx context.Context, predicate ActiveUsersMetricOperationPredicate) (result ListUserInsightDailyActiveUsersCompleteResult, err error) {
	items := make([]beta.ActiveUsersMetric, 0)

	resp, err := c.ListUserInsightDailyActiveUsers(ctx)
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

	result = ListUserInsightDailyActiveUsersCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}

package userinsightmonthlyactiveuser

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

type ListUserInsightMonthlyActiveUsersOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.ActiveUsersMetric
}

type ListUserInsightMonthlyActiveUsersCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.ActiveUsersMetric
}

type ListUserInsightMonthlyActiveUsersCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListUserInsightMonthlyActiveUsersCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListUserInsightMonthlyActiveUsers ...
func (c UserInsightMonthlyActiveUserClient) ListUserInsightMonthlyActiveUsers(ctx context.Context) (result ListUserInsightMonthlyActiveUsersOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListUserInsightMonthlyActiveUsersCustomPager{},
		Path:       "/reports/userInsights/monthly/activeUsers",
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

// ListUserInsightMonthlyActiveUsersComplete retrieves all the results into a single object
func (c UserInsightMonthlyActiveUserClient) ListUserInsightMonthlyActiveUsersComplete(ctx context.Context) (ListUserInsightMonthlyActiveUsersCompleteResult, error) {
	return c.ListUserInsightMonthlyActiveUsersCompleteMatchingPredicate(ctx, ActiveUsersMetricOperationPredicate{})
}

// ListUserInsightMonthlyActiveUsersCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserInsightMonthlyActiveUserClient) ListUserInsightMonthlyActiveUsersCompleteMatchingPredicate(ctx context.Context, predicate ActiveUsersMetricOperationPredicate) (result ListUserInsightMonthlyActiveUsersCompleteResult, err error) {
	items := make([]beta.ActiveUsersMetric, 0)

	resp, err := c.ListUserInsightMonthlyActiveUsers(ctx)
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

	result = ListUserInsightMonthlyActiveUsersCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}

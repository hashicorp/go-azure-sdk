package userinsightdailyinactiveuser

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

type ListUserInsightDailyInactiveUsersOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DailyInactiveUsersMetric
}

type ListUserInsightDailyInactiveUsersCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DailyInactiveUsersMetric
}

type ListUserInsightDailyInactiveUsersCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListUserInsightDailyInactiveUsersCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListUserInsightDailyInactiveUsers ...
func (c UserInsightDailyInactiveUserClient) ListUserInsightDailyInactiveUsers(ctx context.Context) (result ListUserInsightDailyInactiveUsersOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListUserInsightDailyInactiveUsersCustomPager{},
		Path:       "/reports/userInsights/daily/inactiveUsers",
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
		Values *[]beta.DailyInactiveUsersMetric `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserInsightDailyInactiveUsersComplete retrieves all the results into a single object
func (c UserInsightDailyInactiveUserClient) ListUserInsightDailyInactiveUsersComplete(ctx context.Context) (ListUserInsightDailyInactiveUsersCompleteResult, error) {
	return c.ListUserInsightDailyInactiveUsersCompleteMatchingPredicate(ctx, DailyInactiveUsersMetricOperationPredicate{})
}

// ListUserInsightDailyInactiveUsersCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserInsightDailyInactiveUserClient) ListUserInsightDailyInactiveUsersCompleteMatchingPredicate(ctx context.Context, predicate DailyInactiveUsersMetricOperationPredicate) (result ListUserInsightDailyInactiveUsersCompleteResult, err error) {
	items := make([]beta.DailyInactiveUsersMetric, 0)

	resp, err := c.ListUserInsightDailyInactiveUsers(ctx)
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

	result = ListUserInsightDailyInactiveUsersCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}

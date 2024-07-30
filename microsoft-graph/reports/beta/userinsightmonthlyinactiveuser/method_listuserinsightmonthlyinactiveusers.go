package userinsightmonthlyinactiveuser

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

type ListUserInsightMonthlyInactiveUsersOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.MonthlyInactiveUsersMetric
}

type ListUserInsightMonthlyInactiveUsersCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.MonthlyInactiveUsersMetric
}

type ListUserInsightMonthlyInactiveUsersCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListUserInsightMonthlyInactiveUsersCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListUserInsightMonthlyInactiveUsers ...
func (c UserInsightMonthlyInactiveUserClient) ListUserInsightMonthlyInactiveUsers(ctx context.Context) (result ListUserInsightMonthlyInactiveUsersOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListUserInsightMonthlyInactiveUsersCustomPager{},
		Path:       "/reports/userInsights/monthly/inactiveUsers",
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
		Values *[]beta.MonthlyInactiveUsersMetric `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserInsightMonthlyInactiveUsersComplete retrieves all the results into a single object
func (c UserInsightMonthlyInactiveUserClient) ListUserInsightMonthlyInactiveUsersComplete(ctx context.Context) (ListUserInsightMonthlyInactiveUsersCompleteResult, error) {
	return c.ListUserInsightMonthlyInactiveUsersCompleteMatchingPredicate(ctx, MonthlyInactiveUsersMetricOperationPredicate{})
}

// ListUserInsightMonthlyInactiveUsersCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserInsightMonthlyInactiveUserClient) ListUserInsightMonthlyInactiveUsersCompleteMatchingPredicate(ctx context.Context, predicate MonthlyInactiveUsersMetricOperationPredicate) (result ListUserInsightMonthlyInactiveUsersCompleteResult, err error) {
	items := make([]beta.MonthlyInactiveUsersMetric, 0)

	resp, err := c.ListUserInsightMonthlyInactiveUsers(ctx)
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

	result = ListUserInsightMonthlyInactiveUsersCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}

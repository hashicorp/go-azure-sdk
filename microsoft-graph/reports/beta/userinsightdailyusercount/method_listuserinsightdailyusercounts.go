package userinsightdailyusercount

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

type ListUserInsightDailyUserCountsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.UserCountMetric
}

type ListUserInsightDailyUserCountsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.UserCountMetric
}

type ListUserInsightDailyUserCountsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListUserInsightDailyUserCountsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListUserInsightDailyUserCounts ...
func (c UserInsightDailyUserCountClient) ListUserInsightDailyUserCounts(ctx context.Context) (result ListUserInsightDailyUserCountsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListUserInsightDailyUserCountsCustomPager{},
		Path:       "/reports/userInsights/daily/userCount",
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
		Values *[]beta.UserCountMetric `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserInsightDailyUserCountsComplete retrieves all the results into a single object
func (c UserInsightDailyUserCountClient) ListUserInsightDailyUserCountsComplete(ctx context.Context) (ListUserInsightDailyUserCountsCompleteResult, error) {
	return c.ListUserInsightDailyUserCountsCompleteMatchingPredicate(ctx, UserCountMetricOperationPredicate{})
}

// ListUserInsightDailyUserCountsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserInsightDailyUserCountClient) ListUserInsightDailyUserCountsCompleteMatchingPredicate(ctx context.Context, predicate UserCountMetricOperationPredicate) (result ListUserInsightDailyUserCountsCompleteResult, err error) {
	items := make([]beta.UserCountMetric, 0)

	resp, err := c.ListUserInsightDailyUserCounts(ctx)
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

	result = ListUserInsightDailyUserCountsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}

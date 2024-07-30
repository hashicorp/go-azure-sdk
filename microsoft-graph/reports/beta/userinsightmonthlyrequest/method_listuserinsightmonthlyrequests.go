package userinsightmonthlyrequest

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

type ListUserInsightMonthlyRequestsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.UserRequestsMetric
}

type ListUserInsightMonthlyRequestsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.UserRequestsMetric
}

type ListUserInsightMonthlyRequestsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListUserInsightMonthlyRequestsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListUserInsightMonthlyRequests ...
func (c UserInsightMonthlyRequestClient) ListUserInsightMonthlyRequests(ctx context.Context) (result ListUserInsightMonthlyRequestsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListUserInsightMonthlyRequestsCustomPager{},
		Path:       "/reports/userInsights/monthly/requests",
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
		Values *[]beta.UserRequestsMetric `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserInsightMonthlyRequestsComplete retrieves all the results into a single object
func (c UserInsightMonthlyRequestClient) ListUserInsightMonthlyRequestsComplete(ctx context.Context) (ListUserInsightMonthlyRequestsCompleteResult, error) {
	return c.ListUserInsightMonthlyRequestsCompleteMatchingPredicate(ctx, UserRequestsMetricOperationPredicate{})
}

// ListUserInsightMonthlyRequestsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserInsightMonthlyRequestClient) ListUserInsightMonthlyRequestsCompleteMatchingPredicate(ctx context.Context, predicate UserRequestsMetricOperationPredicate) (result ListUserInsightMonthlyRequestsCompleteResult, err error) {
	items := make([]beta.UserRequestsMetric, 0)

	resp, err := c.ListUserInsightMonthlyRequests(ctx)
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

	result = ListUserInsightMonthlyRequestsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}

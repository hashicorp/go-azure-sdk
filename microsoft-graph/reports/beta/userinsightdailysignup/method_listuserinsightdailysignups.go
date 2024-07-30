package userinsightdailysignup

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

type ListUserInsightDailySignUpsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.UserSignUpMetric
}

type ListUserInsightDailySignUpsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.UserSignUpMetric
}

type ListUserInsightDailySignUpsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListUserInsightDailySignUpsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListUserInsightDailySignUps ...
func (c UserInsightDailySignUpClient) ListUserInsightDailySignUps(ctx context.Context) (result ListUserInsightDailySignUpsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListUserInsightDailySignUpsCustomPager{},
		Path:       "/reports/userInsights/daily/signUps",
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
		Values *[]beta.UserSignUpMetric `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserInsightDailySignUpsComplete retrieves all the results into a single object
func (c UserInsightDailySignUpClient) ListUserInsightDailySignUpsComplete(ctx context.Context) (ListUserInsightDailySignUpsCompleteResult, error) {
	return c.ListUserInsightDailySignUpsCompleteMatchingPredicate(ctx, UserSignUpMetricOperationPredicate{})
}

// ListUserInsightDailySignUpsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserInsightDailySignUpClient) ListUserInsightDailySignUpsCompleteMatchingPredicate(ctx context.Context, predicate UserSignUpMetricOperationPredicate) (result ListUserInsightDailySignUpsCompleteResult, err error) {
	items := make([]beta.UserSignUpMetric, 0)

	resp, err := c.ListUserInsightDailySignUps(ctx)
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

	result = ListUserInsightDailySignUpsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}

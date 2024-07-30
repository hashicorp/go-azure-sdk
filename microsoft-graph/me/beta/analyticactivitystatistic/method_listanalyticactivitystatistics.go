package analyticactivitystatistic

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

type ListAnalyticActivityStatisticsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.ActivityStatistics
}

type ListAnalyticActivityStatisticsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.ActivityStatistics
}

type ListAnalyticActivityStatisticsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListAnalyticActivityStatisticsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListAnalyticActivityStatistics ...
func (c AnalyticActivityStatisticClient) ListAnalyticActivityStatistics(ctx context.Context) (result ListAnalyticActivityStatisticsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListAnalyticActivityStatisticsCustomPager{},
		Path:       "/me/analytics/activityStatistics",
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
		Values *[]beta.ActivityStatistics `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListAnalyticActivityStatisticsComplete retrieves all the results into a single object
func (c AnalyticActivityStatisticClient) ListAnalyticActivityStatisticsComplete(ctx context.Context) (ListAnalyticActivityStatisticsCompleteResult, error) {
	return c.ListAnalyticActivityStatisticsCompleteMatchingPredicate(ctx, ActivityStatisticsOperationPredicate{})
}

// ListAnalyticActivityStatisticsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c AnalyticActivityStatisticClient) ListAnalyticActivityStatisticsCompleteMatchingPredicate(ctx context.Context, predicate ActivityStatisticsOperationPredicate) (result ListAnalyticActivityStatisticsCompleteResult, err error) {
	items := make([]beta.ActivityStatistics, 0)

	resp, err := c.ListAnalyticActivityStatistics(ctx)
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

	result = ListAnalyticActivityStatisticsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}

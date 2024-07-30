package userexperienceanalyticsapphealthoverviewmetricvalue

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListUserExperienceAnalyticsAppHealthOverviewMetricValuesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.UserExperienceAnalyticsMetric
}

type ListUserExperienceAnalyticsAppHealthOverviewMetricValuesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.UserExperienceAnalyticsMetric
}

type ListUserExperienceAnalyticsAppHealthOverviewMetricValuesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListUserExperienceAnalyticsAppHealthOverviewMetricValuesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListUserExperienceAnalyticsAppHealthOverviewMetricValues ...
func (c UserExperienceAnalyticsAppHealthOverviewMetricValueClient) ListUserExperienceAnalyticsAppHealthOverviewMetricValues(ctx context.Context) (result ListUserExperienceAnalyticsAppHealthOverviewMetricValuesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListUserExperienceAnalyticsAppHealthOverviewMetricValuesCustomPager{},
		Path:       "/deviceManagement/userExperienceAnalyticsAppHealthOverview/metricValues",
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
		Values *[]stable.UserExperienceAnalyticsMetric `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserExperienceAnalyticsAppHealthOverviewMetricValuesComplete retrieves all the results into a single object
func (c UserExperienceAnalyticsAppHealthOverviewMetricValueClient) ListUserExperienceAnalyticsAppHealthOverviewMetricValuesComplete(ctx context.Context) (ListUserExperienceAnalyticsAppHealthOverviewMetricValuesCompleteResult, error) {
	return c.ListUserExperienceAnalyticsAppHealthOverviewMetricValuesCompleteMatchingPredicate(ctx, UserExperienceAnalyticsMetricOperationPredicate{})
}

// ListUserExperienceAnalyticsAppHealthOverviewMetricValuesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserExperienceAnalyticsAppHealthOverviewMetricValueClient) ListUserExperienceAnalyticsAppHealthOverviewMetricValuesCompleteMatchingPredicate(ctx context.Context, predicate UserExperienceAnalyticsMetricOperationPredicate) (result ListUserExperienceAnalyticsAppHealthOverviewMetricValuesCompleteResult, err error) {
	items := make([]stable.UserExperienceAnalyticsMetric, 0)

	resp, err := c.ListUserExperienceAnalyticsAppHealthOverviewMetricValues(ctx)
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

	result = ListUserExperienceAnalyticsAppHealthOverviewMetricValuesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}

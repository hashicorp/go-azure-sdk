package userexperienceanalyticscategorymetricvalue

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

type ListUserExperienceAnalyticsCategoryMetricValuesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.UserExperienceAnalyticsMetric
}

type ListUserExperienceAnalyticsCategoryMetricValuesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.UserExperienceAnalyticsMetric
}

type ListUserExperienceAnalyticsCategoryMetricValuesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListUserExperienceAnalyticsCategoryMetricValuesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListUserExperienceAnalyticsCategoryMetricValues ...
func (c UserExperienceAnalyticsCategoryMetricValueClient) ListUserExperienceAnalyticsCategoryMetricValues(ctx context.Context, id DeviceManagementUserExperienceAnalyticsCategoryId) (result ListUserExperienceAnalyticsCategoryMetricValuesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListUserExperienceAnalyticsCategoryMetricValuesCustomPager{},
		Path:       fmt.Sprintf("%s/metricValues", id.ID()),
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
		Values *[]beta.UserExperienceAnalyticsMetric `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserExperienceAnalyticsCategoryMetricValuesComplete retrieves all the results into a single object
func (c UserExperienceAnalyticsCategoryMetricValueClient) ListUserExperienceAnalyticsCategoryMetricValuesComplete(ctx context.Context, id DeviceManagementUserExperienceAnalyticsCategoryId) (ListUserExperienceAnalyticsCategoryMetricValuesCompleteResult, error) {
	return c.ListUserExperienceAnalyticsCategoryMetricValuesCompleteMatchingPredicate(ctx, id, UserExperienceAnalyticsMetricOperationPredicate{})
}

// ListUserExperienceAnalyticsCategoryMetricValuesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserExperienceAnalyticsCategoryMetricValueClient) ListUserExperienceAnalyticsCategoryMetricValuesCompleteMatchingPredicate(ctx context.Context, id DeviceManagementUserExperienceAnalyticsCategoryId, predicate UserExperienceAnalyticsMetricOperationPredicate) (result ListUserExperienceAnalyticsCategoryMetricValuesCompleteResult, err error) {
	items := make([]beta.UserExperienceAnalyticsMetric, 0)

	resp, err := c.ListUserExperienceAnalyticsCategoryMetricValues(ctx, id)
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

	result = ListUserExperienceAnalyticsCategoryMetricValuesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}

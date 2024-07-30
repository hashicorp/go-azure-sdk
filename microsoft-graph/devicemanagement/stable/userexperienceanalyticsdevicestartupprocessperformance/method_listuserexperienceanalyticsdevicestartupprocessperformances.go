package userexperienceanalyticsdevicestartupprocessperformance

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

type ListUserExperienceAnalyticsDeviceStartupProcessPerformancesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.UserExperienceAnalyticsDeviceStartupProcessPerformance
}

type ListUserExperienceAnalyticsDeviceStartupProcessPerformancesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.UserExperienceAnalyticsDeviceStartupProcessPerformance
}

type ListUserExperienceAnalyticsDeviceStartupProcessPerformancesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListUserExperienceAnalyticsDeviceStartupProcessPerformancesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListUserExperienceAnalyticsDeviceStartupProcessPerformances ...
func (c UserExperienceAnalyticsDeviceStartupProcessPerformanceClient) ListUserExperienceAnalyticsDeviceStartupProcessPerformances(ctx context.Context) (result ListUserExperienceAnalyticsDeviceStartupProcessPerformancesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListUserExperienceAnalyticsDeviceStartupProcessPerformancesCustomPager{},
		Path:       "/deviceManagement/userExperienceAnalyticsDeviceStartupProcessPerformance",
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
		Values *[]stable.UserExperienceAnalyticsDeviceStartupProcessPerformance `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserExperienceAnalyticsDeviceStartupProcessPerformancesComplete retrieves all the results into a single object
func (c UserExperienceAnalyticsDeviceStartupProcessPerformanceClient) ListUserExperienceAnalyticsDeviceStartupProcessPerformancesComplete(ctx context.Context) (ListUserExperienceAnalyticsDeviceStartupProcessPerformancesCompleteResult, error) {
	return c.ListUserExperienceAnalyticsDeviceStartupProcessPerformancesCompleteMatchingPredicate(ctx, UserExperienceAnalyticsDeviceStartupProcessPerformanceOperationPredicate{})
}

// ListUserExperienceAnalyticsDeviceStartupProcessPerformancesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserExperienceAnalyticsDeviceStartupProcessPerformanceClient) ListUserExperienceAnalyticsDeviceStartupProcessPerformancesCompleteMatchingPredicate(ctx context.Context, predicate UserExperienceAnalyticsDeviceStartupProcessPerformanceOperationPredicate) (result ListUserExperienceAnalyticsDeviceStartupProcessPerformancesCompleteResult, err error) {
	items := make([]stable.UserExperienceAnalyticsDeviceStartupProcessPerformance, 0)

	resp, err := c.ListUserExperienceAnalyticsDeviceStartupProcessPerformances(ctx)
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

	result = ListUserExperienceAnalyticsDeviceStartupProcessPerformancesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}

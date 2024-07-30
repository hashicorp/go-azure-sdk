package userexperienceanalyticsdeviceperformance

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

type ListUserExperienceAnalyticsDevicePerformancesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.UserExperienceAnalyticsDevicePerformance
}

type ListUserExperienceAnalyticsDevicePerformancesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.UserExperienceAnalyticsDevicePerformance
}

type ListUserExperienceAnalyticsDevicePerformancesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListUserExperienceAnalyticsDevicePerformancesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListUserExperienceAnalyticsDevicePerformances ...
func (c UserExperienceAnalyticsDevicePerformanceClient) ListUserExperienceAnalyticsDevicePerformances(ctx context.Context) (result ListUserExperienceAnalyticsDevicePerformancesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListUserExperienceAnalyticsDevicePerformancesCustomPager{},
		Path:       "/deviceManagement/userExperienceAnalyticsDevicePerformance",
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
		Values *[]stable.UserExperienceAnalyticsDevicePerformance `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserExperienceAnalyticsDevicePerformancesComplete retrieves all the results into a single object
func (c UserExperienceAnalyticsDevicePerformanceClient) ListUserExperienceAnalyticsDevicePerformancesComplete(ctx context.Context) (ListUserExperienceAnalyticsDevicePerformancesCompleteResult, error) {
	return c.ListUserExperienceAnalyticsDevicePerformancesCompleteMatchingPredicate(ctx, UserExperienceAnalyticsDevicePerformanceOperationPredicate{})
}

// ListUserExperienceAnalyticsDevicePerformancesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserExperienceAnalyticsDevicePerformanceClient) ListUserExperienceAnalyticsDevicePerformancesCompleteMatchingPredicate(ctx context.Context, predicate UserExperienceAnalyticsDevicePerformanceOperationPredicate) (result ListUserExperienceAnalyticsDevicePerformancesCompleteResult, err error) {
	items := make([]stable.UserExperienceAnalyticsDevicePerformance, 0)

	resp, err := c.ListUserExperienceAnalyticsDevicePerformances(ctx)
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

	result = ListUserExperienceAnalyticsDevicePerformancesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}

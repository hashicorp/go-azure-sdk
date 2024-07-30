package userexperienceanalyticsapphealthdeviceperformance

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

type ListUserExperienceAnalyticsAppHealthDevicePerformancesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.UserExperienceAnalyticsAppHealthDevicePerformance
}

type ListUserExperienceAnalyticsAppHealthDevicePerformancesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.UserExperienceAnalyticsAppHealthDevicePerformance
}

type ListUserExperienceAnalyticsAppHealthDevicePerformancesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListUserExperienceAnalyticsAppHealthDevicePerformancesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListUserExperienceAnalyticsAppHealthDevicePerformances ...
func (c UserExperienceAnalyticsAppHealthDevicePerformanceClient) ListUserExperienceAnalyticsAppHealthDevicePerformances(ctx context.Context) (result ListUserExperienceAnalyticsAppHealthDevicePerformancesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListUserExperienceAnalyticsAppHealthDevicePerformancesCustomPager{},
		Path:       "/deviceManagement/userExperienceAnalyticsAppHealthDevicePerformance",
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
		Values *[]stable.UserExperienceAnalyticsAppHealthDevicePerformance `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserExperienceAnalyticsAppHealthDevicePerformancesComplete retrieves all the results into a single object
func (c UserExperienceAnalyticsAppHealthDevicePerformanceClient) ListUserExperienceAnalyticsAppHealthDevicePerformancesComplete(ctx context.Context) (ListUserExperienceAnalyticsAppHealthDevicePerformancesCompleteResult, error) {
	return c.ListUserExperienceAnalyticsAppHealthDevicePerformancesCompleteMatchingPredicate(ctx, UserExperienceAnalyticsAppHealthDevicePerformanceOperationPredicate{})
}

// ListUserExperienceAnalyticsAppHealthDevicePerformancesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserExperienceAnalyticsAppHealthDevicePerformanceClient) ListUserExperienceAnalyticsAppHealthDevicePerformancesCompleteMatchingPredicate(ctx context.Context, predicate UserExperienceAnalyticsAppHealthDevicePerformanceOperationPredicate) (result ListUserExperienceAnalyticsAppHealthDevicePerformancesCompleteResult, err error) {
	items := make([]stable.UserExperienceAnalyticsAppHealthDevicePerformance, 0)

	resp, err := c.ListUserExperienceAnalyticsAppHealthDevicePerformances(ctx)
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

	result = ListUserExperienceAnalyticsAppHealthDevicePerformancesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}

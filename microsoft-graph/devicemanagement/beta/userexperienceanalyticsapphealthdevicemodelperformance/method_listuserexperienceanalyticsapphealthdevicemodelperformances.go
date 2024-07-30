package userexperienceanalyticsapphealthdevicemodelperformance

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

type ListUserExperienceAnalyticsAppHealthDeviceModelPerformancesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.UserExperienceAnalyticsAppHealthDeviceModelPerformance
}

type ListUserExperienceAnalyticsAppHealthDeviceModelPerformancesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.UserExperienceAnalyticsAppHealthDeviceModelPerformance
}

type ListUserExperienceAnalyticsAppHealthDeviceModelPerformancesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListUserExperienceAnalyticsAppHealthDeviceModelPerformancesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListUserExperienceAnalyticsAppHealthDeviceModelPerformances ...
func (c UserExperienceAnalyticsAppHealthDeviceModelPerformanceClient) ListUserExperienceAnalyticsAppHealthDeviceModelPerformances(ctx context.Context) (result ListUserExperienceAnalyticsAppHealthDeviceModelPerformancesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListUserExperienceAnalyticsAppHealthDeviceModelPerformancesCustomPager{},
		Path:       "/deviceManagement/userExperienceAnalyticsAppHealthDeviceModelPerformance",
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
		Values *[]beta.UserExperienceAnalyticsAppHealthDeviceModelPerformance `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserExperienceAnalyticsAppHealthDeviceModelPerformancesComplete retrieves all the results into a single object
func (c UserExperienceAnalyticsAppHealthDeviceModelPerformanceClient) ListUserExperienceAnalyticsAppHealthDeviceModelPerformancesComplete(ctx context.Context) (ListUserExperienceAnalyticsAppHealthDeviceModelPerformancesCompleteResult, error) {
	return c.ListUserExperienceAnalyticsAppHealthDeviceModelPerformancesCompleteMatchingPredicate(ctx, UserExperienceAnalyticsAppHealthDeviceModelPerformanceOperationPredicate{})
}

// ListUserExperienceAnalyticsAppHealthDeviceModelPerformancesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserExperienceAnalyticsAppHealthDeviceModelPerformanceClient) ListUserExperienceAnalyticsAppHealthDeviceModelPerformancesCompleteMatchingPredicate(ctx context.Context, predicate UserExperienceAnalyticsAppHealthDeviceModelPerformanceOperationPredicate) (result ListUserExperienceAnalyticsAppHealthDeviceModelPerformancesCompleteResult, err error) {
	items := make([]beta.UserExperienceAnalyticsAppHealthDeviceModelPerformance, 0)

	resp, err := c.ListUserExperienceAnalyticsAppHealthDeviceModelPerformances(ctx)
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

	result = ListUserExperienceAnalyticsAppHealthDeviceModelPerformancesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}

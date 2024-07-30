package userexperienceanalyticsapphealthapplicationperformancebyappversion

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

type ListUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.UserExperienceAnalyticsAppHealthAppPerformanceByAppVersion
}

type ListUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.UserExperienceAnalyticsAppHealthAppPerformanceByAppVersion
}

type ListUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersions ...
func (c UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionClient) ListUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersions(ctx context.Context) (result ListUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionsCustomPager{},
		Path:       "/deviceManagement/userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersion",
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
		Values *[]beta.UserExperienceAnalyticsAppHealthAppPerformanceByAppVersion `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionsComplete retrieves all the results into a single object
func (c UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionClient) ListUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionsComplete(ctx context.Context) (ListUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionsCompleteResult, error) {
	return c.ListUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionsCompleteMatchingPredicate(ctx, UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionOperationPredicate{})
}

// ListUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionClient) ListUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionsCompleteMatchingPredicate(ctx context.Context, predicate UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionOperationPredicate) (result ListUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionsCompleteResult, err error) {
	items := make([]beta.UserExperienceAnalyticsAppHealthAppPerformanceByAppVersion, 0)

	resp, err := c.ListUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersions(ctx)
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

	result = ListUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}

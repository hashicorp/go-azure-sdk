package userexperienceanalyticsapphealthapplicationperformance

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

type ListUserExperienceAnalyticsAppHealthApplicationPerformancesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.UserExperienceAnalyticsAppHealthApplicationPerformance
}

type ListUserExperienceAnalyticsAppHealthApplicationPerformancesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.UserExperienceAnalyticsAppHealthApplicationPerformance
}

type ListUserExperienceAnalyticsAppHealthApplicationPerformancesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListUserExperienceAnalyticsAppHealthApplicationPerformancesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListUserExperienceAnalyticsAppHealthApplicationPerformances ...
func (c UserExperienceAnalyticsAppHealthApplicationPerformanceClient) ListUserExperienceAnalyticsAppHealthApplicationPerformances(ctx context.Context) (result ListUserExperienceAnalyticsAppHealthApplicationPerformancesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListUserExperienceAnalyticsAppHealthApplicationPerformancesCustomPager{},
		Path:       "/deviceManagement/userExperienceAnalyticsAppHealthApplicationPerformance",
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
		Values *[]beta.UserExperienceAnalyticsAppHealthApplicationPerformance `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserExperienceAnalyticsAppHealthApplicationPerformancesComplete retrieves all the results into a single object
func (c UserExperienceAnalyticsAppHealthApplicationPerformanceClient) ListUserExperienceAnalyticsAppHealthApplicationPerformancesComplete(ctx context.Context) (ListUserExperienceAnalyticsAppHealthApplicationPerformancesCompleteResult, error) {
	return c.ListUserExperienceAnalyticsAppHealthApplicationPerformancesCompleteMatchingPredicate(ctx, UserExperienceAnalyticsAppHealthApplicationPerformanceOperationPredicate{})
}

// ListUserExperienceAnalyticsAppHealthApplicationPerformancesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserExperienceAnalyticsAppHealthApplicationPerformanceClient) ListUserExperienceAnalyticsAppHealthApplicationPerformancesCompleteMatchingPredicate(ctx context.Context, predicate UserExperienceAnalyticsAppHealthApplicationPerformanceOperationPredicate) (result ListUserExperienceAnalyticsAppHealthApplicationPerformancesCompleteResult, err error) {
	items := make([]beta.UserExperienceAnalyticsAppHealthApplicationPerformance, 0)

	resp, err := c.ListUserExperienceAnalyticsAppHealthApplicationPerformances(ctx)
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

	result = ListUserExperienceAnalyticsAppHealthApplicationPerformancesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}

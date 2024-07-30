package userexperienceanalyticsapphealthapplicationperformancebyappversiondetail

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

type ListUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetails
}

type ListUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetails
}

type ListUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails ...
func (c UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailClient) ListUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails(ctx context.Context) (result ListUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsCustomPager{},
		Path:       "/deviceManagement/userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails",
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
		Values *[]beta.UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetails `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsComplete retrieves all the results into a single object
func (c UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailClient) ListUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsComplete(ctx context.Context) (ListUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsCompleteResult, error) {
	return c.ListUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsCompleteMatchingPredicate(ctx, UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsOperationPredicate{})
}

// ListUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailClient) ListUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsCompleteMatchingPredicate(ctx context.Context, predicate UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsOperationPredicate) (result ListUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsCompleteResult, err error) {
	items := make([]beta.UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetails, 0)

	resp, err := c.ListUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails(ctx)
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

	result = ListUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}

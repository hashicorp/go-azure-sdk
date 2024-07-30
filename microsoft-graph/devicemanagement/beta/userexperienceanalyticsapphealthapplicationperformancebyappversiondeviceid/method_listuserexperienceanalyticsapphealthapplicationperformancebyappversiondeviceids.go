package userexperienceanalyticsapphealthapplicationperformancebyappversiondeviceid

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

type ListUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceId
}

type ListUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceId
}

type ListUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIds ...
func (c UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdClient) ListUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIds(ctx context.Context) (result ListUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdsCustomPager{},
		Path:       "/deviceManagement/userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId",
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
		Values *[]beta.UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceId `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdsComplete retrieves all the results into a single object
func (c UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdClient) ListUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdsComplete(ctx context.Context) (ListUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdsCompleteResult, error) {
	return c.ListUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdsCompleteMatchingPredicate(ctx, UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdOperationPredicate{})
}

// ListUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdClient) ListUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdsCompleteMatchingPredicate(ctx context.Context, predicate UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdOperationPredicate) (result ListUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdsCompleteResult, err error) {
	items := make([]beta.UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceId, 0)

	resp, err := c.ListUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIds(ctx)
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

	result = ListUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}

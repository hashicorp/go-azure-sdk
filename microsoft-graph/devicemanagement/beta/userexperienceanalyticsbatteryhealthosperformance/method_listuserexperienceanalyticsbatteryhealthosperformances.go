package userexperienceanalyticsbatteryhealthosperformance

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

type ListUserExperienceAnalyticsBatteryHealthOsPerformancesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.UserExperienceAnalyticsBatteryHealthOsPerformance
}

type ListUserExperienceAnalyticsBatteryHealthOsPerformancesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.UserExperienceAnalyticsBatteryHealthOsPerformance
}

type ListUserExperienceAnalyticsBatteryHealthOsPerformancesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListUserExperienceAnalyticsBatteryHealthOsPerformancesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListUserExperienceAnalyticsBatteryHealthOsPerformances ...
func (c UserExperienceAnalyticsBatteryHealthOsPerformanceClient) ListUserExperienceAnalyticsBatteryHealthOsPerformances(ctx context.Context) (result ListUserExperienceAnalyticsBatteryHealthOsPerformancesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListUserExperienceAnalyticsBatteryHealthOsPerformancesCustomPager{},
		Path:       "/deviceManagement/userExperienceAnalyticsBatteryHealthOsPerformance",
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
		Values *[]beta.UserExperienceAnalyticsBatteryHealthOsPerformance `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserExperienceAnalyticsBatteryHealthOsPerformancesComplete retrieves all the results into a single object
func (c UserExperienceAnalyticsBatteryHealthOsPerformanceClient) ListUserExperienceAnalyticsBatteryHealthOsPerformancesComplete(ctx context.Context) (ListUserExperienceAnalyticsBatteryHealthOsPerformancesCompleteResult, error) {
	return c.ListUserExperienceAnalyticsBatteryHealthOsPerformancesCompleteMatchingPredicate(ctx, UserExperienceAnalyticsBatteryHealthOsPerformanceOperationPredicate{})
}

// ListUserExperienceAnalyticsBatteryHealthOsPerformancesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserExperienceAnalyticsBatteryHealthOsPerformanceClient) ListUserExperienceAnalyticsBatteryHealthOsPerformancesCompleteMatchingPredicate(ctx context.Context, predicate UserExperienceAnalyticsBatteryHealthOsPerformanceOperationPredicate) (result ListUserExperienceAnalyticsBatteryHealthOsPerformancesCompleteResult, err error) {
	items := make([]beta.UserExperienceAnalyticsBatteryHealthOsPerformance, 0)

	resp, err := c.ListUserExperienceAnalyticsBatteryHealthOsPerformances(ctx)
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

	result = ListUserExperienceAnalyticsBatteryHealthOsPerformancesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}

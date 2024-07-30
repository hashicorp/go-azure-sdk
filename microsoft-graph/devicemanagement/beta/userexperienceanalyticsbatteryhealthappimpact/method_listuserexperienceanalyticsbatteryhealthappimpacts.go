package userexperienceanalyticsbatteryhealthappimpact

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

type ListUserExperienceAnalyticsBatteryHealthAppImpactsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.UserExperienceAnalyticsBatteryHealthAppImpact
}

type ListUserExperienceAnalyticsBatteryHealthAppImpactsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.UserExperienceAnalyticsBatteryHealthAppImpact
}

type ListUserExperienceAnalyticsBatteryHealthAppImpactsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListUserExperienceAnalyticsBatteryHealthAppImpactsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListUserExperienceAnalyticsBatteryHealthAppImpacts ...
func (c UserExperienceAnalyticsBatteryHealthAppImpactClient) ListUserExperienceAnalyticsBatteryHealthAppImpacts(ctx context.Context) (result ListUserExperienceAnalyticsBatteryHealthAppImpactsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListUserExperienceAnalyticsBatteryHealthAppImpactsCustomPager{},
		Path:       "/deviceManagement/userExperienceAnalyticsBatteryHealthAppImpact",
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
		Values *[]beta.UserExperienceAnalyticsBatteryHealthAppImpact `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserExperienceAnalyticsBatteryHealthAppImpactsComplete retrieves all the results into a single object
func (c UserExperienceAnalyticsBatteryHealthAppImpactClient) ListUserExperienceAnalyticsBatteryHealthAppImpactsComplete(ctx context.Context) (ListUserExperienceAnalyticsBatteryHealthAppImpactsCompleteResult, error) {
	return c.ListUserExperienceAnalyticsBatteryHealthAppImpactsCompleteMatchingPredicate(ctx, UserExperienceAnalyticsBatteryHealthAppImpactOperationPredicate{})
}

// ListUserExperienceAnalyticsBatteryHealthAppImpactsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserExperienceAnalyticsBatteryHealthAppImpactClient) ListUserExperienceAnalyticsBatteryHealthAppImpactsCompleteMatchingPredicate(ctx context.Context, predicate UserExperienceAnalyticsBatteryHealthAppImpactOperationPredicate) (result ListUserExperienceAnalyticsBatteryHealthAppImpactsCompleteResult, err error) {
	items := make([]beta.UserExperienceAnalyticsBatteryHealthAppImpact, 0)

	resp, err := c.ListUserExperienceAnalyticsBatteryHealthAppImpacts(ctx)
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

	result = ListUserExperienceAnalyticsBatteryHealthAppImpactsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}

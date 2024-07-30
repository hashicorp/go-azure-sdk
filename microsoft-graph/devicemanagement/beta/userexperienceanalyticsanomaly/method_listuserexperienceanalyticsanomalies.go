package userexperienceanalyticsanomaly

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

type ListUserExperienceAnalyticsAnomaliesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.UserExperienceAnalyticsAnomaly
}

type ListUserExperienceAnalyticsAnomaliesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.UserExperienceAnalyticsAnomaly
}

type ListUserExperienceAnalyticsAnomaliesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListUserExperienceAnalyticsAnomaliesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListUserExperienceAnalyticsAnomalies ...
func (c UserExperienceAnalyticsAnomalyClient) ListUserExperienceAnalyticsAnomalies(ctx context.Context) (result ListUserExperienceAnalyticsAnomaliesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListUserExperienceAnalyticsAnomaliesCustomPager{},
		Path:       "/deviceManagement/userExperienceAnalyticsAnomaly",
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
		Values *[]beta.UserExperienceAnalyticsAnomaly `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserExperienceAnalyticsAnomaliesComplete retrieves all the results into a single object
func (c UserExperienceAnalyticsAnomalyClient) ListUserExperienceAnalyticsAnomaliesComplete(ctx context.Context) (ListUserExperienceAnalyticsAnomaliesCompleteResult, error) {
	return c.ListUserExperienceAnalyticsAnomaliesCompleteMatchingPredicate(ctx, UserExperienceAnalyticsAnomalyOperationPredicate{})
}

// ListUserExperienceAnalyticsAnomaliesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserExperienceAnalyticsAnomalyClient) ListUserExperienceAnalyticsAnomaliesCompleteMatchingPredicate(ctx context.Context, predicate UserExperienceAnalyticsAnomalyOperationPredicate) (result ListUserExperienceAnalyticsAnomaliesCompleteResult, err error) {
	items := make([]beta.UserExperienceAnalyticsAnomaly, 0)

	resp, err := c.ListUserExperienceAnalyticsAnomalies(ctx)
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

	result = ListUserExperienceAnalyticsAnomaliesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}

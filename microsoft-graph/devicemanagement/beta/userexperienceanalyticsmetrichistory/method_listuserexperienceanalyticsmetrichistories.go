package userexperienceanalyticsmetrichistory

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

type ListUserExperienceAnalyticsMetricHistoriesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.UserExperienceAnalyticsMetricHistory
}

type ListUserExperienceAnalyticsMetricHistoriesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.UserExperienceAnalyticsMetricHistory
}

type ListUserExperienceAnalyticsMetricHistoriesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListUserExperienceAnalyticsMetricHistoriesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListUserExperienceAnalyticsMetricHistories ...
func (c UserExperienceAnalyticsMetricHistoryClient) ListUserExperienceAnalyticsMetricHistories(ctx context.Context) (result ListUserExperienceAnalyticsMetricHistoriesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListUserExperienceAnalyticsMetricHistoriesCustomPager{},
		Path:       "/deviceManagement/userExperienceAnalyticsMetricHistory",
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
		Values *[]beta.UserExperienceAnalyticsMetricHistory `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserExperienceAnalyticsMetricHistoriesComplete retrieves all the results into a single object
func (c UserExperienceAnalyticsMetricHistoryClient) ListUserExperienceAnalyticsMetricHistoriesComplete(ctx context.Context) (ListUserExperienceAnalyticsMetricHistoriesCompleteResult, error) {
	return c.ListUserExperienceAnalyticsMetricHistoriesCompleteMatchingPredicate(ctx, UserExperienceAnalyticsMetricHistoryOperationPredicate{})
}

// ListUserExperienceAnalyticsMetricHistoriesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserExperienceAnalyticsMetricHistoryClient) ListUserExperienceAnalyticsMetricHistoriesCompleteMatchingPredicate(ctx context.Context, predicate UserExperienceAnalyticsMetricHistoryOperationPredicate) (result ListUserExperienceAnalyticsMetricHistoriesCompleteResult, err error) {
	items := make([]beta.UserExperienceAnalyticsMetricHistory, 0)

	resp, err := c.ListUserExperienceAnalyticsMetricHistories(ctx)
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

	result = ListUserExperienceAnalyticsMetricHistoriesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}

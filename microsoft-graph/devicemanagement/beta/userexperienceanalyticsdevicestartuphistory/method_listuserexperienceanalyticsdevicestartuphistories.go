package userexperienceanalyticsdevicestartuphistory

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

type ListUserExperienceAnalyticsDeviceStartupHistoriesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.UserExperienceAnalyticsDeviceStartupHistory
}

type ListUserExperienceAnalyticsDeviceStartupHistoriesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.UserExperienceAnalyticsDeviceStartupHistory
}

type ListUserExperienceAnalyticsDeviceStartupHistoriesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListUserExperienceAnalyticsDeviceStartupHistoriesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListUserExperienceAnalyticsDeviceStartupHistories ...
func (c UserExperienceAnalyticsDeviceStartupHistoryClient) ListUserExperienceAnalyticsDeviceStartupHistories(ctx context.Context) (result ListUserExperienceAnalyticsDeviceStartupHistoriesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListUserExperienceAnalyticsDeviceStartupHistoriesCustomPager{},
		Path:       "/deviceManagement/userExperienceAnalyticsDeviceStartupHistory",
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
		Values *[]beta.UserExperienceAnalyticsDeviceStartupHistory `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserExperienceAnalyticsDeviceStartupHistoriesComplete retrieves all the results into a single object
func (c UserExperienceAnalyticsDeviceStartupHistoryClient) ListUserExperienceAnalyticsDeviceStartupHistoriesComplete(ctx context.Context) (ListUserExperienceAnalyticsDeviceStartupHistoriesCompleteResult, error) {
	return c.ListUserExperienceAnalyticsDeviceStartupHistoriesCompleteMatchingPredicate(ctx, UserExperienceAnalyticsDeviceStartupHistoryOperationPredicate{})
}

// ListUserExperienceAnalyticsDeviceStartupHistoriesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserExperienceAnalyticsDeviceStartupHistoryClient) ListUserExperienceAnalyticsDeviceStartupHistoriesCompleteMatchingPredicate(ctx context.Context, predicate UserExperienceAnalyticsDeviceStartupHistoryOperationPredicate) (result ListUserExperienceAnalyticsDeviceStartupHistoriesCompleteResult, err error) {
	items := make([]beta.UserExperienceAnalyticsDeviceStartupHistory, 0)

	resp, err := c.ListUserExperienceAnalyticsDeviceStartupHistories(ctx)
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

	result = ListUserExperienceAnalyticsDeviceStartupHistoriesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}

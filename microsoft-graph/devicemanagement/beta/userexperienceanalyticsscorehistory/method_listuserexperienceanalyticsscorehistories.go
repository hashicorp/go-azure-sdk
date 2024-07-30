package userexperienceanalyticsscorehistory

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

type ListUserExperienceAnalyticsScoreHistoriesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.UserExperienceAnalyticsScoreHistory
}

type ListUserExperienceAnalyticsScoreHistoriesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.UserExperienceAnalyticsScoreHistory
}

type ListUserExperienceAnalyticsScoreHistoriesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListUserExperienceAnalyticsScoreHistoriesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListUserExperienceAnalyticsScoreHistories ...
func (c UserExperienceAnalyticsScoreHistoryClient) ListUserExperienceAnalyticsScoreHistories(ctx context.Context) (result ListUserExperienceAnalyticsScoreHistoriesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListUserExperienceAnalyticsScoreHistoriesCustomPager{},
		Path:       "/deviceManagement/userExperienceAnalyticsScoreHistory",
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
		Values *[]beta.UserExperienceAnalyticsScoreHistory `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserExperienceAnalyticsScoreHistoriesComplete retrieves all the results into a single object
func (c UserExperienceAnalyticsScoreHistoryClient) ListUserExperienceAnalyticsScoreHistoriesComplete(ctx context.Context) (ListUserExperienceAnalyticsScoreHistoriesCompleteResult, error) {
	return c.ListUserExperienceAnalyticsScoreHistoriesCompleteMatchingPredicate(ctx, UserExperienceAnalyticsScoreHistoryOperationPredicate{})
}

// ListUserExperienceAnalyticsScoreHistoriesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserExperienceAnalyticsScoreHistoryClient) ListUserExperienceAnalyticsScoreHistoriesCompleteMatchingPredicate(ctx context.Context, predicate UserExperienceAnalyticsScoreHistoryOperationPredicate) (result ListUserExperienceAnalyticsScoreHistoriesCompleteResult, err error) {
	items := make([]beta.UserExperienceAnalyticsScoreHistory, 0)

	resp, err := c.ListUserExperienceAnalyticsScoreHistories(ctx)
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

	result = ListUserExperienceAnalyticsScoreHistoriesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}

package userexperienceanalyticsdevicescore

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListUserExperienceAnalyticsDeviceScoresOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.UserExperienceAnalyticsDeviceScores
}

type ListUserExperienceAnalyticsDeviceScoresCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.UserExperienceAnalyticsDeviceScores
}

type ListUserExperienceAnalyticsDeviceScoresCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListUserExperienceAnalyticsDeviceScoresCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListUserExperienceAnalyticsDeviceScores ...
func (c UserExperienceAnalyticsDeviceScoreClient) ListUserExperienceAnalyticsDeviceScores(ctx context.Context) (result ListUserExperienceAnalyticsDeviceScoresOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListUserExperienceAnalyticsDeviceScoresCustomPager{},
		Path:       "/deviceManagement/userExperienceAnalyticsDeviceScores",
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
		Values *[]stable.UserExperienceAnalyticsDeviceScores `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserExperienceAnalyticsDeviceScoresComplete retrieves all the results into a single object
func (c UserExperienceAnalyticsDeviceScoreClient) ListUserExperienceAnalyticsDeviceScoresComplete(ctx context.Context) (ListUserExperienceAnalyticsDeviceScoresCompleteResult, error) {
	return c.ListUserExperienceAnalyticsDeviceScoresCompleteMatchingPredicate(ctx, UserExperienceAnalyticsDeviceScoresOperationPredicate{})
}

// ListUserExperienceAnalyticsDeviceScoresCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserExperienceAnalyticsDeviceScoreClient) ListUserExperienceAnalyticsDeviceScoresCompleteMatchingPredicate(ctx context.Context, predicate UserExperienceAnalyticsDeviceScoresOperationPredicate) (result ListUserExperienceAnalyticsDeviceScoresCompleteResult, err error) {
	items := make([]stable.UserExperienceAnalyticsDeviceScores, 0)

	resp, err := c.ListUserExperienceAnalyticsDeviceScores(ctx)
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

	result = ListUserExperienceAnalyticsDeviceScoresCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}

package userexperienceanalyticsmodelscore

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

type ListUserExperienceAnalyticsModelScoresOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.UserExperienceAnalyticsModelScores
}

type ListUserExperienceAnalyticsModelScoresCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.UserExperienceAnalyticsModelScores
}

type ListUserExperienceAnalyticsModelScoresCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListUserExperienceAnalyticsModelScoresCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListUserExperienceAnalyticsModelScores ...
func (c UserExperienceAnalyticsModelScoreClient) ListUserExperienceAnalyticsModelScores(ctx context.Context) (result ListUserExperienceAnalyticsModelScoresOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListUserExperienceAnalyticsModelScoresCustomPager{},
		Path:       "/deviceManagement/userExperienceAnalyticsModelScores",
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
		Values *[]stable.UserExperienceAnalyticsModelScores `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserExperienceAnalyticsModelScoresComplete retrieves all the results into a single object
func (c UserExperienceAnalyticsModelScoreClient) ListUserExperienceAnalyticsModelScoresComplete(ctx context.Context) (ListUserExperienceAnalyticsModelScoresCompleteResult, error) {
	return c.ListUserExperienceAnalyticsModelScoresCompleteMatchingPredicate(ctx, UserExperienceAnalyticsModelScoresOperationPredicate{})
}

// ListUserExperienceAnalyticsModelScoresCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserExperienceAnalyticsModelScoreClient) ListUserExperienceAnalyticsModelScoresCompleteMatchingPredicate(ctx context.Context, predicate UserExperienceAnalyticsModelScoresOperationPredicate) (result ListUserExperienceAnalyticsModelScoresCompleteResult, err error) {
	items := make([]stable.UserExperienceAnalyticsModelScores, 0)

	resp, err := c.ListUserExperienceAnalyticsModelScores(ctx)
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

	result = ListUserExperienceAnalyticsModelScoresCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}

package userpendingaccessreviewinstancedecisioninstancestagedecisioninsight

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common/beta/models"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListUserByIdPendingAccessReviewInstanceByIdDecisionByIdInstanceStageByIdDecisionByIdInsightsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.GovernanceInsightCollectionResponse
}

type ListUserByIdPendingAccessReviewInstanceByIdDecisionByIdInstanceStageByIdDecisionByIdInsightsCompleteResult struct {
	Items []models.GovernanceInsightCollectionResponse
}

// ListUserByIdPendingAccessReviewInstanceByIdDecisionByIdInstanceStageByIdDecisionByIdInsights ...
func (c UserPendingAccessReviewInstanceDecisionInstanceStageDecisionInsightClient) ListUserByIdPendingAccessReviewInstanceByIdDecisionByIdInstanceStageByIdDecisionByIdInsights(ctx context.Context, id UserPendingAccessReviewInstanceDecisionInstanceStageDecisionId) (result ListUserByIdPendingAccessReviewInstanceByIdDecisionByIdInstanceStageByIdDecisionByIdInsightsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/insights", id.ID()),
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
		Values *[]models.GovernanceInsightCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserByIdPendingAccessReviewInstanceByIdDecisionByIdInstanceStageByIdDecisionByIdInsightsComplete retrieves all the results into a single object
func (c UserPendingAccessReviewInstanceDecisionInstanceStageDecisionInsightClient) ListUserByIdPendingAccessReviewInstanceByIdDecisionByIdInstanceStageByIdDecisionByIdInsightsComplete(ctx context.Context, id UserPendingAccessReviewInstanceDecisionInstanceStageDecisionId) (ListUserByIdPendingAccessReviewInstanceByIdDecisionByIdInstanceStageByIdDecisionByIdInsightsCompleteResult, error) {
	return c.ListUserByIdPendingAccessReviewInstanceByIdDecisionByIdInstanceStageByIdDecisionByIdInsightsCompleteMatchingPredicate(ctx, id, models.GovernanceInsightCollectionResponseOperationPredicate{})
}

// ListUserByIdPendingAccessReviewInstanceByIdDecisionByIdInstanceStageByIdDecisionByIdInsightsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserPendingAccessReviewInstanceDecisionInstanceStageDecisionInsightClient) ListUserByIdPendingAccessReviewInstanceByIdDecisionByIdInstanceStageByIdDecisionByIdInsightsCompleteMatchingPredicate(ctx context.Context, id UserPendingAccessReviewInstanceDecisionInstanceStageDecisionId, predicate models.GovernanceInsightCollectionResponseOperationPredicate) (result ListUserByIdPendingAccessReviewInstanceByIdDecisionByIdInstanceStageByIdDecisionByIdInsightsCompleteResult, err error) {
	items := make([]models.GovernanceInsightCollectionResponse, 0)

	resp, err := c.ListUserByIdPendingAccessReviewInstanceByIdDecisionByIdInstanceStageByIdDecisionByIdInsights(ctx, id)
	if err != nil {
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

	result = ListUserByIdPendingAccessReviewInstanceByIdDecisionByIdInstanceStageByIdDecisionByIdInsightsCompleteResult{
		Items: items,
	}
	return
}

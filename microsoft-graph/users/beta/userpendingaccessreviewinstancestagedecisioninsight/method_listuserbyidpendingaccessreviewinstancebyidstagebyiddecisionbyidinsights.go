package userpendingaccessreviewinstancestagedecisioninsight

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

type ListUserByIdPendingAccessReviewInstanceByIdStageByIdDecisionByIdInsightsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.GovernanceInsightCollectionResponse
}

type ListUserByIdPendingAccessReviewInstanceByIdStageByIdDecisionByIdInsightsCompleteResult struct {
	Items []models.GovernanceInsightCollectionResponse
}

// ListUserByIdPendingAccessReviewInstanceByIdStageByIdDecisionByIdInsights ...
func (c UserPendingAccessReviewInstanceStageDecisionInsightClient) ListUserByIdPendingAccessReviewInstanceByIdStageByIdDecisionByIdInsights(ctx context.Context, id UserPendingAccessReviewInstanceStageDecisionId) (result ListUserByIdPendingAccessReviewInstanceByIdStageByIdDecisionByIdInsightsOperationResponse, err error) {
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

// ListUserByIdPendingAccessReviewInstanceByIdStageByIdDecisionByIdInsightsComplete retrieves all the results into a single object
func (c UserPendingAccessReviewInstanceStageDecisionInsightClient) ListUserByIdPendingAccessReviewInstanceByIdStageByIdDecisionByIdInsightsComplete(ctx context.Context, id UserPendingAccessReviewInstanceStageDecisionId) (ListUserByIdPendingAccessReviewInstanceByIdStageByIdDecisionByIdInsightsCompleteResult, error) {
	return c.ListUserByIdPendingAccessReviewInstanceByIdStageByIdDecisionByIdInsightsCompleteMatchingPredicate(ctx, id, models.GovernanceInsightCollectionResponseOperationPredicate{})
}

// ListUserByIdPendingAccessReviewInstanceByIdStageByIdDecisionByIdInsightsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserPendingAccessReviewInstanceStageDecisionInsightClient) ListUserByIdPendingAccessReviewInstanceByIdStageByIdDecisionByIdInsightsCompleteMatchingPredicate(ctx context.Context, id UserPendingAccessReviewInstanceStageDecisionId, predicate models.GovernanceInsightCollectionResponseOperationPredicate) (result ListUserByIdPendingAccessReviewInstanceByIdStageByIdDecisionByIdInsightsCompleteResult, err error) {
	items := make([]models.GovernanceInsightCollectionResponse, 0)

	resp, err := c.ListUserByIdPendingAccessReviewInstanceByIdStageByIdDecisionByIdInsights(ctx, id)
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

	result = ListUserByIdPendingAccessReviewInstanceByIdStageByIdDecisionByIdInsightsCompleteResult{
		Items: items,
	}
	return
}

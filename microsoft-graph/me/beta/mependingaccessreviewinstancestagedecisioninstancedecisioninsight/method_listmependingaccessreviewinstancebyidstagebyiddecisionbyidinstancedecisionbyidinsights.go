package mependingaccessreviewinstancestagedecisioninstancedecisioninsight

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

type ListMePendingAccessReviewInstanceByIdStageByIdDecisionByIdInstanceDecisionByIdInsightsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.GovernanceInsightCollectionResponse
}

type ListMePendingAccessReviewInstanceByIdStageByIdDecisionByIdInstanceDecisionByIdInsightsCompleteResult struct {
	Items []models.GovernanceInsightCollectionResponse
}

// ListMePendingAccessReviewInstanceByIdStageByIdDecisionByIdInstanceDecisionByIdInsights ...
func (c MePendingAccessReviewInstanceStageDecisionInstanceDecisionInsightClient) ListMePendingAccessReviewInstanceByIdStageByIdDecisionByIdInstanceDecisionByIdInsights(ctx context.Context, id MePendingAccessReviewInstanceStageDecisionInstanceDecisionId) (result ListMePendingAccessReviewInstanceByIdStageByIdDecisionByIdInstanceDecisionByIdInsightsOperationResponse, err error) {
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

// ListMePendingAccessReviewInstanceByIdStageByIdDecisionByIdInstanceDecisionByIdInsightsComplete retrieves all the results into a single object
func (c MePendingAccessReviewInstanceStageDecisionInstanceDecisionInsightClient) ListMePendingAccessReviewInstanceByIdStageByIdDecisionByIdInstanceDecisionByIdInsightsComplete(ctx context.Context, id MePendingAccessReviewInstanceStageDecisionInstanceDecisionId) (ListMePendingAccessReviewInstanceByIdStageByIdDecisionByIdInstanceDecisionByIdInsightsCompleteResult, error) {
	return c.ListMePendingAccessReviewInstanceByIdStageByIdDecisionByIdInstanceDecisionByIdInsightsCompleteMatchingPredicate(ctx, id, models.GovernanceInsightCollectionResponseOperationPredicate{})
}

// ListMePendingAccessReviewInstanceByIdStageByIdDecisionByIdInstanceDecisionByIdInsightsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MePendingAccessReviewInstanceStageDecisionInstanceDecisionInsightClient) ListMePendingAccessReviewInstanceByIdStageByIdDecisionByIdInstanceDecisionByIdInsightsCompleteMatchingPredicate(ctx context.Context, id MePendingAccessReviewInstanceStageDecisionInstanceDecisionId, predicate models.GovernanceInsightCollectionResponseOperationPredicate) (result ListMePendingAccessReviewInstanceByIdStageByIdDecisionByIdInstanceDecisionByIdInsightsCompleteResult, err error) {
	items := make([]models.GovernanceInsightCollectionResponse, 0)

	resp, err := c.ListMePendingAccessReviewInstanceByIdStageByIdDecisionByIdInstanceDecisionByIdInsights(ctx, id)
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

	result = ListMePendingAccessReviewInstanceByIdStageByIdDecisionByIdInstanceDecisionByIdInsightsCompleteResult{
		Items: items,
	}
	return
}

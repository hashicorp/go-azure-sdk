package mependingaccessreviewinstancedecisioninstancestagedecision

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

type ListMePendingAccessReviewInstanceByIdDecisionByIdInstanceStageByIdDecisionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.AccessReviewInstanceDecisionItemCollectionResponse
}

type ListMePendingAccessReviewInstanceByIdDecisionByIdInstanceStageByIdDecisionsCompleteResult struct {
	Items []models.AccessReviewInstanceDecisionItemCollectionResponse
}

// ListMePendingAccessReviewInstanceByIdDecisionByIdInstanceStageByIdDecisions ...
func (c MePendingAccessReviewInstanceDecisionInstanceStageDecisionClient) ListMePendingAccessReviewInstanceByIdDecisionByIdInstanceStageByIdDecisions(ctx context.Context, id MePendingAccessReviewInstanceDecisionInstanceStageId) (result ListMePendingAccessReviewInstanceByIdDecisionByIdInstanceStageByIdDecisionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/decisions", id.ID()),
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
		Values *[]models.AccessReviewInstanceDecisionItemCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListMePendingAccessReviewInstanceByIdDecisionByIdInstanceStageByIdDecisionsComplete retrieves all the results into a single object
func (c MePendingAccessReviewInstanceDecisionInstanceStageDecisionClient) ListMePendingAccessReviewInstanceByIdDecisionByIdInstanceStageByIdDecisionsComplete(ctx context.Context, id MePendingAccessReviewInstanceDecisionInstanceStageId) (ListMePendingAccessReviewInstanceByIdDecisionByIdInstanceStageByIdDecisionsCompleteResult, error) {
	return c.ListMePendingAccessReviewInstanceByIdDecisionByIdInstanceStageByIdDecisionsCompleteMatchingPredicate(ctx, id, models.AccessReviewInstanceDecisionItemCollectionResponseOperationPredicate{})
}

// ListMePendingAccessReviewInstanceByIdDecisionByIdInstanceStageByIdDecisionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MePendingAccessReviewInstanceDecisionInstanceStageDecisionClient) ListMePendingAccessReviewInstanceByIdDecisionByIdInstanceStageByIdDecisionsCompleteMatchingPredicate(ctx context.Context, id MePendingAccessReviewInstanceDecisionInstanceStageId, predicate models.AccessReviewInstanceDecisionItemCollectionResponseOperationPredicate) (result ListMePendingAccessReviewInstanceByIdDecisionByIdInstanceStageByIdDecisionsCompleteResult, err error) {
	items := make([]models.AccessReviewInstanceDecisionItemCollectionResponse, 0)

	resp, err := c.ListMePendingAccessReviewInstanceByIdDecisionByIdInstanceStageByIdDecisions(ctx, id)
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

	result = ListMePendingAccessReviewInstanceByIdDecisionByIdInstanceStageByIdDecisionsCompleteResult{
		Items: items,
	}
	return
}

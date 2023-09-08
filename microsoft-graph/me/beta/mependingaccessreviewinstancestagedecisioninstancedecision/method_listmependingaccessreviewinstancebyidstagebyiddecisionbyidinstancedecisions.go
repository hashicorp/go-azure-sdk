package mependingaccessreviewinstancestagedecisioninstancedecision

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

type ListMePendingAccessReviewInstanceByIdStageByIdDecisionByIdInstanceDecisionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.AccessReviewInstanceDecisionItemCollectionResponse
}

type ListMePendingAccessReviewInstanceByIdStageByIdDecisionByIdInstanceDecisionsCompleteResult struct {
	Items []models.AccessReviewInstanceDecisionItemCollectionResponse
}

// ListMePendingAccessReviewInstanceByIdStageByIdDecisionByIdInstanceDecisions ...
func (c MePendingAccessReviewInstanceStageDecisionInstanceDecisionClient) ListMePendingAccessReviewInstanceByIdStageByIdDecisionByIdInstanceDecisions(ctx context.Context, id MePendingAccessReviewInstanceStageDecisionId) (result ListMePendingAccessReviewInstanceByIdStageByIdDecisionByIdInstanceDecisionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/instance/decisions", id.ID()),
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

// ListMePendingAccessReviewInstanceByIdStageByIdDecisionByIdInstanceDecisionsComplete retrieves all the results into a single object
func (c MePendingAccessReviewInstanceStageDecisionInstanceDecisionClient) ListMePendingAccessReviewInstanceByIdStageByIdDecisionByIdInstanceDecisionsComplete(ctx context.Context, id MePendingAccessReviewInstanceStageDecisionId) (ListMePendingAccessReviewInstanceByIdStageByIdDecisionByIdInstanceDecisionsCompleteResult, error) {
	return c.ListMePendingAccessReviewInstanceByIdStageByIdDecisionByIdInstanceDecisionsCompleteMatchingPredicate(ctx, id, models.AccessReviewInstanceDecisionItemCollectionResponseOperationPredicate{})
}

// ListMePendingAccessReviewInstanceByIdStageByIdDecisionByIdInstanceDecisionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MePendingAccessReviewInstanceStageDecisionInstanceDecisionClient) ListMePendingAccessReviewInstanceByIdStageByIdDecisionByIdInstanceDecisionsCompleteMatchingPredicate(ctx context.Context, id MePendingAccessReviewInstanceStageDecisionId, predicate models.AccessReviewInstanceDecisionItemCollectionResponseOperationPredicate) (result ListMePendingAccessReviewInstanceByIdStageByIdDecisionByIdInstanceDecisionsCompleteResult, err error) {
	items := make([]models.AccessReviewInstanceDecisionItemCollectionResponse, 0)

	resp, err := c.ListMePendingAccessReviewInstanceByIdStageByIdDecisionByIdInstanceDecisions(ctx, id)
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

	result = ListMePendingAccessReviewInstanceByIdStageByIdDecisionByIdInstanceDecisionsCompleteResult{
		Items: items,
	}
	return
}

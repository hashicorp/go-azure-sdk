package userpendingaccessreviewinstancedecisioninstancestagedecision

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

type ListUserByIdPendingAccessReviewInstanceByIdDecisionByIdInstanceStageByIdDecisionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.AccessReviewInstanceDecisionItemCollectionResponse
}

type ListUserByIdPendingAccessReviewInstanceByIdDecisionByIdInstanceStageByIdDecisionsCompleteResult struct {
	Items []models.AccessReviewInstanceDecisionItemCollectionResponse
}

// ListUserByIdPendingAccessReviewInstanceByIdDecisionByIdInstanceStageByIdDecisions ...
func (c UserPendingAccessReviewInstanceDecisionInstanceStageDecisionClient) ListUserByIdPendingAccessReviewInstanceByIdDecisionByIdInstanceStageByIdDecisions(ctx context.Context, id UserPendingAccessReviewInstanceDecisionInstanceStageId) (result ListUserByIdPendingAccessReviewInstanceByIdDecisionByIdInstanceStageByIdDecisionsOperationResponse, err error) {
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

// ListUserByIdPendingAccessReviewInstanceByIdDecisionByIdInstanceStageByIdDecisionsComplete retrieves all the results into a single object
func (c UserPendingAccessReviewInstanceDecisionInstanceStageDecisionClient) ListUserByIdPendingAccessReviewInstanceByIdDecisionByIdInstanceStageByIdDecisionsComplete(ctx context.Context, id UserPendingAccessReviewInstanceDecisionInstanceStageId) (ListUserByIdPendingAccessReviewInstanceByIdDecisionByIdInstanceStageByIdDecisionsCompleteResult, error) {
	return c.ListUserByIdPendingAccessReviewInstanceByIdDecisionByIdInstanceStageByIdDecisionsCompleteMatchingPredicate(ctx, id, models.AccessReviewInstanceDecisionItemCollectionResponseOperationPredicate{})
}

// ListUserByIdPendingAccessReviewInstanceByIdDecisionByIdInstanceStageByIdDecisionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserPendingAccessReviewInstanceDecisionInstanceStageDecisionClient) ListUserByIdPendingAccessReviewInstanceByIdDecisionByIdInstanceStageByIdDecisionsCompleteMatchingPredicate(ctx context.Context, id UserPendingAccessReviewInstanceDecisionInstanceStageId, predicate models.AccessReviewInstanceDecisionItemCollectionResponseOperationPredicate) (result ListUserByIdPendingAccessReviewInstanceByIdDecisionByIdInstanceStageByIdDecisionsCompleteResult, err error) {
	items := make([]models.AccessReviewInstanceDecisionItemCollectionResponse, 0)

	resp, err := c.ListUserByIdPendingAccessReviewInstanceByIdDecisionByIdInstanceStageByIdDecisions(ctx, id)
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

	result = ListUserByIdPendingAccessReviewInstanceByIdDecisionByIdInstanceStageByIdDecisionsCompleteResult{
		Items: items,
	}
	return
}

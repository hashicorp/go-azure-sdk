package userpendingaccessreviewinstancestagedecisioninstancedecision

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

type ListUserByIdPendingAccessReviewInstanceByIdStageByIdDecisionByIdInstanceDecisionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.AccessReviewInstanceDecisionItemCollectionResponse
}

type ListUserByIdPendingAccessReviewInstanceByIdStageByIdDecisionByIdInstanceDecisionsCompleteResult struct {
	Items []models.AccessReviewInstanceDecisionItemCollectionResponse
}

// ListUserByIdPendingAccessReviewInstanceByIdStageByIdDecisionByIdInstanceDecisions ...
func (c UserPendingAccessReviewInstanceStageDecisionInstanceDecisionClient) ListUserByIdPendingAccessReviewInstanceByIdStageByIdDecisionByIdInstanceDecisions(ctx context.Context, id UserPendingAccessReviewInstanceStageDecisionId) (result ListUserByIdPendingAccessReviewInstanceByIdStageByIdDecisionByIdInstanceDecisionsOperationResponse, err error) {
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

// ListUserByIdPendingAccessReviewInstanceByIdStageByIdDecisionByIdInstanceDecisionsComplete retrieves all the results into a single object
func (c UserPendingAccessReviewInstanceStageDecisionInstanceDecisionClient) ListUserByIdPendingAccessReviewInstanceByIdStageByIdDecisionByIdInstanceDecisionsComplete(ctx context.Context, id UserPendingAccessReviewInstanceStageDecisionId) (ListUserByIdPendingAccessReviewInstanceByIdStageByIdDecisionByIdInstanceDecisionsCompleteResult, error) {
	return c.ListUserByIdPendingAccessReviewInstanceByIdStageByIdDecisionByIdInstanceDecisionsCompleteMatchingPredicate(ctx, id, models.AccessReviewInstanceDecisionItemCollectionResponseOperationPredicate{})
}

// ListUserByIdPendingAccessReviewInstanceByIdStageByIdDecisionByIdInstanceDecisionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserPendingAccessReviewInstanceStageDecisionInstanceDecisionClient) ListUserByIdPendingAccessReviewInstanceByIdStageByIdDecisionByIdInstanceDecisionsCompleteMatchingPredicate(ctx context.Context, id UserPendingAccessReviewInstanceStageDecisionId, predicate models.AccessReviewInstanceDecisionItemCollectionResponseOperationPredicate) (result ListUserByIdPendingAccessReviewInstanceByIdStageByIdDecisionByIdInstanceDecisionsCompleteResult, err error) {
	items := make([]models.AccessReviewInstanceDecisionItemCollectionResponse, 0)

	resp, err := c.ListUserByIdPendingAccessReviewInstanceByIdStageByIdDecisionByIdInstanceDecisions(ctx, id)
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

	result = ListUserByIdPendingAccessReviewInstanceByIdStageByIdDecisionByIdInstanceDecisionsCompleteResult{
		Items: items,
	}
	return
}

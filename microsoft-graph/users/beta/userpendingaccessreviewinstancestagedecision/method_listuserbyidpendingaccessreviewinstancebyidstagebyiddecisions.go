package userpendingaccessreviewinstancestagedecision

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

type ListUserByIdPendingAccessReviewInstanceByIdStageByIdDecisionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.AccessReviewInstanceDecisionItemCollectionResponse
}

type ListUserByIdPendingAccessReviewInstanceByIdStageByIdDecisionsCompleteResult struct {
	Items []models.AccessReviewInstanceDecisionItemCollectionResponse
}

// ListUserByIdPendingAccessReviewInstanceByIdStageByIdDecisions ...
func (c UserPendingAccessReviewInstanceStageDecisionClient) ListUserByIdPendingAccessReviewInstanceByIdStageByIdDecisions(ctx context.Context, id UserPendingAccessReviewInstanceStageId) (result ListUserByIdPendingAccessReviewInstanceByIdStageByIdDecisionsOperationResponse, err error) {
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

// ListUserByIdPendingAccessReviewInstanceByIdStageByIdDecisionsComplete retrieves all the results into a single object
func (c UserPendingAccessReviewInstanceStageDecisionClient) ListUserByIdPendingAccessReviewInstanceByIdStageByIdDecisionsComplete(ctx context.Context, id UserPendingAccessReviewInstanceStageId) (ListUserByIdPendingAccessReviewInstanceByIdStageByIdDecisionsCompleteResult, error) {
	return c.ListUserByIdPendingAccessReviewInstanceByIdStageByIdDecisionsCompleteMatchingPredicate(ctx, id, models.AccessReviewInstanceDecisionItemCollectionResponseOperationPredicate{})
}

// ListUserByIdPendingAccessReviewInstanceByIdStageByIdDecisionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserPendingAccessReviewInstanceStageDecisionClient) ListUserByIdPendingAccessReviewInstanceByIdStageByIdDecisionsCompleteMatchingPredicate(ctx context.Context, id UserPendingAccessReviewInstanceStageId, predicate models.AccessReviewInstanceDecisionItemCollectionResponseOperationPredicate) (result ListUserByIdPendingAccessReviewInstanceByIdStageByIdDecisionsCompleteResult, err error) {
	items := make([]models.AccessReviewInstanceDecisionItemCollectionResponse, 0)

	resp, err := c.ListUserByIdPendingAccessReviewInstanceByIdStageByIdDecisions(ctx, id)
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

	result = ListUserByIdPendingAccessReviewInstanceByIdStageByIdDecisionsCompleteResult{
		Items: items,
	}
	return
}

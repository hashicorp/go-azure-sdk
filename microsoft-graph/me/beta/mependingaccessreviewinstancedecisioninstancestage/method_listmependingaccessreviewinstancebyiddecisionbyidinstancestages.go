package mependingaccessreviewinstancedecisioninstancestage

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

type ListMePendingAccessReviewInstanceByIdDecisionByIdInstanceStagesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.AccessReviewStageCollectionResponse
}

type ListMePendingAccessReviewInstanceByIdDecisionByIdInstanceStagesCompleteResult struct {
	Items []models.AccessReviewStageCollectionResponse
}

// ListMePendingAccessReviewInstanceByIdDecisionByIdInstanceStages ...
func (c MePendingAccessReviewInstanceDecisionInstanceStageClient) ListMePendingAccessReviewInstanceByIdDecisionByIdInstanceStages(ctx context.Context, id MePendingAccessReviewInstanceDecisionId) (result ListMePendingAccessReviewInstanceByIdDecisionByIdInstanceStagesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/instance/stages", id.ID()),
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
		Values *[]models.AccessReviewStageCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListMePendingAccessReviewInstanceByIdDecisionByIdInstanceStagesComplete retrieves all the results into a single object
func (c MePendingAccessReviewInstanceDecisionInstanceStageClient) ListMePendingAccessReviewInstanceByIdDecisionByIdInstanceStagesComplete(ctx context.Context, id MePendingAccessReviewInstanceDecisionId) (ListMePendingAccessReviewInstanceByIdDecisionByIdInstanceStagesCompleteResult, error) {
	return c.ListMePendingAccessReviewInstanceByIdDecisionByIdInstanceStagesCompleteMatchingPredicate(ctx, id, models.AccessReviewStageCollectionResponseOperationPredicate{})
}

// ListMePendingAccessReviewInstanceByIdDecisionByIdInstanceStagesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MePendingAccessReviewInstanceDecisionInstanceStageClient) ListMePendingAccessReviewInstanceByIdDecisionByIdInstanceStagesCompleteMatchingPredicate(ctx context.Context, id MePendingAccessReviewInstanceDecisionId, predicate models.AccessReviewStageCollectionResponseOperationPredicate) (result ListMePendingAccessReviewInstanceByIdDecisionByIdInstanceStagesCompleteResult, err error) {
	items := make([]models.AccessReviewStageCollectionResponse, 0)

	resp, err := c.ListMePendingAccessReviewInstanceByIdDecisionByIdInstanceStages(ctx, id)
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

	result = ListMePendingAccessReviewInstanceByIdDecisionByIdInstanceStagesCompleteResult{
		Items: items,
	}
	return
}

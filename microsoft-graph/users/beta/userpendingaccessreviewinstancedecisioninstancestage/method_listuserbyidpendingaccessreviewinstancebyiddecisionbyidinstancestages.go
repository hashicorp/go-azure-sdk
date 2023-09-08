package userpendingaccessreviewinstancedecisioninstancestage

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

type ListUserByIdPendingAccessReviewInstanceByIdDecisionByIdInstanceStagesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.AccessReviewStageCollectionResponse
}

type ListUserByIdPendingAccessReviewInstanceByIdDecisionByIdInstanceStagesCompleteResult struct {
	Items []models.AccessReviewStageCollectionResponse
}

// ListUserByIdPendingAccessReviewInstanceByIdDecisionByIdInstanceStages ...
func (c UserPendingAccessReviewInstanceDecisionInstanceStageClient) ListUserByIdPendingAccessReviewInstanceByIdDecisionByIdInstanceStages(ctx context.Context, id UserPendingAccessReviewInstanceDecisionId) (result ListUserByIdPendingAccessReviewInstanceByIdDecisionByIdInstanceStagesOperationResponse, err error) {
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

// ListUserByIdPendingAccessReviewInstanceByIdDecisionByIdInstanceStagesComplete retrieves all the results into a single object
func (c UserPendingAccessReviewInstanceDecisionInstanceStageClient) ListUserByIdPendingAccessReviewInstanceByIdDecisionByIdInstanceStagesComplete(ctx context.Context, id UserPendingAccessReviewInstanceDecisionId) (ListUserByIdPendingAccessReviewInstanceByIdDecisionByIdInstanceStagesCompleteResult, error) {
	return c.ListUserByIdPendingAccessReviewInstanceByIdDecisionByIdInstanceStagesCompleteMatchingPredicate(ctx, id, models.AccessReviewStageCollectionResponseOperationPredicate{})
}

// ListUserByIdPendingAccessReviewInstanceByIdDecisionByIdInstanceStagesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserPendingAccessReviewInstanceDecisionInstanceStageClient) ListUserByIdPendingAccessReviewInstanceByIdDecisionByIdInstanceStagesCompleteMatchingPredicate(ctx context.Context, id UserPendingAccessReviewInstanceDecisionId, predicate models.AccessReviewStageCollectionResponseOperationPredicate) (result ListUserByIdPendingAccessReviewInstanceByIdDecisionByIdInstanceStagesCompleteResult, err error) {
	items := make([]models.AccessReviewStageCollectionResponse, 0)

	resp, err := c.ListUserByIdPendingAccessReviewInstanceByIdDecisionByIdInstanceStages(ctx, id)
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

	result = ListUserByIdPendingAccessReviewInstanceByIdDecisionByIdInstanceStagesCompleteResult{
		Items: items,
	}
	return
}

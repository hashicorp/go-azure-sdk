package mependingaccessreviewinstancestagedecisioninstancecontactedreviewer

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

type ListMePendingAccessReviewInstanceByIdStageByIdDecisionByIdInstanceContactedReviewersOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.AccessReviewReviewerCollectionResponse
}

type ListMePendingAccessReviewInstanceByIdStageByIdDecisionByIdInstanceContactedReviewersCompleteResult struct {
	Items []models.AccessReviewReviewerCollectionResponse
}

// ListMePendingAccessReviewInstanceByIdStageByIdDecisionByIdInstanceContactedReviewers ...
func (c MePendingAccessReviewInstanceStageDecisionInstanceContactedReviewerClient) ListMePendingAccessReviewInstanceByIdStageByIdDecisionByIdInstanceContactedReviewers(ctx context.Context, id MePendingAccessReviewInstanceStageDecisionId) (result ListMePendingAccessReviewInstanceByIdStageByIdDecisionByIdInstanceContactedReviewersOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/instance/contactedReviewers", id.ID()),
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
		Values *[]models.AccessReviewReviewerCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListMePendingAccessReviewInstanceByIdStageByIdDecisionByIdInstanceContactedReviewersComplete retrieves all the results into a single object
func (c MePendingAccessReviewInstanceStageDecisionInstanceContactedReviewerClient) ListMePendingAccessReviewInstanceByIdStageByIdDecisionByIdInstanceContactedReviewersComplete(ctx context.Context, id MePendingAccessReviewInstanceStageDecisionId) (ListMePendingAccessReviewInstanceByIdStageByIdDecisionByIdInstanceContactedReviewersCompleteResult, error) {
	return c.ListMePendingAccessReviewInstanceByIdStageByIdDecisionByIdInstanceContactedReviewersCompleteMatchingPredicate(ctx, id, models.AccessReviewReviewerCollectionResponseOperationPredicate{})
}

// ListMePendingAccessReviewInstanceByIdStageByIdDecisionByIdInstanceContactedReviewersCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MePendingAccessReviewInstanceStageDecisionInstanceContactedReviewerClient) ListMePendingAccessReviewInstanceByIdStageByIdDecisionByIdInstanceContactedReviewersCompleteMatchingPredicate(ctx context.Context, id MePendingAccessReviewInstanceStageDecisionId, predicate models.AccessReviewReviewerCollectionResponseOperationPredicate) (result ListMePendingAccessReviewInstanceByIdStageByIdDecisionByIdInstanceContactedReviewersCompleteResult, err error) {
	items := make([]models.AccessReviewReviewerCollectionResponse, 0)

	resp, err := c.ListMePendingAccessReviewInstanceByIdStageByIdDecisionByIdInstanceContactedReviewers(ctx, id)
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

	result = ListMePendingAccessReviewInstanceByIdStageByIdDecisionByIdInstanceContactedReviewersCompleteResult{
		Items: items,
	}
	return
}

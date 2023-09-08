package mependingaccessreviewinstancedecision

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

type ListMePendingAccessReviewInstanceByIdDecisionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.AccessReviewInstanceDecisionItemCollectionResponse
}

type ListMePendingAccessReviewInstanceByIdDecisionsCompleteResult struct {
	Items []models.AccessReviewInstanceDecisionItemCollectionResponse
}

// ListMePendingAccessReviewInstanceByIdDecisions ...
func (c MePendingAccessReviewInstanceDecisionClient) ListMePendingAccessReviewInstanceByIdDecisions(ctx context.Context, id MePendingAccessReviewInstanceId) (result ListMePendingAccessReviewInstanceByIdDecisionsOperationResponse, err error) {
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

// ListMePendingAccessReviewInstanceByIdDecisionsComplete retrieves all the results into a single object
func (c MePendingAccessReviewInstanceDecisionClient) ListMePendingAccessReviewInstanceByIdDecisionsComplete(ctx context.Context, id MePendingAccessReviewInstanceId) (ListMePendingAccessReviewInstanceByIdDecisionsCompleteResult, error) {
	return c.ListMePendingAccessReviewInstanceByIdDecisionsCompleteMatchingPredicate(ctx, id, models.AccessReviewInstanceDecisionItemCollectionResponseOperationPredicate{})
}

// ListMePendingAccessReviewInstanceByIdDecisionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MePendingAccessReviewInstanceDecisionClient) ListMePendingAccessReviewInstanceByIdDecisionsCompleteMatchingPredicate(ctx context.Context, id MePendingAccessReviewInstanceId, predicate models.AccessReviewInstanceDecisionItemCollectionResponseOperationPredicate) (result ListMePendingAccessReviewInstanceByIdDecisionsCompleteResult, err error) {
	items := make([]models.AccessReviewInstanceDecisionItemCollectionResponse, 0)

	resp, err := c.ListMePendingAccessReviewInstanceByIdDecisions(ctx, id)
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

	result = ListMePendingAccessReviewInstanceByIdDecisionsCompleteResult{
		Items: items,
	}
	return
}

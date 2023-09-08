package userpendingaccessreviewinstancedecision

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

type ListUserByIdPendingAccessReviewInstanceByIdDecisionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.AccessReviewInstanceDecisionItemCollectionResponse
}

type ListUserByIdPendingAccessReviewInstanceByIdDecisionsCompleteResult struct {
	Items []models.AccessReviewInstanceDecisionItemCollectionResponse
}

// ListUserByIdPendingAccessReviewInstanceByIdDecisions ...
func (c UserPendingAccessReviewInstanceDecisionClient) ListUserByIdPendingAccessReviewInstanceByIdDecisions(ctx context.Context, id UserPendingAccessReviewInstanceId) (result ListUserByIdPendingAccessReviewInstanceByIdDecisionsOperationResponse, err error) {
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

// ListUserByIdPendingAccessReviewInstanceByIdDecisionsComplete retrieves all the results into a single object
func (c UserPendingAccessReviewInstanceDecisionClient) ListUserByIdPendingAccessReviewInstanceByIdDecisionsComplete(ctx context.Context, id UserPendingAccessReviewInstanceId) (ListUserByIdPendingAccessReviewInstanceByIdDecisionsCompleteResult, error) {
	return c.ListUserByIdPendingAccessReviewInstanceByIdDecisionsCompleteMatchingPredicate(ctx, id, models.AccessReviewInstanceDecisionItemCollectionResponseOperationPredicate{})
}

// ListUserByIdPendingAccessReviewInstanceByIdDecisionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserPendingAccessReviewInstanceDecisionClient) ListUserByIdPendingAccessReviewInstanceByIdDecisionsCompleteMatchingPredicate(ctx context.Context, id UserPendingAccessReviewInstanceId, predicate models.AccessReviewInstanceDecisionItemCollectionResponseOperationPredicate) (result ListUserByIdPendingAccessReviewInstanceByIdDecisionsCompleteResult, err error) {
	items := make([]models.AccessReviewInstanceDecisionItemCollectionResponse, 0)

	resp, err := c.ListUserByIdPendingAccessReviewInstanceByIdDecisions(ctx, id)
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

	result = ListUserByIdPendingAccessReviewInstanceByIdDecisionsCompleteResult{
		Items: items,
	}
	return
}

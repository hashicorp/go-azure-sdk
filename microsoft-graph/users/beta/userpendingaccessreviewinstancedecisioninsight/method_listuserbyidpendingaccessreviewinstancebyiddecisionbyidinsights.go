package userpendingaccessreviewinstancedecisioninsight

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

type ListUserByIdPendingAccessReviewInstanceByIdDecisionByIdInsightsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.GovernanceInsightCollectionResponse
}

type ListUserByIdPendingAccessReviewInstanceByIdDecisionByIdInsightsCompleteResult struct {
	Items []models.GovernanceInsightCollectionResponse
}

// ListUserByIdPendingAccessReviewInstanceByIdDecisionByIdInsights ...
func (c UserPendingAccessReviewInstanceDecisionInsightClient) ListUserByIdPendingAccessReviewInstanceByIdDecisionByIdInsights(ctx context.Context, id UserPendingAccessReviewInstanceDecisionId) (result ListUserByIdPendingAccessReviewInstanceByIdDecisionByIdInsightsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/insights", id.ID()),
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
		Values *[]models.GovernanceInsightCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserByIdPendingAccessReviewInstanceByIdDecisionByIdInsightsComplete retrieves all the results into a single object
func (c UserPendingAccessReviewInstanceDecisionInsightClient) ListUserByIdPendingAccessReviewInstanceByIdDecisionByIdInsightsComplete(ctx context.Context, id UserPendingAccessReviewInstanceDecisionId) (ListUserByIdPendingAccessReviewInstanceByIdDecisionByIdInsightsCompleteResult, error) {
	return c.ListUserByIdPendingAccessReviewInstanceByIdDecisionByIdInsightsCompleteMatchingPredicate(ctx, id, models.GovernanceInsightCollectionResponseOperationPredicate{})
}

// ListUserByIdPendingAccessReviewInstanceByIdDecisionByIdInsightsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserPendingAccessReviewInstanceDecisionInsightClient) ListUserByIdPendingAccessReviewInstanceByIdDecisionByIdInsightsCompleteMatchingPredicate(ctx context.Context, id UserPendingAccessReviewInstanceDecisionId, predicate models.GovernanceInsightCollectionResponseOperationPredicate) (result ListUserByIdPendingAccessReviewInstanceByIdDecisionByIdInsightsCompleteResult, err error) {
	items := make([]models.GovernanceInsightCollectionResponse, 0)

	resp, err := c.ListUserByIdPendingAccessReviewInstanceByIdDecisionByIdInsights(ctx, id)
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

	result = ListUserByIdPendingAccessReviewInstanceByIdDecisionByIdInsightsCompleteResult{
		Items: items,
	}
	return
}

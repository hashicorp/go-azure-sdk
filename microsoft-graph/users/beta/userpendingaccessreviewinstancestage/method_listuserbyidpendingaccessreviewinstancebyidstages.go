package userpendingaccessreviewinstancestage

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

type ListUserByIdPendingAccessReviewInstanceByIdStagesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.AccessReviewStageCollectionResponse
}

type ListUserByIdPendingAccessReviewInstanceByIdStagesCompleteResult struct {
	Items []models.AccessReviewStageCollectionResponse
}

// ListUserByIdPendingAccessReviewInstanceByIdStages ...
func (c UserPendingAccessReviewInstanceStageClient) ListUserByIdPendingAccessReviewInstanceByIdStages(ctx context.Context, id UserPendingAccessReviewInstanceId) (result ListUserByIdPendingAccessReviewInstanceByIdStagesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/stages", id.ID()),
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

// ListUserByIdPendingAccessReviewInstanceByIdStagesComplete retrieves all the results into a single object
func (c UserPendingAccessReviewInstanceStageClient) ListUserByIdPendingAccessReviewInstanceByIdStagesComplete(ctx context.Context, id UserPendingAccessReviewInstanceId) (ListUserByIdPendingAccessReviewInstanceByIdStagesCompleteResult, error) {
	return c.ListUserByIdPendingAccessReviewInstanceByIdStagesCompleteMatchingPredicate(ctx, id, models.AccessReviewStageCollectionResponseOperationPredicate{})
}

// ListUserByIdPendingAccessReviewInstanceByIdStagesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserPendingAccessReviewInstanceStageClient) ListUserByIdPendingAccessReviewInstanceByIdStagesCompleteMatchingPredicate(ctx context.Context, id UserPendingAccessReviewInstanceId, predicate models.AccessReviewStageCollectionResponseOperationPredicate) (result ListUserByIdPendingAccessReviewInstanceByIdStagesCompleteResult, err error) {
	items := make([]models.AccessReviewStageCollectionResponse, 0)

	resp, err := c.ListUserByIdPendingAccessReviewInstanceByIdStages(ctx, id)
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

	result = ListUserByIdPendingAccessReviewInstanceByIdStagesCompleteResult{
		Items: items,
	}
	return
}

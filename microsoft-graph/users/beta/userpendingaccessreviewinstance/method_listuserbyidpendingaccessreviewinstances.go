package userpendingaccessreviewinstance

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

type ListUserByIdPendingAccessReviewInstancesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.AccessReviewInstanceCollectionResponse
}

type ListUserByIdPendingAccessReviewInstancesCompleteResult struct {
	Items []models.AccessReviewInstanceCollectionResponse
}

// ListUserByIdPendingAccessReviewInstances ...
func (c UserPendingAccessReviewInstanceClient) ListUserByIdPendingAccessReviewInstances(ctx context.Context, id UserId) (result ListUserByIdPendingAccessReviewInstancesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/pendingAccessReviewInstances", id.ID()),
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
		Values *[]models.AccessReviewInstanceCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserByIdPendingAccessReviewInstancesComplete retrieves all the results into a single object
func (c UserPendingAccessReviewInstanceClient) ListUserByIdPendingAccessReviewInstancesComplete(ctx context.Context, id UserId) (ListUserByIdPendingAccessReviewInstancesCompleteResult, error) {
	return c.ListUserByIdPendingAccessReviewInstancesCompleteMatchingPredicate(ctx, id, models.AccessReviewInstanceCollectionResponseOperationPredicate{})
}

// ListUserByIdPendingAccessReviewInstancesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserPendingAccessReviewInstanceClient) ListUserByIdPendingAccessReviewInstancesCompleteMatchingPredicate(ctx context.Context, id UserId, predicate models.AccessReviewInstanceCollectionResponseOperationPredicate) (result ListUserByIdPendingAccessReviewInstancesCompleteResult, err error) {
	items := make([]models.AccessReviewInstanceCollectionResponse, 0)

	resp, err := c.ListUserByIdPendingAccessReviewInstances(ctx, id)
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

	result = ListUserByIdPendingAccessReviewInstancesCompleteResult{
		Items: items,
	}
	return
}

package mependingaccessreviewinstance

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

type ListMePendingAccessReviewInstancesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.AccessReviewInstanceCollectionResponse
}

type ListMePendingAccessReviewInstancesCompleteResult struct {
	Items []models.AccessReviewInstanceCollectionResponse
}

// ListMePendingAccessReviewInstances ...
func (c MePendingAccessReviewInstanceClient) ListMePendingAccessReviewInstances(ctx context.Context) (result ListMePendingAccessReviewInstancesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       "/me/pendingAccessReviewInstances",
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

// ListMePendingAccessReviewInstancesComplete retrieves all the results into a single object
func (c MePendingAccessReviewInstanceClient) ListMePendingAccessReviewInstancesComplete(ctx context.Context) (ListMePendingAccessReviewInstancesCompleteResult, error) {
	return c.ListMePendingAccessReviewInstancesCompleteMatchingPredicate(ctx, models.AccessReviewInstanceCollectionResponseOperationPredicate{})
}

// ListMePendingAccessReviewInstancesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MePendingAccessReviewInstanceClient) ListMePendingAccessReviewInstancesCompleteMatchingPredicate(ctx context.Context, predicate models.AccessReviewInstanceCollectionResponseOperationPredicate) (result ListMePendingAccessReviewInstancesCompleteResult, err error) {
	items := make([]models.AccessReviewInstanceCollectionResponse, 0)

	resp, err := c.ListMePendingAccessReviewInstances(ctx)
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

	result = ListMePendingAccessReviewInstancesCompleteResult{
		Items: items,
	}
	return
}

package userplannerplanbucket

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

type ListUserByIdPlannerPlanByIdBucketsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.PlannerBucketCollectionResponse
}

type ListUserByIdPlannerPlanByIdBucketsCompleteResult struct {
	Items []models.PlannerBucketCollectionResponse
}

// ListUserByIdPlannerPlanByIdBuckets ...
func (c UserPlannerPlanBucketClient) ListUserByIdPlannerPlanByIdBuckets(ctx context.Context, id UserPlannerPlanId) (result ListUserByIdPlannerPlanByIdBucketsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/buckets", id.ID()),
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
		Values *[]models.PlannerBucketCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserByIdPlannerPlanByIdBucketsComplete retrieves all the results into a single object
func (c UserPlannerPlanBucketClient) ListUserByIdPlannerPlanByIdBucketsComplete(ctx context.Context, id UserPlannerPlanId) (ListUserByIdPlannerPlanByIdBucketsCompleteResult, error) {
	return c.ListUserByIdPlannerPlanByIdBucketsCompleteMatchingPredicate(ctx, id, models.PlannerBucketCollectionResponseOperationPredicate{})
}

// ListUserByIdPlannerPlanByIdBucketsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserPlannerPlanBucketClient) ListUserByIdPlannerPlanByIdBucketsCompleteMatchingPredicate(ctx context.Context, id UserPlannerPlanId, predicate models.PlannerBucketCollectionResponseOperationPredicate) (result ListUserByIdPlannerPlanByIdBucketsCompleteResult, err error) {
	items := make([]models.PlannerBucketCollectionResponse, 0)

	resp, err := c.ListUserByIdPlannerPlanByIdBuckets(ctx, id)
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

	result = ListUserByIdPlannerPlanByIdBucketsCompleteResult{
		Items: items,
	}
	return
}

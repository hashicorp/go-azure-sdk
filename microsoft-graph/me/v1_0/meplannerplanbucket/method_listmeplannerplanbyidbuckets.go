package meplannerplanbucket

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common/v1_0/models"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListMePlannerPlanByIdBucketsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.PlannerBucketCollectionResponse
}

type ListMePlannerPlanByIdBucketsCompleteResult struct {
	Items []models.PlannerBucketCollectionResponse
}

// ListMePlannerPlanByIdBuckets ...
func (c MePlannerPlanBucketClient) ListMePlannerPlanByIdBuckets(ctx context.Context, id MePlannerPlanId) (result ListMePlannerPlanByIdBucketsOperationResponse, err error) {
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

// ListMePlannerPlanByIdBucketsComplete retrieves all the results into a single object
func (c MePlannerPlanBucketClient) ListMePlannerPlanByIdBucketsComplete(ctx context.Context, id MePlannerPlanId) (ListMePlannerPlanByIdBucketsCompleteResult, error) {
	return c.ListMePlannerPlanByIdBucketsCompleteMatchingPredicate(ctx, id, models.PlannerBucketCollectionResponseOperationPredicate{})
}

// ListMePlannerPlanByIdBucketsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MePlannerPlanBucketClient) ListMePlannerPlanByIdBucketsCompleteMatchingPredicate(ctx context.Context, id MePlannerPlanId, predicate models.PlannerBucketCollectionResponseOperationPredicate) (result ListMePlannerPlanByIdBucketsCompleteResult, err error) {
	items := make([]models.PlannerBucketCollectionResponse, 0)

	resp, err := c.ListMePlannerPlanByIdBuckets(ctx, id)
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

	result = ListMePlannerPlanByIdBucketsCompleteResult{
		Items: items,
	}
	return
}

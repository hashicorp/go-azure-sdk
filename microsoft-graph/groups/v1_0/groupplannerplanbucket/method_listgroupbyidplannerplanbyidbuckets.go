package groupplannerplanbucket

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

type ListGroupByIdPlannerPlanByIdBucketsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.PlannerBucketCollectionResponse
}

type ListGroupByIdPlannerPlanByIdBucketsCompleteResult struct {
	Items []models.PlannerBucketCollectionResponse
}

// ListGroupByIdPlannerPlanByIdBuckets ...
func (c GroupPlannerPlanBucketClient) ListGroupByIdPlannerPlanByIdBuckets(ctx context.Context, id GroupPlannerPlanId) (result ListGroupByIdPlannerPlanByIdBucketsOperationResponse, err error) {
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

// ListGroupByIdPlannerPlanByIdBucketsComplete retrieves all the results into a single object
func (c GroupPlannerPlanBucketClient) ListGroupByIdPlannerPlanByIdBucketsComplete(ctx context.Context, id GroupPlannerPlanId) (ListGroupByIdPlannerPlanByIdBucketsCompleteResult, error) {
	return c.ListGroupByIdPlannerPlanByIdBucketsCompleteMatchingPredicate(ctx, id, models.PlannerBucketCollectionResponseOperationPredicate{})
}

// ListGroupByIdPlannerPlanByIdBucketsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GroupPlannerPlanBucketClient) ListGroupByIdPlannerPlanByIdBucketsCompleteMatchingPredicate(ctx context.Context, id GroupPlannerPlanId, predicate models.PlannerBucketCollectionResponseOperationPredicate) (result ListGroupByIdPlannerPlanByIdBucketsCompleteResult, err error) {
	items := make([]models.PlannerBucketCollectionResponse, 0)

	resp, err := c.ListGroupByIdPlannerPlanByIdBuckets(ctx, id)
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

	result = ListGroupByIdPlannerPlanByIdBucketsCompleteResult{
		Items: items,
	}
	return
}

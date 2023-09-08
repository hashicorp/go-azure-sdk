package meplannerplanbuckettask

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

type ListMePlannerPlanByIdBucketByIdTasksOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.PlannerTaskCollectionResponse
}

type ListMePlannerPlanByIdBucketByIdTasksCompleteResult struct {
	Items []models.PlannerTaskCollectionResponse
}

// ListMePlannerPlanByIdBucketByIdTasks ...
func (c MePlannerPlanBucketTaskClient) ListMePlannerPlanByIdBucketByIdTasks(ctx context.Context, id MePlannerPlanBucketId) (result ListMePlannerPlanByIdBucketByIdTasksOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/tasks", id.ID()),
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
		Values *[]models.PlannerTaskCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListMePlannerPlanByIdBucketByIdTasksComplete retrieves all the results into a single object
func (c MePlannerPlanBucketTaskClient) ListMePlannerPlanByIdBucketByIdTasksComplete(ctx context.Context, id MePlannerPlanBucketId) (ListMePlannerPlanByIdBucketByIdTasksCompleteResult, error) {
	return c.ListMePlannerPlanByIdBucketByIdTasksCompleteMatchingPredicate(ctx, id, models.PlannerTaskCollectionResponseOperationPredicate{})
}

// ListMePlannerPlanByIdBucketByIdTasksCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MePlannerPlanBucketTaskClient) ListMePlannerPlanByIdBucketByIdTasksCompleteMatchingPredicate(ctx context.Context, id MePlannerPlanBucketId, predicate models.PlannerTaskCollectionResponseOperationPredicate) (result ListMePlannerPlanByIdBucketByIdTasksCompleteResult, err error) {
	items := make([]models.PlannerTaskCollectionResponse, 0)

	resp, err := c.ListMePlannerPlanByIdBucketByIdTasks(ctx, id)
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

	result = ListMePlannerPlanByIdBucketByIdTasksCompleteResult{
		Items: items,
	}
	return
}

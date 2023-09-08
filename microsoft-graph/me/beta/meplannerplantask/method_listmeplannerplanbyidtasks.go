package meplannerplantask

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

type ListMePlannerPlanByIdTasksOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.PlannerTaskCollectionResponse
}

type ListMePlannerPlanByIdTasksCompleteResult struct {
	Items []models.PlannerTaskCollectionResponse
}

// ListMePlannerPlanByIdTasks ...
func (c MePlannerPlanTaskClient) ListMePlannerPlanByIdTasks(ctx context.Context, id MePlannerPlanId) (result ListMePlannerPlanByIdTasksOperationResponse, err error) {
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

// ListMePlannerPlanByIdTasksComplete retrieves all the results into a single object
func (c MePlannerPlanTaskClient) ListMePlannerPlanByIdTasksComplete(ctx context.Context, id MePlannerPlanId) (ListMePlannerPlanByIdTasksCompleteResult, error) {
	return c.ListMePlannerPlanByIdTasksCompleteMatchingPredicate(ctx, id, models.PlannerTaskCollectionResponseOperationPredicate{})
}

// ListMePlannerPlanByIdTasksCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MePlannerPlanTaskClient) ListMePlannerPlanByIdTasksCompleteMatchingPredicate(ctx context.Context, id MePlannerPlanId, predicate models.PlannerTaskCollectionResponseOperationPredicate) (result ListMePlannerPlanByIdTasksCompleteResult, err error) {
	items := make([]models.PlannerTaskCollectionResponse, 0)

	resp, err := c.ListMePlannerPlanByIdTasks(ctx, id)
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

	result = ListMePlannerPlanByIdTasksCompleteResult{
		Items: items,
	}
	return
}

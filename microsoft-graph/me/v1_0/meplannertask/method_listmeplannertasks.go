package meplannertask

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

type ListMePlannerTasksOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.PlannerTaskCollectionResponse
}

type ListMePlannerTasksCompleteResult struct {
	Items []models.PlannerTaskCollectionResponse
}

// ListMePlannerTasks ...
func (c MePlannerTaskClient) ListMePlannerTasks(ctx context.Context) (result ListMePlannerTasksOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       "/me/planner/tasks",
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

// ListMePlannerTasksComplete retrieves all the results into a single object
func (c MePlannerTaskClient) ListMePlannerTasksComplete(ctx context.Context) (ListMePlannerTasksCompleteResult, error) {
	return c.ListMePlannerTasksCompleteMatchingPredicate(ctx, models.PlannerTaskCollectionResponseOperationPredicate{})
}

// ListMePlannerTasksCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MePlannerTaskClient) ListMePlannerTasksCompleteMatchingPredicate(ctx context.Context, predicate models.PlannerTaskCollectionResponseOperationPredicate) (result ListMePlannerTasksCompleteResult, err error) {
	items := make([]models.PlannerTaskCollectionResponse, 0)

	resp, err := c.ListMePlannerTasks(ctx)
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

	result = ListMePlannerTasksCompleteResult{
		Items: items,
	}
	return
}

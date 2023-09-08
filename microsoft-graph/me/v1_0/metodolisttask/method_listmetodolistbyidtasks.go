package metodolisttask

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

type ListMeTodoListByIdTasksOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.TodoTaskCollectionResponse
}

type ListMeTodoListByIdTasksCompleteResult struct {
	Items []models.TodoTaskCollectionResponse
}

// ListMeTodoListByIdTasks ...
func (c MeTodoListTaskClient) ListMeTodoListByIdTasks(ctx context.Context, id MeTodoListId) (result ListMeTodoListByIdTasksOperationResponse, err error) {
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
		Values *[]models.TodoTaskCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListMeTodoListByIdTasksComplete retrieves all the results into a single object
func (c MeTodoListTaskClient) ListMeTodoListByIdTasksComplete(ctx context.Context, id MeTodoListId) (ListMeTodoListByIdTasksCompleteResult, error) {
	return c.ListMeTodoListByIdTasksCompleteMatchingPredicate(ctx, id, models.TodoTaskCollectionResponseOperationPredicate{})
}

// ListMeTodoListByIdTasksCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeTodoListTaskClient) ListMeTodoListByIdTasksCompleteMatchingPredicate(ctx context.Context, id MeTodoListId, predicate models.TodoTaskCollectionResponseOperationPredicate) (result ListMeTodoListByIdTasksCompleteResult, err error) {
	items := make([]models.TodoTaskCollectionResponse, 0)

	resp, err := c.ListMeTodoListByIdTasks(ctx, id)
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

	result = ListMeTodoListByIdTasksCompleteResult{
		Items: items,
	}
	return
}

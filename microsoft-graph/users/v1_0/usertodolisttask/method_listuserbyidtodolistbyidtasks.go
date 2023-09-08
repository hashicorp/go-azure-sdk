package usertodolisttask

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

type ListUserByIdTodoListByIdTasksOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.TodoTaskCollectionResponse
}

type ListUserByIdTodoListByIdTasksCompleteResult struct {
	Items []models.TodoTaskCollectionResponse
}

// ListUserByIdTodoListByIdTasks ...
func (c UserTodoListTaskClient) ListUserByIdTodoListByIdTasks(ctx context.Context, id UserTodoListId) (result ListUserByIdTodoListByIdTasksOperationResponse, err error) {
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

// ListUserByIdTodoListByIdTasksComplete retrieves all the results into a single object
func (c UserTodoListTaskClient) ListUserByIdTodoListByIdTasksComplete(ctx context.Context, id UserTodoListId) (ListUserByIdTodoListByIdTasksCompleteResult, error) {
	return c.ListUserByIdTodoListByIdTasksCompleteMatchingPredicate(ctx, id, models.TodoTaskCollectionResponseOperationPredicate{})
}

// ListUserByIdTodoListByIdTasksCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserTodoListTaskClient) ListUserByIdTodoListByIdTasksCompleteMatchingPredicate(ctx context.Context, id UserTodoListId, predicate models.TodoTaskCollectionResponseOperationPredicate) (result ListUserByIdTodoListByIdTasksCompleteResult, err error) {
	items := make([]models.TodoTaskCollectionResponse, 0)

	resp, err := c.ListUserByIdTodoListByIdTasks(ctx, id)
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

	result = ListUserByIdTodoListByIdTasksCompleteResult{
		Items: items,
	}
	return
}

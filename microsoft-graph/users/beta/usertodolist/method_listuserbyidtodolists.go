package usertodolist

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

type ListUserByIdTodoListsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.TodoTaskListCollectionResponse
}

type ListUserByIdTodoListsCompleteResult struct {
	Items []models.TodoTaskListCollectionResponse
}

// ListUserByIdTodoLists ...
func (c UserTodoListClient) ListUserByIdTodoLists(ctx context.Context, id UserId) (result ListUserByIdTodoListsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/todo/lists", id.ID()),
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
		Values *[]models.TodoTaskListCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserByIdTodoListsComplete retrieves all the results into a single object
func (c UserTodoListClient) ListUserByIdTodoListsComplete(ctx context.Context, id UserId) (ListUserByIdTodoListsCompleteResult, error) {
	return c.ListUserByIdTodoListsCompleteMatchingPredicate(ctx, id, models.TodoTaskListCollectionResponseOperationPredicate{})
}

// ListUserByIdTodoListsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserTodoListClient) ListUserByIdTodoListsCompleteMatchingPredicate(ctx context.Context, id UserId, predicate models.TodoTaskListCollectionResponseOperationPredicate) (result ListUserByIdTodoListsCompleteResult, err error) {
	items := make([]models.TodoTaskListCollectionResponse, 0)

	resp, err := c.ListUserByIdTodoLists(ctx, id)
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

	result = ListUserByIdTodoListsCompleteResult{
		Items: items,
	}
	return
}

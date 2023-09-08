package usertodolisttaskchecklistitem

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

type ListUserByIdTodoListByIdTaskByIdChecklistItemsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.ChecklistItemCollectionResponse
}

type ListUserByIdTodoListByIdTaskByIdChecklistItemsCompleteResult struct {
	Items []models.ChecklistItemCollectionResponse
}

// ListUserByIdTodoListByIdTaskByIdChecklistItems ...
func (c UserTodoListTaskChecklistItemClient) ListUserByIdTodoListByIdTaskByIdChecklistItems(ctx context.Context, id UserTodoListTaskId) (result ListUserByIdTodoListByIdTaskByIdChecklistItemsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/checklistItems", id.ID()),
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
		Values *[]models.ChecklistItemCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserByIdTodoListByIdTaskByIdChecklistItemsComplete retrieves all the results into a single object
func (c UserTodoListTaskChecklistItemClient) ListUserByIdTodoListByIdTaskByIdChecklistItemsComplete(ctx context.Context, id UserTodoListTaskId) (ListUserByIdTodoListByIdTaskByIdChecklistItemsCompleteResult, error) {
	return c.ListUserByIdTodoListByIdTaskByIdChecklistItemsCompleteMatchingPredicate(ctx, id, models.ChecklistItemCollectionResponseOperationPredicate{})
}

// ListUserByIdTodoListByIdTaskByIdChecklistItemsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserTodoListTaskChecklistItemClient) ListUserByIdTodoListByIdTaskByIdChecklistItemsCompleteMatchingPredicate(ctx context.Context, id UserTodoListTaskId, predicate models.ChecklistItemCollectionResponseOperationPredicate) (result ListUserByIdTodoListByIdTaskByIdChecklistItemsCompleteResult, err error) {
	items := make([]models.ChecklistItemCollectionResponse, 0)

	resp, err := c.ListUserByIdTodoListByIdTaskByIdChecklistItems(ctx, id)
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

	result = ListUserByIdTodoListByIdTaskByIdChecklistItemsCompleteResult{
		Items: items,
	}
	return
}

package metodolisttaskchecklistitem

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

type ListMeTodoListByIdTaskByIdChecklistItemsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.ChecklistItemCollectionResponse
}

type ListMeTodoListByIdTaskByIdChecklistItemsCompleteResult struct {
	Items []models.ChecklistItemCollectionResponse
}

// ListMeTodoListByIdTaskByIdChecklistItems ...
func (c MeTodoListTaskChecklistItemClient) ListMeTodoListByIdTaskByIdChecklistItems(ctx context.Context, id MeTodoListTaskId) (result ListMeTodoListByIdTaskByIdChecklistItemsOperationResponse, err error) {
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

// ListMeTodoListByIdTaskByIdChecklistItemsComplete retrieves all the results into a single object
func (c MeTodoListTaskChecklistItemClient) ListMeTodoListByIdTaskByIdChecklistItemsComplete(ctx context.Context, id MeTodoListTaskId) (ListMeTodoListByIdTaskByIdChecklistItemsCompleteResult, error) {
	return c.ListMeTodoListByIdTaskByIdChecklistItemsCompleteMatchingPredicate(ctx, id, models.ChecklistItemCollectionResponseOperationPredicate{})
}

// ListMeTodoListByIdTaskByIdChecklistItemsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeTodoListTaskChecklistItemClient) ListMeTodoListByIdTaskByIdChecklistItemsCompleteMatchingPredicate(ctx context.Context, id MeTodoListTaskId, predicate models.ChecklistItemCollectionResponseOperationPredicate) (result ListMeTodoListByIdTaskByIdChecklistItemsCompleteResult, err error) {
	items := make([]models.ChecklistItemCollectionResponse, 0)

	resp, err := c.ListMeTodoListByIdTaskByIdChecklistItems(ctx, id)
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

	result = ListMeTodoListByIdTaskByIdChecklistItemsCompleteResult{
		Items: items,
	}
	return
}

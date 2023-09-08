package metodolist

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

type ListMeTodoListsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.TodoTaskListCollectionResponse
}

type ListMeTodoListsCompleteResult struct {
	Items []models.TodoTaskListCollectionResponse
}

// ListMeTodoLists ...
func (c MeTodoListClient) ListMeTodoLists(ctx context.Context) (result ListMeTodoListsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       "/me/todo/lists",
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

// ListMeTodoListsComplete retrieves all the results into a single object
func (c MeTodoListClient) ListMeTodoListsComplete(ctx context.Context) (ListMeTodoListsCompleteResult, error) {
	return c.ListMeTodoListsCompleteMatchingPredicate(ctx, models.TodoTaskListCollectionResponseOperationPredicate{})
}

// ListMeTodoListsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeTodoListClient) ListMeTodoListsCompleteMatchingPredicate(ctx context.Context, predicate models.TodoTaskListCollectionResponseOperationPredicate) (result ListMeTodoListsCompleteResult, err error) {
	items := make([]models.TodoTaskListCollectionResponse, 0)

	resp, err := c.ListMeTodoLists(ctx)
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

	result = ListMeTodoListsCompleteResult{
		Items: items,
	}
	return
}

package todolist

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListTodoListsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.TodoTaskList
}

type ListTodoListsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.TodoTaskList
}

type ListTodoListsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListTodoListsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListTodoLists ...
func (c TodoListClient) ListTodoLists(ctx context.Context, id UserId) (result ListTodoListsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListTodoListsCustomPager{},
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
		Values *[]beta.TodoTaskList `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListTodoListsComplete retrieves all the results into a single object
func (c TodoListClient) ListTodoListsComplete(ctx context.Context, id UserId) (ListTodoListsCompleteResult, error) {
	return c.ListTodoListsCompleteMatchingPredicate(ctx, id, TodoTaskListOperationPredicate{})
}

// ListTodoListsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c TodoListClient) ListTodoListsCompleteMatchingPredicate(ctx context.Context, id UserId, predicate TodoTaskListOperationPredicate) (result ListTodoListsCompleteResult, err error) {
	items := make([]beta.TodoTaskList, 0)

	resp, err := c.ListTodoLists(ctx, id)
	if err != nil {
		result.LatestHttpResponse = resp.HttpResponse
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

	result = ListTodoListsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}

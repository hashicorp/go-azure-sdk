package todolisttask

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

type ListTodoListTasksOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.TodoTask
}

type ListTodoListTasksCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.TodoTask
}

type ListTodoListTasksOperationOptions struct {
	Count     *bool
	Expand    *odata.Expand
	Filter    *string
	Metadata  *odata.Metadata
	OrderBy   *odata.OrderBy
	RetryFunc client.RequestRetryFunc
	Search    *string
	Select    *[]string
	Skip      *int64
	Top       *int64
}

func DefaultListTodoListTasksOperationOptions() ListTodoListTasksOperationOptions {
	return ListTodoListTasksOperationOptions{}
}

func (o ListTodoListTasksOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListTodoListTasksOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Count != nil {
		out.Count = *o.Count
	}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Filter != nil {
		out.Filter = *o.Filter
	}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.OrderBy != nil {
		out.OrderBy = *o.OrderBy
	}
	if o.Search != nil {
		out.Search = *o.Search
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	if o.Skip != nil {
		out.Skip = int(*o.Skip)
	}
	if o.Top != nil {
		out.Top = int(*o.Top)
	}
	return &out
}

func (o ListTodoListTasksOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListTodoListTasksCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListTodoListTasksCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListTodoListTasks - List tasks. Get the todoTask resources from the tasks navigation property of a specified
// todoTaskList.
func (c TodoListTaskClient) ListTodoListTasks(ctx context.Context, id beta.MeTodoListId, options ListTodoListTasksOperationOptions) (result ListTodoListTasksOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListTodoListTasksCustomPager{},
		Path:          fmt.Sprintf("%s/tasks", id.ID()),
		RetryFunc:     options.RetryFunc,
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
		Values *[]beta.TodoTask `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListTodoListTasksComplete retrieves all the results into a single object
func (c TodoListTaskClient) ListTodoListTasksComplete(ctx context.Context, id beta.MeTodoListId, options ListTodoListTasksOperationOptions) (ListTodoListTasksCompleteResult, error) {
	return c.ListTodoListTasksCompleteMatchingPredicate(ctx, id, options, TodoTaskOperationPredicate{})
}

// ListTodoListTasksCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c TodoListTaskClient) ListTodoListTasksCompleteMatchingPredicate(ctx context.Context, id beta.MeTodoListId, options ListTodoListTasksOperationOptions, predicate TodoTaskOperationPredicate) (result ListTodoListTasksCompleteResult, err error) {
	items := make([]beta.TodoTask, 0)

	resp, err := c.ListTodoListTasks(ctx, id, options)
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

	result = ListTodoListTasksCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}

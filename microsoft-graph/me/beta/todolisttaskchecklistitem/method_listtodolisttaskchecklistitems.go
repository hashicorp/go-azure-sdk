package todolisttaskchecklistitem

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

type ListTodoListTaskChecklistItemsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.ChecklistItem
}

type ListTodoListTaskChecklistItemsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.ChecklistItem
}

type ListTodoListTaskChecklistItemsOperationOptions struct {
	Count   *bool
	Expand  *odata.Expand
	Filter  *string
	OrderBy *odata.OrderBy
	Search  *string
	Select  *[]string
	Skip    *int64
	Top     *int64
}

func DefaultListTodoListTaskChecklistItemsOperationOptions() ListTodoListTaskChecklistItemsOperationOptions {
	return ListTodoListTaskChecklistItemsOperationOptions{}
}

func (o ListTodoListTaskChecklistItemsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListTodoListTaskChecklistItemsOperationOptions) ToOData() *odata.Query {
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

func (o ListTodoListTaskChecklistItemsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListTodoListTaskChecklistItemsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListTodoListTaskChecklistItemsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListTodoListTaskChecklistItems - List checklistItems. Get the checklistItem resources associated to the
// checklistItems navigation property of a todoTask.
func (c TodoListTaskChecklistItemClient) ListTodoListTaskChecklistItems(ctx context.Context, id beta.MeTodoListIdTaskId, options ListTodoListTaskChecklistItemsOperationOptions) (result ListTodoListTaskChecklistItemsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListTodoListTaskChecklistItemsCustomPager{},
		Path:          fmt.Sprintf("%s/checklistItems", id.ID()),
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
		Values *[]beta.ChecklistItem `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListTodoListTaskChecklistItemsComplete retrieves all the results into a single object
func (c TodoListTaskChecklistItemClient) ListTodoListTaskChecklistItemsComplete(ctx context.Context, id beta.MeTodoListIdTaskId, options ListTodoListTaskChecklistItemsOperationOptions) (ListTodoListTaskChecklistItemsCompleteResult, error) {
	return c.ListTodoListTaskChecklistItemsCompleteMatchingPredicate(ctx, id, options, ChecklistItemOperationPredicate{})
}

// ListTodoListTaskChecklistItemsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c TodoListTaskChecklistItemClient) ListTodoListTaskChecklistItemsCompleteMatchingPredicate(ctx context.Context, id beta.MeTodoListIdTaskId, options ListTodoListTaskChecklistItemsOperationOptions, predicate ChecklistItemOperationPredicate) (result ListTodoListTaskChecklistItemsCompleteResult, err error) {
	items := make([]beta.ChecklistItem, 0)

	resp, err := c.ListTodoListTaskChecklistItems(ctx, id, options)
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

	result = ListTodoListTaskChecklistItemsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}

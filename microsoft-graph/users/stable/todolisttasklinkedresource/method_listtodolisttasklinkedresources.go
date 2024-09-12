package todolisttasklinkedresource

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListTodoListTaskLinkedResourcesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.LinkedResource
}

type ListTodoListTaskLinkedResourcesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.LinkedResource
}

type ListTodoListTaskLinkedResourcesOperationOptions struct {
	Count   *bool
	Expand  *odata.Expand
	Filter  *string
	OrderBy *odata.OrderBy
	Search  *string
	Select  *[]string
	Skip    *int64
	Top     *int64
}

func DefaultListTodoListTaskLinkedResourcesOperationOptions() ListTodoListTaskLinkedResourcesOperationOptions {
	return ListTodoListTaskLinkedResourcesOperationOptions{}
}

func (o ListTodoListTaskLinkedResourcesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListTodoListTaskLinkedResourcesOperationOptions) ToOData() *odata.Query {
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

func (o ListTodoListTaskLinkedResourcesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListTodoListTaskLinkedResourcesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListTodoListTaskLinkedResourcesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListTodoListTaskLinkedResources - Get linkedResources from users. A collection of resources linked to the task.
func (c TodoListTaskLinkedResourceClient) ListTodoListTaskLinkedResources(ctx context.Context, id stable.UserIdTodoListIdTaskId, options ListTodoListTaskLinkedResourcesOperationOptions) (result ListTodoListTaskLinkedResourcesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListTodoListTaskLinkedResourcesCustomPager{},
		Path:          fmt.Sprintf("%s/linkedResources", id.ID()),
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
		Values *[]stable.LinkedResource `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListTodoListTaskLinkedResourcesComplete retrieves all the results into a single object
func (c TodoListTaskLinkedResourceClient) ListTodoListTaskLinkedResourcesComplete(ctx context.Context, id stable.UserIdTodoListIdTaskId, options ListTodoListTaskLinkedResourcesOperationOptions) (ListTodoListTaskLinkedResourcesCompleteResult, error) {
	return c.ListTodoListTaskLinkedResourcesCompleteMatchingPredicate(ctx, id, options, LinkedResourceOperationPredicate{})
}

// ListTodoListTaskLinkedResourcesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c TodoListTaskLinkedResourceClient) ListTodoListTaskLinkedResourcesCompleteMatchingPredicate(ctx context.Context, id stable.UserIdTodoListIdTaskId, options ListTodoListTaskLinkedResourcesOperationOptions, predicate LinkedResourceOperationPredicate) (result ListTodoListTaskLinkedResourcesCompleteResult, err error) {
	items := make([]stable.LinkedResource, 0)

	resp, err := c.ListTodoListTaskLinkedResources(ctx, id, options)
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

	result = ListTodoListTaskLinkedResourcesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}

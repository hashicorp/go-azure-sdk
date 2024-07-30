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

type ListTodoListTaskLinkedResourcesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListTodoListTaskLinkedResourcesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListTodoListTaskLinkedResources ...
func (c TodoListTaskLinkedResourceClient) ListTodoListTaskLinkedResources(ctx context.Context, id UserIdTodoListIdTaskId) (result ListTodoListTaskLinkedResourcesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListTodoListTaskLinkedResourcesCustomPager{},
		Path:       fmt.Sprintf("%s/linkedResources", id.ID()),
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
func (c TodoListTaskLinkedResourceClient) ListTodoListTaskLinkedResourcesComplete(ctx context.Context, id UserIdTodoListIdTaskId) (ListTodoListTaskLinkedResourcesCompleteResult, error) {
	return c.ListTodoListTaskLinkedResourcesCompleteMatchingPredicate(ctx, id, LinkedResourceOperationPredicate{})
}

// ListTodoListTaskLinkedResourcesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c TodoListTaskLinkedResourceClient) ListTodoListTaskLinkedResourcesCompleteMatchingPredicate(ctx context.Context, id UserIdTodoListIdTaskId, predicate LinkedResourceOperationPredicate) (result ListTodoListTaskLinkedResourcesCompleteResult, err error) {
	items := make([]stable.LinkedResource, 0)

	resp, err := c.ListTodoListTaskLinkedResources(ctx, id)
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

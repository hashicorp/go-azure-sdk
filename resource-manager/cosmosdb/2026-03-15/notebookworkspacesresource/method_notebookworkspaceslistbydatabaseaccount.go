package notebookworkspacesresource

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NotebookWorkspacesListByDatabaseAccountOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]NotebookWorkspace
}

type NotebookWorkspacesListByDatabaseAccountCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []NotebookWorkspace
}

type NotebookWorkspacesListByDatabaseAccountCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *NotebookWorkspacesListByDatabaseAccountCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// NotebookWorkspacesListByDatabaseAccount ...
func (c NotebookWorkspacesResourceClient) NotebookWorkspacesListByDatabaseAccount(ctx context.Context, id DatabaseAccountId) (result NotebookWorkspacesListByDatabaseAccountOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &NotebookWorkspacesListByDatabaseAccountCustomPager{},
		Path:       fmt.Sprintf("%s/notebookWorkspaces", id.ID()),
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
		Values *[]NotebookWorkspace `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// NotebookWorkspacesListByDatabaseAccountComplete retrieves all the results into a single object
func (c NotebookWorkspacesResourceClient) NotebookWorkspacesListByDatabaseAccountComplete(ctx context.Context, id DatabaseAccountId) (NotebookWorkspacesListByDatabaseAccountCompleteResult, error) {
	return c.NotebookWorkspacesListByDatabaseAccountCompleteMatchingPredicate(ctx, id, NotebookWorkspaceOperationPredicate{})
}

// NotebookWorkspacesListByDatabaseAccountCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c NotebookWorkspacesResourceClient) NotebookWorkspacesListByDatabaseAccountCompleteMatchingPredicate(ctx context.Context, id DatabaseAccountId, predicate NotebookWorkspaceOperationPredicate) (result NotebookWorkspacesListByDatabaseAccountCompleteResult, err error) {
	items := make([]NotebookWorkspace, 0)

	resp, err := c.NotebookWorkspacesListByDatabaseAccount(ctx, id)
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

	result = NotebookWorkspacesListByDatabaseAccountCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}

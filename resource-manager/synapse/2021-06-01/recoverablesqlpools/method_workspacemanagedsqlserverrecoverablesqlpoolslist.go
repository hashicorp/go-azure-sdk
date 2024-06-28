package recoverablesqlpools

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkspaceManagedSqlServerRecoverableSqlPoolsListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]RecoverableSqlPool
}

type WorkspaceManagedSqlServerRecoverableSqlPoolsListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []RecoverableSqlPool
}

type WorkspaceManagedSqlServerRecoverableSqlPoolsListCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *WorkspaceManagedSqlServerRecoverableSqlPoolsListCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// WorkspaceManagedSqlServerRecoverableSqlPoolsList ...
func (c RecoverableSqlPoolsClient) WorkspaceManagedSqlServerRecoverableSqlPoolsList(ctx context.Context, id WorkspaceId) (result WorkspaceManagedSqlServerRecoverableSqlPoolsListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &WorkspaceManagedSqlServerRecoverableSqlPoolsListCustomPager{},
		Path:       fmt.Sprintf("%s/recoverableSqlPools", id.ID()),
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
		Values *[]RecoverableSqlPool `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// WorkspaceManagedSqlServerRecoverableSqlPoolsListComplete retrieves all the results into a single object
func (c RecoverableSqlPoolsClient) WorkspaceManagedSqlServerRecoverableSqlPoolsListComplete(ctx context.Context, id WorkspaceId) (WorkspaceManagedSqlServerRecoverableSqlPoolsListCompleteResult, error) {
	return c.WorkspaceManagedSqlServerRecoverableSqlPoolsListCompleteMatchingPredicate(ctx, id, RecoverableSqlPoolOperationPredicate{})
}

// WorkspaceManagedSqlServerRecoverableSqlPoolsListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c RecoverableSqlPoolsClient) WorkspaceManagedSqlServerRecoverableSqlPoolsListCompleteMatchingPredicate(ctx context.Context, id WorkspaceId, predicate RecoverableSqlPoolOperationPredicate) (result WorkspaceManagedSqlServerRecoverableSqlPoolsListCompleteResult, err error) {
	items := make([]RecoverableSqlPool, 0)

	resp, err := c.WorkspaceManagedSqlServerRecoverableSqlPoolsList(ctx, id)
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

	result = WorkspaceManagedSqlServerRecoverableSqlPoolsListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}

package workspacemanagedsqlserver

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RecoverableSqlPoolsListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]RecoverableSqlPool
}

type RecoverableSqlPoolsListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []RecoverableSqlPool
}

type RecoverableSqlPoolsListCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *RecoverableSqlPoolsListCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// RecoverableSqlPoolsList ...
func (c WorkspaceManagedSqlServerClient) RecoverableSqlPoolsList(ctx context.Context, id WorkspaceId) (result RecoverableSqlPoolsListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &RecoverableSqlPoolsListCustomPager{},
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

// RecoverableSqlPoolsListComplete retrieves all the results into a single object
func (c WorkspaceManagedSqlServerClient) RecoverableSqlPoolsListComplete(ctx context.Context, id WorkspaceId) (RecoverableSqlPoolsListCompleteResult, error) {
	return c.RecoverableSqlPoolsListCompleteMatchingPredicate(ctx, id, RecoverableSqlPoolOperationPredicate{})
}

// RecoverableSqlPoolsListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c WorkspaceManagedSqlServerClient) RecoverableSqlPoolsListCompleteMatchingPredicate(ctx context.Context, id WorkspaceId, predicate RecoverableSqlPoolOperationPredicate) (result RecoverableSqlPoolsListCompleteResult, err error) {
	items := make([]RecoverableSqlPool, 0)

	resp, err := c.RecoverableSqlPoolsList(ctx, id)
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

	result = RecoverableSqlPoolsListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}

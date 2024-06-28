package workspacemanagedsqlserversqlusages

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkspaceManagedSqlServerUsagesListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]ServerUsage
}

type WorkspaceManagedSqlServerUsagesListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []ServerUsage
}

type WorkspaceManagedSqlServerUsagesListCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *WorkspaceManagedSqlServerUsagesListCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// WorkspaceManagedSqlServerUsagesList ...
func (c WorkspaceManagedSqlServerSqlUsagesClient) WorkspaceManagedSqlServerUsagesList(ctx context.Context, id WorkspaceId) (result WorkspaceManagedSqlServerUsagesListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &WorkspaceManagedSqlServerUsagesListCustomPager{},
		Path:       fmt.Sprintf("%s/sqlUsages", id.ID()),
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
		Values *[]ServerUsage `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// WorkspaceManagedSqlServerUsagesListComplete retrieves all the results into a single object
func (c WorkspaceManagedSqlServerSqlUsagesClient) WorkspaceManagedSqlServerUsagesListComplete(ctx context.Context, id WorkspaceId) (WorkspaceManagedSqlServerUsagesListCompleteResult, error) {
	return c.WorkspaceManagedSqlServerUsagesListCompleteMatchingPredicate(ctx, id, ServerUsageOperationPredicate{})
}

// WorkspaceManagedSqlServerUsagesListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c WorkspaceManagedSqlServerSqlUsagesClient) WorkspaceManagedSqlServerUsagesListCompleteMatchingPredicate(ctx context.Context, id WorkspaceId, predicate ServerUsageOperationPredicate) (result WorkspaceManagedSqlServerUsagesListCompleteResult, err error) {
	items := make([]ServerUsage, 0)

	resp, err := c.WorkspaceManagedSqlServerUsagesList(ctx, id)
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

	result = WorkspaceManagedSqlServerUsagesListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}

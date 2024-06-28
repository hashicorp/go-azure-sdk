package workspacemanagedsqlserverblobauditing

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkspaceManagedSqlServerBlobAuditingPoliciesListByWorkspaceOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]ServerBlobAuditingPolicy
}

type WorkspaceManagedSqlServerBlobAuditingPoliciesListByWorkspaceCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []ServerBlobAuditingPolicy
}

type WorkspaceManagedSqlServerBlobAuditingPoliciesListByWorkspaceCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *WorkspaceManagedSqlServerBlobAuditingPoliciesListByWorkspaceCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// WorkspaceManagedSqlServerBlobAuditingPoliciesListByWorkspace ...
func (c WorkspaceManagedSqlServerBlobAuditingClient) WorkspaceManagedSqlServerBlobAuditingPoliciesListByWorkspace(ctx context.Context, id WorkspaceId) (result WorkspaceManagedSqlServerBlobAuditingPoliciesListByWorkspaceOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &WorkspaceManagedSqlServerBlobAuditingPoliciesListByWorkspaceCustomPager{},
		Path:       fmt.Sprintf("%s/auditingSettings", id.ID()),
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
		Values *[]ServerBlobAuditingPolicy `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// WorkspaceManagedSqlServerBlobAuditingPoliciesListByWorkspaceComplete retrieves all the results into a single object
func (c WorkspaceManagedSqlServerBlobAuditingClient) WorkspaceManagedSqlServerBlobAuditingPoliciesListByWorkspaceComplete(ctx context.Context, id WorkspaceId) (WorkspaceManagedSqlServerBlobAuditingPoliciesListByWorkspaceCompleteResult, error) {
	return c.WorkspaceManagedSqlServerBlobAuditingPoliciesListByWorkspaceCompleteMatchingPredicate(ctx, id, ServerBlobAuditingPolicyOperationPredicate{})
}

// WorkspaceManagedSqlServerBlobAuditingPoliciesListByWorkspaceCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c WorkspaceManagedSqlServerBlobAuditingClient) WorkspaceManagedSqlServerBlobAuditingPoliciesListByWorkspaceCompleteMatchingPredicate(ctx context.Context, id WorkspaceId, predicate ServerBlobAuditingPolicyOperationPredicate) (result WorkspaceManagedSqlServerBlobAuditingPoliciesListByWorkspaceCompleteResult, err error) {
	items := make([]ServerBlobAuditingPolicy, 0)

	resp, err := c.WorkspaceManagedSqlServerBlobAuditingPoliciesListByWorkspace(ctx, id)
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

	result = WorkspaceManagedSqlServerBlobAuditingPoliciesListByWorkspaceCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}

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

type WorkspaceManagedSqlServerExtendedBlobAuditingPoliciesListByWorkspaceOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]ExtendedServerBlobAuditingPolicy
}

type WorkspaceManagedSqlServerExtendedBlobAuditingPoliciesListByWorkspaceCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []ExtendedServerBlobAuditingPolicy
}

// WorkspaceManagedSqlServerExtendedBlobAuditingPoliciesListByWorkspace ...
func (c WorkspaceManagedSqlServerBlobAuditingClient) WorkspaceManagedSqlServerExtendedBlobAuditingPoliciesListByWorkspace(ctx context.Context, id WorkspaceId) (result WorkspaceManagedSqlServerExtendedBlobAuditingPoliciesListByWorkspaceOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/extendedAuditingSettings", id.ID()),
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
		Values *[]ExtendedServerBlobAuditingPolicy `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// WorkspaceManagedSqlServerExtendedBlobAuditingPoliciesListByWorkspaceComplete retrieves all the results into a single object
func (c WorkspaceManagedSqlServerBlobAuditingClient) WorkspaceManagedSqlServerExtendedBlobAuditingPoliciesListByWorkspaceComplete(ctx context.Context, id WorkspaceId) (WorkspaceManagedSqlServerExtendedBlobAuditingPoliciesListByWorkspaceCompleteResult, error) {
	return c.WorkspaceManagedSqlServerExtendedBlobAuditingPoliciesListByWorkspaceCompleteMatchingPredicate(ctx, id, ExtendedServerBlobAuditingPolicyOperationPredicate{})
}

// WorkspaceManagedSqlServerExtendedBlobAuditingPoliciesListByWorkspaceCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c WorkspaceManagedSqlServerBlobAuditingClient) WorkspaceManagedSqlServerExtendedBlobAuditingPoliciesListByWorkspaceCompleteMatchingPredicate(ctx context.Context, id WorkspaceId, predicate ExtendedServerBlobAuditingPolicyOperationPredicate) (result WorkspaceManagedSqlServerExtendedBlobAuditingPoliciesListByWorkspaceCompleteResult, err error) {
	items := make([]ExtendedServerBlobAuditingPolicy, 0)

	resp, err := c.WorkspaceManagedSqlServerExtendedBlobAuditingPoliciesListByWorkspace(ctx, id)
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

	result = WorkspaceManagedSqlServerExtendedBlobAuditingPoliciesListByWorkspaceCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}

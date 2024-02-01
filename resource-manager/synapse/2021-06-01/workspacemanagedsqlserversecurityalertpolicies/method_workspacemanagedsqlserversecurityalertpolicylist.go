package workspacemanagedsqlserversecurityalertpolicies

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkspaceManagedSqlServerSecurityAlertPolicyListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]ServerSecurityAlertPolicy
}

type WorkspaceManagedSqlServerSecurityAlertPolicyListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []ServerSecurityAlertPolicy
}

// WorkspaceManagedSqlServerSecurityAlertPolicyList ...
func (c WorkspaceManagedSqlServerSecurityAlertPoliciesClient) WorkspaceManagedSqlServerSecurityAlertPolicyList(ctx context.Context, id WorkspaceId) (result WorkspaceManagedSqlServerSecurityAlertPolicyListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/securityAlertPolicies", id.ID()),
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
		Values *[]ServerSecurityAlertPolicy `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// WorkspaceManagedSqlServerSecurityAlertPolicyListComplete retrieves all the results into a single object
func (c WorkspaceManagedSqlServerSecurityAlertPoliciesClient) WorkspaceManagedSqlServerSecurityAlertPolicyListComplete(ctx context.Context, id WorkspaceId) (WorkspaceManagedSqlServerSecurityAlertPolicyListCompleteResult, error) {
	return c.WorkspaceManagedSqlServerSecurityAlertPolicyListCompleteMatchingPredicate(ctx, id, ServerSecurityAlertPolicyOperationPredicate{})
}

// WorkspaceManagedSqlServerSecurityAlertPolicyListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c WorkspaceManagedSqlServerSecurityAlertPoliciesClient) WorkspaceManagedSqlServerSecurityAlertPolicyListCompleteMatchingPredicate(ctx context.Context, id WorkspaceId, predicate ServerSecurityAlertPolicyOperationPredicate) (result WorkspaceManagedSqlServerSecurityAlertPolicyListCompleteResult, err error) {
	items := make([]ServerSecurityAlertPolicy, 0)

	resp, err := c.WorkspaceManagedSqlServerSecurityAlertPolicyList(ctx, id)
	if err != nil {
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

	result = WorkspaceManagedSqlServerSecurityAlertPolicyListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}

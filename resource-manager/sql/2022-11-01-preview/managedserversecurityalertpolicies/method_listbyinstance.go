package managedserversecurityalertpolicies

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByInstanceOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]ManagedServerSecurityAlertPolicy
}

type ListByInstanceCompleteResult struct {
	Items []ManagedServerSecurityAlertPolicy
}

// ListByInstance ...
func (c ManagedServerSecurityAlertPoliciesClient) ListByInstance(ctx context.Context, id ManagedInstanceId) (result ListByInstanceOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json",
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
		Values *[]ManagedServerSecurityAlertPolicy `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListByInstanceComplete retrieves all the results into a single object
func (c ManagedServerSecurityAlertPoliciesClient) ListByInstanceComplete(ctx context.Context, id ManagedInstanceId) (ListByInstanceCompleteResult, error) {
	return c.ListByInstanceCompleteMatchingPredicate(ctx, id, ManagedServerSecurityAlertPolicyOperationPredicate{})
}

// ListByInstanceCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ManagedServerSecurityAlertPoliciesClient) ListByInstanceCompleteMatchingPredicate(ctx context.Context, id ManagedInstanceId, predicate ManagedServerSecurityAlertPolicyOperationPredicate) (result ListByInstanceCompleteResult, err error) {
	items := make([]ManagedServerSecurityAlertPolicy, 0)

	resp, err := c.ListByInstance(ctx, id)
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

	result = ListByInstanceCompleteResult{
		Items: items,
	}
	return
}

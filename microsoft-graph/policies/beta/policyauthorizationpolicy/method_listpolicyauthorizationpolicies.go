package policyauthorizationpolicy

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common/beta/models"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListPolicyAuthorizationPoliciesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.AuthorizationPolicyCollectionResponse
}

type ListPolicyAuthorizationPoliciesCompleteResult struct {
	Items []models.AuthorizationPolicyCollectionResponse
}

// ListPolicyAuthorizationPolicies ...
func (c PolicyAuthorizationPolicyClient) ListPolicyAuthorizationPolicies(ctx context.Context) (result ListPolicyAuthorizationPoliciesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       "/policies/authorizationPolicy",
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
		Values *[]models.AuthorizationPolicyCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListPolicyAuthorizationPoliciesComplete retrieves all the results into a single object
func (c PolicyAuthorizationPolicyClient) ListPolicyAuthorizationPoliciesComplete(ctx context.Context) (ListPolicyAuthorizationPoliciesCompleteResult, error) {
	return c.ListPolicyAuthorizationPoliciesCompleteMatchingPredicate(ctx, models.AuthorizationPolicyCollectionResponseOperationPredicate{})
}

// ListPolicyAuthorizationPoliciesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c PolicyAuthorizationPolicyClient) ListPolicyAuthorizationPoliciesCompleteMatchingPredicate(ctx context.Context, predicate models.AuthorizationPolicyCollectionResponseOperationPredicate) (result ListPolicyAuthorizationPoliciesCompleteResult, err error) {
	items := make([]models.AuthorizationPolicyCollectionResponse, 0)

	resp, err := c.ListPolicyAuthorizationPolicies(ctx)
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

	result = ListPolicyAuthorizationPoliciesCompleteResult{
		Items: items,
	}
	return
}

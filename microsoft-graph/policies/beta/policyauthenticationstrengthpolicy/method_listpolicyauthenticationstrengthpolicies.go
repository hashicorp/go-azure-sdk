package policyauthenticationstrengthpolicy

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

type ListPolicyAuthenticationStrengthPoliciesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.AuthenticationStrengthPolicyCollectionResponse
}

type ListPolicyAuthenticationStrengthPoliciesCompleteResult struct {
	Items []models.AuthenticationStrengthPolicyCollectionResponse
}

// ListPolicyAuthenticationStrengthPolicies ...
func (c PolicyAuthenticationStrengthPolicyClient) ListPolicyAuthenticationStrengthPolicies(ctx context.Context) (result ListPolicyAuthenticationStrengthPoliciesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       "/policies/authenticationStrengthPolicies",
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
		Values *[]models.AuthenticationStrengthPolicyCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListPolicyAuthenticationStrengthPoliciesComplete retrieves all the results into a single object
func (c PolicyAuthenticationStrengthPolicyClient) ListPolicyAuthenticationStrengthPoliciesComplete(ctx context.Context) (ListPolicyAuthenticationStrengthPoliciesCompleteResult, error) {
	return c.ListPolicyAuthenticationStrengthPoliciesCompleteMatchingPredicate(ctx, models.AuthenticationStrengthPolicyCollectionResponseOperationPredicate{})
}

// ListPolicyAuthenticationStrengthPoliciesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c PolicyAuthenticationStrengthPolicyClient) ListPolicyAuthenticationStrengthPoliciesCompleteMatchingPredicate(ctx context.Context, predicate models.AuthenticationStrengthPolicyCollectionResponseOperationPredicate) (result ListPolicyAuthenticationStrengthPoliciesCompleteResult, err error) {
	items := make([]models.AuthenticationStrengthPolicyCollectionResponse, 0)

	resp, err := c.ListPolicyAuthenticationStrengthPolicies(ctx)
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

	result = ListPolicyAuthenticationStrengthPoliciesCompleteResult{
		Items: items,
	}
	return
}

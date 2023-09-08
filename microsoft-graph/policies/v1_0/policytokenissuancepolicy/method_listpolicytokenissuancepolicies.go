package policytokenissuancepolicy

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common/v1_0/models"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListPolicyTokenIssuancePoliciesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.TokenIssuancePolicyCollectionResponse
}

type ListPolicyTokenIssuancePoliciesCompleteResult struct {
	Items []models.TokenIssuancePolicyCollectionResponse
}

// ListPolicyTokenIssuancePolicies ...
func (c PolicyTokenIssuancePolicyClient) ListPolicyTokenIssuancePolicies(ctx context.Context) (result ListPolicyTokenIssuancePoliciesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       "/policies/tokenIssuancePolicies",
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
		Values *[]models.TokenIssuancePolicyCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListPolicyTokenIssuancePoliciesComplete retrieves all the results into a single object
func (c PolicyTokenIssuancePolicyClient) ListPolicyTokenIssuancePoliciesComplete(ctx context.Context) (ListPolicyTokenIssuancePoliciesCompleteResult, error) {
	return c.ListPolicyTokenIssuancePoliciesCompleteMatchingPredicate(ctx, models.TokenIssuancePolicyCollectionResponseOperationPredicate{})
}

// ListPolicyTokenIssuancePoliciesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c PolicyTokenIssuancePolicyClient) ListPolicyTokenIssuancePoliciesCompleteMatchingPredicate(ctx context.Context, predicate models.TokenIssuancePolicyCollectionResponseOperationPredicate) (result ListPolicyTokenIssuancePoliciesCompleteResult, err error) {
	items := make([]models.TokenIssuancePolicyCollectionResponse, 0)

	resp, err := c.ListPolicyTokenIssuancePolicies(ctx)
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

	result = ListPolicyTokenIssuancePoliciesCompleteResult{
		Items: items,
	}
	return
}

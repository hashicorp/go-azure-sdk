package policyclaimsmappingpolicy

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

type ListPolicyClaimsMappingPoliciesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.ClaimsMappingPolicyCollectionResponse
}

type ListPolicyClaimsMappingPoliciesCompleteResult struct {
	Items []models.ClaimsMappingPolicyCollectionResponse
}

// ListPolicyClaimsMappingPolicies ...
func (c PolicyClaimsMappingPolicyClient) ListPolicyClaimsMappingPolicies(ctx context.Context) (result ListPolicyClaimsMappingPoliciesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       "/policies/claimsMappingPolicies",
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
		Values *[]models.ClaimsMappingPolicyCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListPolicyClaimsMappingPoliciesComplete retrieves all the results into a single object
func (c PolicyClaimsMappingPolicyClient) ListPolicyClaimsMappingPoliciesComplete(ctx context.Context) (ListPolicyClaimsMappingPoliciesCompleteResult, error) {
	return c.ListPolicyClaimsMappingPoliciesCompleteMatchingPredicate(ctx, models.ClaimsMappingPolicyCollectionResponseOperationPredicate{})
}

// ListPolicyClaimsMappingPoliciesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c PolicyClaimsMappingPolicyClient) ListPolicyClaimsMappingPoliciesCompleteMatchingPredicate(ctx context.Context, predicate models.ClaimsMappingPolicyCollectionResponseOperationPredicate) (result ListPolicyClaimsMappingPoliciesCompleteResult, err error) {
	items := make([]models.ClaimsMappingPolicyCollectionResponse, 0)

	resp, err := c.ListPolicyClaimsMappingPolicies(ctx)
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

	result = ListPolicyClaimsMappingPoliciesCompleteResult{
		Items: items,
	}
	return
}

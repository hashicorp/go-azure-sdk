package claimsmappingpolicy

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListClaimsMappingPoliciesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.ClaimsMappingPolicy
}

type ListClaimsMappingPoliciesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.ClaimsMappingPolicy
}

type ListClaimsMappingPoliciesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListClaimsMappingPoliciesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListClaimsMappingPolicies ...
func (c ClaimsMappingPolicyClient) ListClaimsMappingPolicies(ctx context.Context) (result ListClaimsMappingPoliciesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListClaimsMappingPoliciesCustomPager{},
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
		Values *[]stable.ClaimsMappingPolicy `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListClaimsMappingPoliciesComplete retrieves all the results into a single object
func (c ClaimsMappingPolicyClient) ListClaimsMappingPoliciesComplete(ctx context.Context) (ListClaimsMappingPoliciesCompleteResult, error) {
	return c.ListClaimsMappingPoliciesCompleteMatchingPredicate(ctx, ClaimsMappingPolicyOperationPredicate{})
}

// ListClaimsMappingPoliciesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ClaimsMappingPolicyClient) ListClaimsMappingPoliciesCompleteMatchingPredicate(ctx context.Context, predicate ClaimsMappingPolicyOperationPredicate) (result ListClaimsMappingPoliciesCompleteResult, err error) {
	items := make([]stable.ClaimsMappingPolicy, 0)

	resp, err := c.ListClaimsMappingPolicies(ctx)
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

	result = ListClaimsMappingPoliciesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}

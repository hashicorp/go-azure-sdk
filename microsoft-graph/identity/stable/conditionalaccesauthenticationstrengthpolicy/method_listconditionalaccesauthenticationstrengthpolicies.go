package conditionalaccesauthenticationstrengthpolicy

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

type ListConditionalAccesAuthenticationStrengthPoliciesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.AuthenticationStrengthPolicy
}

type ListConditionalAccesAuthenticationStrengthPoliciesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.AuthenticationStrengthPolicy
}

type ListConditionalAccesAuthenticationStrengthPoliciesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListConditionalAccesAuthenticationStrengthPoliciesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListConditionalAccesAuthenticationStrengthPolicies ...
func (c ConditionalAccesAuthenticationStrengthPolicyClient) ListConditionalAccesAuthenticationStrengthPolicies(ctx context.Context) (result ListConditionalAccesAuthenticationStrengthPoliciesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListConditionalAccesAuthenticationStrengthPoliciesCustomPager{},
		Path:       "/identity/conditionalAccess/authenticationStrength/policies",
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
		Values *[]stable.AuthenticationStrengthPolicy `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListConditionalAccesAuthenticationStrengthPoliciesComplete retrieves all the results into a single object
func (c ConditionalAccesAuthenticationStrengthPolicyClient) ListConditionalAccesAuthenticationStrengthPoliciesComplete(ctx context.Context) (ListConditionalAccesAuthenticationStrengthPoliciesCompleteResult, error) {
	return c.ListConditionalAccesAuthenticationStrengthPoliciesCompleteMatchingPredicate(ctx, AuthenticationStrengthPolicyOperationPredicate{})
}

// ListConditionalAccesAuthenticationStrengthPoliciesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ConditionalAccesAuthenticationStrengthPolicyClient) ListConditionalAccesAuthenticationStrengthPoliciesCompleteMatchingPredicate(ctx context.Context, predicate AuthenticationStrengthPolicyOperationPredicate) (result ListConditionalAccesAuthenticationStrengthPoliciesCompleteResult, err error) {
	items := make([]stable.AuthenticationStrengthPolicy, 0)

	resp, err := c.ListConditionalAccesAuthenticationStrengthPolicies(ctx)
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

	result = ListConditionalAccesAuthenticationStrengthPoliciesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}

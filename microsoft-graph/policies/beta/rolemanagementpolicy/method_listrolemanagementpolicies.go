package rolemanagementpolicy

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListRoleManagementPoliciesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.UnifiedRoleManagementPolicy
}

type ListRoleManagementPoliciesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.UnifiedRoleManagementPolicy
}

type ListRoleManagementPoliciesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListRoleManagementPoliciesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListRoleManagementPolicies ...
func (c RoleManagementPolicyClient) ListRoleManagementPolicies(ctx context.Context) (result ListRoleManagementPoliciesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListRoleManagementPoliciesCustomPager{},
		Path:       "/policies/roleManagementPolicies",
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
		Values *[]beta.UnifiedRoleManagementPolicy `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListRoleManagementPoliciesComplete retrieves all the results into a single object
func (c RoleManagementPolicyClient) ListRoleManagementPoliciesComplete(ctx context.Context) (ListRoleManagementPoliciesCompleteResult, error) {
	return c.ListRoleManagementPoliciesCompleteMatchingPredicate(ctx, UnifiedRoleManagementPolicyOperationPredicate{})
}

// ListRoleManagementPoliciesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c RoleManagementPolicyClient) ListRoleManagementPoliciesCompleteMatchingPredicate(ctx context.Context, predicate UnifiedRoleManagementPolicyOperationPredicate) (result ListRoleManagementPoliciesCompleteResult, err error) {
	items := make([]beta.UnifiedRoleManagementPolicy, 0)

	resp, err := c.ListRoleManagementPolicies(ctx)
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

	result = ListRoleManagementPoliciesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}

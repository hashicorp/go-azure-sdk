package entitlementmanagementassignmentpolicy

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

type ListEntitlementManagementAssignmentPoliciesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.AccessPackageAssignmentPolicy
}

type ListEntitlementManagementAssignmentPoliciesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.AccessPackageAssignmentPolicy
}

type ListEntitlementManagementAssignmentPoliciesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListEntitlementManagementAssignmentPoliciesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListEntitlementManagementAssignmentPolicies ...
func (c EntitlementManagementAssignmentPolicyClient) ListEntitlementManagementAssignmentPolicies(ctx context.Context) (result ListEntitlementManagementAssignmentPoliciesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListEntitlementManagementAssignmentPoliciesCustomPager{},
		Path:       "/identityGovernance/entitlementManagement/assignmentPolicies",
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
		Values *[]stable.AccessPackageAssignmentPolicy `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListEntitlementManagementAssignmentPoliciesComplete retrieves all the results into a single object
func (c EntitlementManagementAssignmentPolicyClient) ListEntitlementManagementAssignmentPoliciesComplete(ctx context.Context) (ListEntitlementManagementAssignmentPoliciesCompleteResult, error) {
	return c.ListEntitlementManagementAssignmentPoliciesCompleteMatchingPredicate(ctx, AccessPackageAssignmentPolicyOperationPredicate{})
}

// ListEntitlementManagementAssignmentPoliciesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c EntitlementManagementAssignmentPolicyClient) ListEntitlementManagementAssignmentPoliciesCompleteMatchingPredicate(ctx context.Context, predicate AccessPackageAssignmentPolicyOperationPredicate) (result ListEntitlementManagementAssignmentPoliciesCompleteResult, err error) {
	items := make([]stable.AccessPackageAssignmentPolicy, 0)

	resp, err := c.ListEntitlementManagementAssignmentPolicies(ctx)
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

	result = ListEntitlementManagementAssignmentPoliciesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}

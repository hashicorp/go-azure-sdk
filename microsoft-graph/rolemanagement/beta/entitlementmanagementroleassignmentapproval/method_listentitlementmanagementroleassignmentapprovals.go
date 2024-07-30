package entitlementmanagementroleassignmentapproval

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

type ListEntitlementManagementRoleAssignmentApprovalsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.Approval
}

type ListEntitlementManagementRoleAssignmentApprovalsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.Approval
}

type ListEntitlementManagementRoleAssignmentApprovalsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListEntitlementManagementRoleAssignmentApprovalsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListEntitlementManagementRoleAssignmentApprovals ...
func (c EntitlementManagementRoleAssignmentApprovalClient) ListEntitlementManagementRoleAssignmentApprovals(ctx context.Context) (result ListEntitlementManagementRoleAssignmentApprovalsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListEntitlementManagementRoleAssignmentApprovalsCustomPager{},
		Path:       "/roleManagement/entitlementManagement/roleAssignmentApprovals",
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
		Values *[]beta.Approval `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListEntitlementManagementRoleAssignmentApprovalsComplete retrieves all the results into a single object
func (c EntitlementManagementRoleAssignmentApprovalClient) ListEntitlementManagementRoleAssignmentApprovalsComplete(ctx context.Context) (ListEntitlementManagementRoleAssignmentApprovalsCompleteResult, error) {
	return c.ListEntitlementManagementRoleAssignmentApprovalsCompleteMatchingPredicate(ctx, ApprovalOperationPredicate{})
}

// ListEntitlementManagementRoleAssignmentApprovalsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c EntitlementManagementRoleAssignmentApprovalClient) ListEntitlementManagementRoleAssignmentApprovalsCompleteMatchingPredicate(ctx context.Context, predicate ApprovalOperationPredicate) (result ListEntitlementManagementRoleAssignmentApprovalsCompleteResult, err error) {
	items := make([]beta.Approval, 0)

	resp, err := c.ListEntitlementManagementRoleAssignmentApprovals(ctx)
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

	result = ListEntitlementManagementRoleAssignmentApprovalsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}

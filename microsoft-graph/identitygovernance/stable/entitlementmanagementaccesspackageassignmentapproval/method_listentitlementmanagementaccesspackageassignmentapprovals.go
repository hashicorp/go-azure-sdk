package entitlementmanagementaccesspackageassignmentapproval

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

type ListEntitlementManagementAccessPackageAssignmentApprovalsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.Approval
}

type ListEntitlementManagementAccessPackageAssignmentApprovalsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.Approval
}

type ListEntitlementManagementAccessPackageAssignmentApprovalsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListEntitlementManagementAccessPackageAssignmentApprovalsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListEntitlementManagementAccessPackageAssignmentApprovals ...
func (c EntitlementManagementAccessPackageAssignmentApprovalClient) ListEntitlementManagementAccessPackageAssignmentApprovals(ctx context.Context) (result ListEntitlementManagementAccessPackageAssignmentApprovalsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListEntitlementManagementAccessPackageAssignmentApprovalsCustomPager{},
		Path:       "/identityGovernance/entitlementManagement/accessPackageAssignmentApprovals",
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
		Values *[]stable.Approval `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListEntitlementManagementAccessPackageAssignmentApprovalsComplete retrieves all the results into a single object
func (c EntitlementManagementAccessPackageAssignmentApprovalClient) ListEntitlementManagementAccessPackageAssignmentApprovalsComplete(ctx context.Context) (ListEntitlementManagementAccessPackageAssignmentApprovalsCompleteResult, error) {
	return c.ListEntitlementManagementAccessPackageAssignmentApprovalsCompleteMatchingPredicate(ctx, ApprovalOperationPredicate{})
}

// ListEntitlementManagementAccessPackageAssignmentApprovalsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c EntitlementManagementAccessPackageAssignmentApprovalClient) ListEntitlementManagementAccessPackageAssignmentApprovalsCompleteMatchingPredicate(ctx context.Context, predicate ApprovalOperationPredicate) (result ListEntitlementManagementAccessPackageAssignmentApprovalsCompleteResult, err error) {
	items := make([]stable.Approval, 0)

	resp, err := c.ListEntitlementManagementAccessPackageAssignmentApprovals(ctx)
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

	result = ListEntitlementManagementAccessPackageAssignmentApprovalsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}

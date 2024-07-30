package entitlementmanagementroleassignmentapprovalstep

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

type ListEntitlementManagementRoleAssignmentApprovalStepsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.ApprovalStep
}

type ListEntitlementManagementRoleAssignmentApprovalStepsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.ApprovalStep
}

type ListEntitlementManagementRoleAssignmentApprovalStepsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListEntitlementManagementRoleAssignmentApprovalStepsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListEntitlementManagementRoleAssignmentApprovalSteps ...
func (c EntitlementManagementRoleAssignmentApprovalStepClient) ListEntitlementManagementRoleAssignmentApprovalSteps(ctx context.Context, id RoleManagementEntitlementManagementRoleAssignmentApprovalId) (result ListEntitlementManagementRoleAssignmentApprovalStepsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListEntitlementManagementRoleAssignmentApprovalStepsCustomPager{},
		Path:       fmt.Sprintf("%s/steps", id.ID()),
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
		Values *[]beta.ApprovalStep `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListEntitlementManagementRoleAssignmentApprovalStepsComplete retrieves all the results into a single object
func (c EntitlementManagementRoleAssignmentApprovalStepClient) ListEntitlementManagementRoleAssignmentApprovalStepsComplete(ctx context.Context, id RoleManagementEntitlementManagementRoleAssignmentApprovalId) (ListEntitlementManagementRoleAssignmentApprovalStepsCompleteResult, error) {
	return c.ListEntitlementManagementRoleAssignmentApprovalStepsCompleteMatchingPredicate(ctx, id, ApprovalStepOperationPredicate{})
}

// ListEntitlementManagementRoleAssignmentApprovalStepsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c EntitlementManagementRoleAssignmentApprovalStepClient) ListEntitlementManagementRoleAssignmentApprovalStepsCompleteMatchingPredicate(ctx context.Context, id RoleManagementEntitlementManagementRoleAssignmentApprovalId, predicate ApprovalStepOperationPredicate) (result ListEntitlementManagementRoleAssignmentApprovalStepsCompleteResult, err error) {
	items := make([]beta.ApprovalStep, 0)

	resp, err := c.ListEntitlementManagementRoleAssignmentApprovalSteps(ctx, id)
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

	result = ListEntitlementManagementRoleAssignmentApprovalStepsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}

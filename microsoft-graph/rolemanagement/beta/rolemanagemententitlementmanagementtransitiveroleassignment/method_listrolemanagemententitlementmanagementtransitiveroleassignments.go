package rolemanagemententitlementmanagementtransitiveroleassignment

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

type ListRoleManagementEntitlementManagementTransitiveRoleAssignmentsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.UnifiedRoleAssignmentCollectionResponse
}

type ListRoleManagementEntitlementManagementTransitiveRoleAssignmentsCompleteResult struct {
	Items []models.UnifiedRoleAssignmentCollectionResponse
}

// ListRoleManagementEntitlementManagementTransitiveRoleAssignments ...
func (c RoleManagementEntitlementManagementTransitiveRoleAssignmentClient) ListRoleManagementEntitlementManagementTransitiveRoleAssignments(ctx context.Context) (result ListRoleManagementEntitlementManagementTransitiveRoleAssignmentsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       "/roleManagement/entitlementManagement/transitiveRoleAssignments",
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
		Values *[]models.UnifiedRoleAssignmentCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListRoleManagementEntitlementManagementTransitiveRoleAssignmentsComplete retrieves all the results into a single object
func (c RoleManagementEntitlementManagementTransitiveRoleAssignmentClient) ListRoleManagementEntitlementManagementTransitiveRoleAssignmentsComplete(ctx context.Context) (ListRoleManagementEntitlementManagementTransitiveRoleAssignmentsCompleteResult, error) {
	return c.ListRoleManagementEntitlementManagementTransitiveRoleAssignmentsCompleteMatchingPredicate(ctx, models.UnifiedRoleAssignmentCollectionResponseOperationPredicate{})
}

// ListRoleManagementEntitlementManagementTransitiveRoleAssignmentsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c RoleManagementEntitlementManagementTransitiveRoleAssignmentClient) ListRoleManagementEntitlementManagementTransitiveRoleAssignmentsCompleteMatchingPredicate(ctx context.Context, predicate models.UnifiedRoleAssignmentCollectionResponseOperationPredicate) (result ListRoleManagementEntitlementManagementTransitiveRoleAssignmentsCompleteResult, err error) {
	items := make([]models.UnifiedRoleAssignmentCollectionResponse, 0)

	resp, err := c.ListRoleManagementEntitlementManagementTransitiveRoleAssignments(ctx)
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

	result = ListRoleManagementEntitlementManagementTransitiveRoleAssignmentsCompleteResult{
		Items: items,
	}
	return
}

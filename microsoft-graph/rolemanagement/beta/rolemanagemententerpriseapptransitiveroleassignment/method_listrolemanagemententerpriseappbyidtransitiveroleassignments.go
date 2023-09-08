package rolemanagemententerpriseapptransitiveroleassignment

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

type ListRoleManagementEnterpriseAppByIdTransitiveRoleAssignmentsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.UnifiedRoleAssignmentCollectionResponse
}

type ListRoleManagementEnterpriseAppByIdTransitiveRoleAssignmentsCompleteResult struct {
	Items []models.UnifiedRoleAssignmentCollectionResponse
}

// ListRoleManagementEnterpriseAppByIdTransitiveRoleAssignments ...
func (c RoleManagementEnterpriseAppTransitiveRoleAssignmentClient) ListRoleManagementEnterpriseAppByIdTransitiveRoleAssignments(ctx context.Context, id RoleManagementEnterpriseAppId) (result ListRoleManagementEnterpriseAppByIdTransitiveRoleAssignmentsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/transitiveRoleAssignments", id.ID()),
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

// ListRoleManagementEnterpriseAppByIdTransitiveRoleAssignmentsComplete retrieves all the results into a single object
func (c RoleManagementEnterpriseAppTransitiveRoleAssignmentClient) ListRoleManagementEnterpriseAppByIdTransitiveRoleAssignmentsComplete(ctx context.Context, id RoleManagementEnterpriseAppId) (ListRoleManagementEnterpriseAppByIdTransitiveRoleAssignmentsCompleteResult, error) {
	return c.ListRoleManagementEnterpriseAppByIdTransitiveRoleAssignmentsCompleteMatchingPredicate(ctx, id, models.UnifiedRoleAssignmentCollectionResponseOperationPredicate{})
}

// ListRoleManagementEnterpriseAppByIdTransitiveRoleAssignmentsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c RoleManagementEnterpriseAppTransitiveRoleAssignmentClient) ListRoleManagementEnterpriseAppByIdTransitiveRoleAssignmentsCompleteMatchingPredicate(ctx context.Context, id RoleManagementEnterpriseAppId, predicate models.UnifiedRoleAssignmentCollectionResponseOperationPredicate) (result ListRoleManagementEnterpriseAppByIdTransitiveRoleAssignmentsCompleteResult, err error) {
	items := make([]models.UnifiedRoleAssignmentCollectionResponse, 0)

	resp, err := c.ListRoleManagementEnterpriseAppByIdTransitiveRoleAssignments(ctx, id)
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

	result = ListRoleManagementEnterpriseAppByIdTransitiveRoleAssignmentsCompleteResult{
		Items: items,
	}
	return
}

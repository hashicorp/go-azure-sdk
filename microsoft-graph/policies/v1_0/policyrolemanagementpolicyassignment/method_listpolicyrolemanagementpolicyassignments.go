package policyrolemanagementpolicyassignment

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common/v1_0/models"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListPolicyRoleManagementPolicyAssignmentsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.UnifiedRoleManagementPolicyAssignmentCollectionResponse
}

type ListPolicyRoleManagementPolicyAssignmentsCompleteResult struct {
	Items []models.UnifiedRoleManagementPolicyAssignmentCollectionResponse
}

// ListPolicyRoleManagementPolicyAssignments ...
func (c PolicyRoleManagementPolicyAssignmentClient) ListPolicyRoleManagementPolicyAssignments(ctx context.Context) (result ListPolicyRoleManagementPolicyAssignmentsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       "/policies/roleManagementPolicyAssignments",
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
		Values *[]models.UnifiedRoleManagementPolicyAssignmentCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListPolicyRoleManagementPolicyAssignmentsComplete retrieves all the results into a single object
func (c PolicyRoleManagementPolicyAssignmentClient) ListPolicyRoleManagementPolicyAssignmentsComplete(ctx context.Context) (ListPolicyRoleManagementPolicyAssignmentsCompleteResult, error) {
	return c.ListPolicyRoleManagementPolicyAssignmentsCompleteMatchingPredicate(ctx, models.UnifiedRoleManagementPolicyAssignmentCollectionResponseOperationPredicate{})
}

// ListPolicyRoleManagementPolicyAssignmentsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c PolicyRoleManagementPolicyAssignmentClient) ListPolicyRoleManagementPolicyAssignmentsCompleteMatchingPredicate(ctx context.Context, predicate models.UnifiedRoleManagementPolicyAssignmentCollectionResponseOperationPredicate) (result ListPolicyRoleManagementPolicyAssignmentsCompleteResult, err error) {
	items := make([]models.UnifiedRoleManagementPolicyAssignmentCollectionResponse, 0)

	resp, err := c.ListPolicyRoleManagementPolicyAssignments(ctx)
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

	result = ListPolicyRoleManagementPolicyAssignmentsCompleteResult{
		Items: items,
	}
	return
}

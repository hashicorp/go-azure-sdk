package rolemanagementexchangetransitiveroleassignment

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

type ListRoleManagementExchangeTransitiveRoleAssignmentsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.UnifiedRoleAssignmentCollectionResponse
}

type ListRoleManagementExchangeTransitiveRoleAssignmentsCompleteResult struct {
	Items []models.UnifiedRoleAssignmentCollectionResponse
}

// ListRoleManagementExchangeTransitiveRoleAssignments ...
func (c RoleManagementExchangeTransitiveRoleAssignmentClient) ListRoleManagementExchangeTransitiveRoleAssignments(ctx context.Context) (result ListRoleManagementExchangeTransitiveRoleAssignmentsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       "/roleManagement/exchange/transitiveRoleAssignments",
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

// ListRoleManagementExchangeTransitiveRoleAssignmentsComplete retrieves all the results into a single object
func (c RoleManagementExchangeTransitiveRoleAssignmentClient) ListRoleManagementExchangeTransitiveRoleAssignmentsComplete(ctx context.Context) (ListRoleManagementExchangeTransitiveRoleAssignmentsCompleteResult, error) {
	return c.ListRoleManagementExchangeTransitiveRoleAssignmentsCompleteMatchingPredicate(ctx, models.UnifiedRoleAssignmentCollectionResponseOperationPredicate{})
}

// ListRoleManagementExchangeTransitiveRoleAssignmentsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c RoleManagementExchangeTransitiveRoleAssignmentClient) ListRoleManagementExchangeTransitiveRoleAssignmentsCompleteMatchingPredicate(ctx context.Context, predicate models.UnifiedRoleAssignmentCollectionResponseOperationPredicate) (result ListRoleManagementExchangeTransitiveRoleAssignmentsCompleteResult, err error) {
	items := make([]models.UnifiedRoleAssignmentCollectionResponse, 0)

	resp, err := c.ListRoleManagementExchangeTransitiveRoleAssignments(ctx)
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

	result = ListRoleManagementExchangeTransitiveRoleAssignmentsCompleteResult{
		Items: items,
	}
	return
}

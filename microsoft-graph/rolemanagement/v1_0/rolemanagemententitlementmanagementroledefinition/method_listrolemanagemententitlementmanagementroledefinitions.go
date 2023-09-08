package rolemanagemententitlementmanagementroledefinition

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

type ListRoleManagementEntitlementManagementRoleDefinitionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.UnifiedRoleDefinitionCollectionResponse
}

type ListRoleManagementEntitlementManagementRoleDefinitionsCompleteResult struct {
	Items []models.UnifiedRoleDefinitionCollectionResponse
}

// ListRoleManagementEntitlementManagementRoleDefinitions ...
func (c RoleManagementEntitlementManagementRoleDefinitionClient) ListRoleManagementEntitlementManagementRoleDefinitions(ctx context.Context) (result ListRoleManagementEntitlementManagementRoleDefinitionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       "/roleManagement/entitlementManagement/roleDefinitions",
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
		Values *[]models.UnifiedRoleDefinitionCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListRoleManagementEntitlementManagementRoleDefinitionsComplete retrieves all the results into a single object
func (c RoleManagementEntitlementManagementRoleDefinitionClient) ListRoleManagementEntitlementManagementRoleDefinitionsComplete(ctx context.Context) (ListRoleManagementEntitlementManagementRoleDefinitionsCompleteResult, error) {
	return c.ListRoleManagementEntitlementManagementRoleDefinitionsCompleteMatchingPredicate(ctx, models.UnifiedRoleDefinitionCollectionResponseOperationPredicate{})
}

// ListRoleManagementEntitlementManagementRoleDefinitionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c RoleManagementEntitlementManagementRoleDefinitionClient) ListRoleManagementEntitlementManagementRoleDefinitionsCompleteMatchingPredicate(ctx context.Context, predicate models.UnifiedRoleDefinitionCollectionResponseOperationPredicate) (result ListRoleManagementEntitlementManagementRoleDefinitionsCompleteResult, err error) {
	items := make([]models.UnifiedRoleDefinitionCollectionResponse, 0)

	resp, err := c.ListRoleManagementEntitlementManagementRoleDefinitions(ctx)
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

	result = ListRoleManagementEntitlementManagementRoleDefinitionsCompleteResult{
		Items: items,
	}
	return
}

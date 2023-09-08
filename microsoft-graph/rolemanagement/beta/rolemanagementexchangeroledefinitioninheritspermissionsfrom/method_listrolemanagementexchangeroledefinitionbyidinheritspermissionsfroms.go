package rolemanagementexchangeroledefinitioninheritspermissionsfrom

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

type ListRoleManagementExchangeRoleDefinitionByIdInheritsPermissionsFromsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.UnifiedRoleDefinitionCollectionResponse
}

type ListRoleManagementExchangeRoleDefinitionByIdInheritsPermissionsFromsCompleteResult struct {
	Items []models.UnifiedRoleDefinitionCollectionResponse
}

// ListRoleManagementExchangeRoleDefinitionByIdInheritsPermissionsFroms ...
func (c RoleManagementExchangeRoleDefinitionInheritsPermissionsFromClient) ListRoleManagementExchangeRoleDefinitionByIdInheritsPermissionsFroms(ctx context.Context, id RoleManagementExchangeRoleDefinitionId) (result ListRoleManagementExchangeRoleDefinitionByIdInheritsPermissionsFromsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/inheritsPermissionsFrom", id.ID()),
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

// ListRoleManagementExchangeRoleDefinitionByIdInheritsPermissionsFromsComplete retrieves all the results into a single object
func (c RoleManagementExchangeRoleDefinitionInheritsPermissionsFromClient) ListRoleManagementExchangeRoleDefinitionByIdInheritsPermissionsFromsComplete(ctx context.Context, id RoleManagementExchangeRoleDefinitionId) (ListRoleManagementExchangeRoleDefinitionByIdInheritsPermissionsFromsCompleteResult, error) {
	return c.ListRoleManagementExchangeRoleDefinitionByIdInheritsPermissionsFromsCompleteMatchingPredicate(ctx, id, models.UnifiedRoleDefinitionCollectionResponseOperationPredicate{})
}

// ListRoleManagementExchangeRoleDefinitionByIdInheritsPermissionsFromsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c RoleManagementExchangeRoleDefinitionInheritsPermissionsFromClient) ListRoleManagementExchangeRoleDefinitionByIdInheritsPermissionsFromsCompleteMatchingPredicate(ctx context.Context, id RoleManagementExchangeRoleDefinitionId, predicate models.UnifiedRoleDefinitionCollectionResponseOperationPredicate) (result ListRoleManagementExchangeRoleDefinitionByIdInheritsPermissionsFromsCompleteResult, err error) {
	items := make([]models.UnifiedRoleDefinitionCollectionResponse, 0)

	resp, err := c.ListRoleManagementExchangeRoleDefinitionByIdInheritsPermissionsFroms(ctx, id)
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

	result = ListRoleManagementExchangeRoleDefinitionByIdInheritsPermissionsFromsCompleteResult{
		Items: items,
	}
	return
}

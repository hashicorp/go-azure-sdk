package entitlementmanagementroledefinitioninheritspermissionsfrom

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

type CreateEntitlementManagementRoleDefinitionInheritsPermissionsFromOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.UnifiedRoleDefinition
}

// CreateEntitlementManagementRoleDefinitionInheritsPermissionsFrom - Create new navigation property to
// inheritsPermissionsFrom for roleManagement
func (c EntitlementManagementRoleDefinitionInheritsPermissionsFromClient) CreateEntitlementManagementRoleDefinitionInheritsPermissionsFrom(ctx context.Context, id stable.RoleManagementEntitlementManagementRoleDefinitionId, input stable.UnifiedRoleDefinition) (result CreateEntitlementManagementRoleDefinitionInheritsPermissionsFromOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
		},
		HttpMethod: http.MethodPost,
		Path:       fmt.Sprintf("%s/inheritsPermissionsFrom", id.ID()),
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	if err = req.Marshal(input); err != nil {
		return
	}

	var resp *client.Response
	resp, err = req.Execute(ctx)
	if resp != nil {
		result.OData = resp.OData
		result.HttpResponse = resp.Response
	}
	if err != nil {
		return
	}

	var model stable.UnifiedRoleDefinition
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}

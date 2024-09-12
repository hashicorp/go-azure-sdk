package entitlementmanagementaccesspackageaccesspackageresourcerolescope

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

type CreateEntitlementManagementAccessPackageResourceRoleScopeOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.AccessPackageResourceRoleScope
}

// CreateEntitlementManagementAccessPackageResourceRoleScope - Create accessPackageResourceRoleScope. Create a new
// accessPackageResourceRoleScope for adding a resource role to an access package. The access package resource, for a
// group, an app, or a SharePoint Online site, must already exist in the access package catalog, and the originId for
// the resource role retrieved from the list of the resource roles. Once you add the resource role scope to the access
// package, the user will receive this resource role through any current and future access package assignments.
func (c EntitlementManagementAccessPackageAccessPackageResourceRoleScopeClient) CreateEntitlementManagementAccessPackageResourceRoleScope(ctx context.Context, id beta.IdentityGovernanceEntitlementManagementAccessPackageId, input beta.AccessPackageResourceRoleScope) (result CreateEntitlementManagementAccessPackageResourceRoleScopeOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
		},
		HttpMethod: http.MethodPost,
		Path:       fmt.Sprintf("%s/accessPackageResourceRoleScopes", id.ID()),
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

	var model beta.AccessPackageResourceRoleScope
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}

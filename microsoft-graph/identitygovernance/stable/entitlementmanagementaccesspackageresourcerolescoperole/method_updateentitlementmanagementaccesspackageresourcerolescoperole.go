package entitlementmanagementaccesspackageresourcerolescoperole

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

type UpdateEntitlementManagementAccessPackageResourceRoleScopeRoleOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateEntitlementManagementAccessPackageResourceRoleScopeRoleOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultUpdateEntitlementManagementAccessPackageResourceRoleScopeRoleOperationOptions() UpdateEntitlementManagementAccessPackageResourceRoleScopeRoleOperationOptions {
	return UpdateEntitlementManagementAccessPackageResourceRoleScopeRoleOperationOptions{}
}

func (o UpdateEntitlementManagementAccessPackageResourceRoleScopeRoleOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateEntitlementManagementAccessPackageResourceRoleScopeRoleOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateEntitlementManagementAccessPackageResourceRoleScopeRoleOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateEntitlementManagementAccessPackageResourceRoleScopeRole - Update the navigation property role in
// identityGovernance
func (c EntitlementManagementAccessPackageResourceRoleScopeRoleClient) UpdateEntitlementManagementAccessPackageResourceRoleScopeRole(ctx context.Context, id stable.IdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeId, input stable.AccessPackageResourceRole, options UpdateEntitlementManagementAccessPackageResourceRoleScopeRoleOperationOptions) (result UpdateEntitlementManagementAccessPackageResourceRoleScopeRoleOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/role", id.ID()),
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

	return
}

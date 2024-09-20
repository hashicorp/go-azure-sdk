package entitlementmanagementaccesspackageassignmentaccesspackageaccesspackageresourcerolescopeaccesspackageresourceroleaccesspackageresource

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

type UpdateEntitlementManagementAccessPackageAssignmentAccessPackageResourceRoleScopeResourceRoleAccessPackageResourceOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateEntitlementManagementAccessPackageAssignmentAccessPackageResourceRoleScopeResourceRoleAccessPackageResourceOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultUpdateEntitlementManagementAccessPackageAssignmentAccessPackageResourceRoleScopeResourceRoleAccessPackageResourceOperationOptions() UpdateEntitlementManagementAccessPackageAssignmentAccessPackageResourceRoleScopeResourceRoleAccessPackageResourceOperationOptions {
	return UpdateEntitlementManagementAccessPackageAssignmentAccessPackageResourceRoleScopeResourceRoleAccessPackageResourceOperationOptions{}
}

func (o UpdateEntitlementManagementAccessPackageAssignmentAccessPackageResourceRoleScopeResourceRoleAccessPackageResourceOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateEntitlementManagementAccessPackageAssignmentAccessPackageResourceRoleScopeResourceRoleAccessPackageResourceOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateEntitlementManagementAccessPackageAssignmentAccessPackageResourceRoleScopeResourceRoleAccessPackageResourceOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateEntitlementManagementAccessPackageAssignmentAccessPackageResourceRoleScopeResourceRoleAccessPackageResource -
// Update the navigation property accessPackageResource in identityGovernance
func (c EntitlementManagementAccessPackageAssignmentAccessPackageAccessPackageResourceRoleScopeAccessPackageResourceRoleAccessPackageResourceClient) UpdateEntitlementManagementAccessPackageAssignmentAccessPackageResourceRoleScopeResourceRoleAccessPackageResource(ctx context.Context, id beta.IdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAccessPackageResourceRoleScopeId, input beta.AccessPackageResource, options UpdateEntitlementManagementAccessPackageAssignmentAccessPackageResourceRoleScopeResourceRoleAccessPackageResourceOperationOptions) (result UpdateEntitlementManagementAccessPackageAssignmentAccessPackageResourceRoleScopeResourceRoleAccessPackageResourceOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/accessPackageResourceRole/accessPackageResource", id.ID()),
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

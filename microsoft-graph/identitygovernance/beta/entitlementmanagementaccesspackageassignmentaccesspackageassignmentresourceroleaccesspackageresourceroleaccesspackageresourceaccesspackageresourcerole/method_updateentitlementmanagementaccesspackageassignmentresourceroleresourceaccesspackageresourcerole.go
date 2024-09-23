package entitlementmanagementaccesspackageassignmentaccesspackageassignmentresourceroleaccesspackageresourceroleaccesspackageresourceaccesspackageresourcerole

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateEntitlementManagementAccessPackageAssignmentResourceRoleResourceAccessPackageResourceRoleOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateEntitlementManagementAccessPackageAssignmentResourceRoleResourceAccessPackageResourceRoleOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdateEntitlementManagementAccessPackageAssignmentResourceRoleResourceAccessPackageResourceRoleOperationOptions() UpdateEntitlementManagementAccessPackageAssignmentResourceRoleResourceAccessPackageResourceRoleOperationOptions {
	return UpdateEntitlementManagementAccessPackageAssignmentResourceRoleResourceAccessPackageResourceRoleOperationOptions{}
}

func (o UpdateEntitlementManagementAccessPackageAssignmentResourceRoleResourceAccessPackageResourceRoleOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateEntitlementManagementAccessPackageAssignmentResourceRoleResourceAccessPackageResourceRoleOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateEntitlementManagementAccessPackageAssignmentResourceRoleResourceAccessPackageResourceRoleOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateEntitlementManagementAccessPackageAssignmentResourceRoleResourceAccessPackageResourceRole - Update the
// navigation property accessPackageResourceRoles in identityGovernance
func (c EntitlementManagementAccessPackageAssignmentAccessPackageAssignmentResourceRoleAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceRoleClient) UpdateEntitlementManagementAccessPackageAssignmentResourceRoleResourceAccessPackageResourceRole(ctx context.Context, id beta.IdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAssignmentResourceRoleIdAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceRoleId, input beta.AccessPackageResourceRole, options UpdateEntitlementManagementAccessPackageAssignmentResourceRoleResourceAccessPackageResourceRoleOperationOptions) (result UpdateEntitlementManagementAccessPackageAssignmentResourceRoleResourceAccessPackageResourceRoleOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          id.ID(),
		RetryFunc:     options.RetryFunc,
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

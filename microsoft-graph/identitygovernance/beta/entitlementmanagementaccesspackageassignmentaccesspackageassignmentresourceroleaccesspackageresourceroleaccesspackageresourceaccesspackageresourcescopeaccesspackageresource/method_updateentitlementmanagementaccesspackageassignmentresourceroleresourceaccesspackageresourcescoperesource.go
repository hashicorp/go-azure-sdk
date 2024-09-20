package entitlementmanagementaccesspackageassignmentaccesspackageassignmentresourceroleaccesspackageresourceroleaccesspackageresourceaccesspackageresourcescopeaccesspackageresource

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

type UpdateEntitlementManagementAccessPackageAssignmentResourceRoleResourceAccessPackageResourceScopeResourceOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateEntitlementManagementAccessPackageAssignmentResourceRoleResourceAccessPackageResourceScopeResourceOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultUpdateEntitlementManagementAccessPackageAssignmentResourceRoleResourceAccessPackageResourceScopeResourceOperationOptions() UpdateEntitlementManagementAccessPackageAssignmentResourceRoleResourceAccessPackageResourceScopeResourceOperationOptions {
	return UpdateEntitlementManagementAccessPackageAssignmentResourceRoleResourceAccessPackageResourceScopeResourceOperationOptions{}
}

func (o UpdateEntitlementManagementAccessPackageAssignmentResourceRoleResourceAccessPackageResourceScopeResourceOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateEntitlementManagementAccessPackageAssignmentResourceRoleResourceAccessPackageResourceScopeResourceOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateEntitlementManagementAccessPackageAssignmentResourceRoleResourceAccessPackageResourceScopeResourceOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateEntitlementManagementAccessPackageAssignmentResourceRoleResourceAccessPackageResourceScopeResource - Update the
// navigation property accessPackageResource in identityGovernance
func (c EntitlementManagementAccessPackageAssignmentAccessPackageAssignmentResourceRoleAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceScopeAccessPackageResourceClient) UpdateEntitlementManagementAccessPackageAssignmentResourceRoleResourceAccessPackageResourceScopeResource(ctx context.Context, id beta.IdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAssignmentResourceRoleIdAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceScopeId, input beta.AccessPackageResource, options UpdateEntitlementManagementAccessPackageAssignmentResourceRoleResourceAccessPackageResourceScopeResourceOperationOptions) (result UpdateEntitlementManagementAccessPackageAssignmentResourceRoleResourceAccessPackageResourceScopeResourceOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/accessPackageResource", id.ID()),
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

package entitlementmanagementaccesspackageaccesspackageresourcerolescopeaccesspackageresourceroleaccesspackageresourceaccesspackageresourcescopeaccesspackageresource

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

type UpdateEntitlementManagementAccessPackageResourceRoleScopeResourceRoleAccessPackageResourceResourceScopeAccessPackageResourceOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateEntitlementManagementAccessPackageResourceRoleScopeResourceRoleAccessPackageResourceResourceScopeAccessPackageResourceOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdateEntitlementManagementAccessPackageResourceRoleScopeResourceRoleAccessPackageResourceResourceScopeAccessPackageResourceOperationOptions() UpdateEntitlementManagementAccessPackageResourceRoleScopeResourceRoleAccessPackageResourceResourceScopeAccessPackageResourceOperationOptions {
	return UpdateEntitlementManagementAccessPackageResourceRoleScopeResourceRoleAccessPackageResourceResourceScopeAccessPackageResourceOperationOptions{}
}

func (o UpdateEntitlementManagementAccessPackageResourceRoleScopeResourceRoleAccessPackageResourceResourceScopeAccessPackageResourceOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateEntitlementManagementAccessPackageResourceRoleScopeResourceRoleAccessPackageResourceResourceScopeAccessPackageResourceOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateEntitlementManagementAccessPackageResourceRoleScopeResourceRoleAccessPackageResourceResourceScopeAccessPackageResourceOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateEntitlementManagementAccessPackageResourceRoleScopeResourceRoleAccessPackageResourceResourceScopeAccessPackageResource
// - Update the navigation property accessPackageResource in identityGovernance
func (c EntitlementManagementAccessPackageAccessPackageResourceRoleScopeAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceScopeAccessPackageResourceClient) UpdateEntitlementManagementAccessPackageResourceRoleScopeResourceRoleAccessPackageResourceResourceScopeAccessPackageResource(ctx context.Context, id beta.IdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageResourceRoleScopeIdAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceScopeId, input beta.AccessPackageResource, options UpdateEntitlementManagementAccessPackageResourceRoleScopeResourceRoleAccessPackageResourceResourceScopeAccessPackageResourceOperationOptions) (result UpdateEntitlementManagementAccessPackageResourceRoleScopeResourceRoleAccessPackageResourceResourceScopeAccessPackageResourceOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/accessPackageResource", id.ID()),
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

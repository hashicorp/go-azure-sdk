package entitlementmanagementaccesspackageresourcerolescopeaccesspackageresourceroleaccesspackageresourceaccesspackageresourcescopeaccesspackageresource

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

type UpdateEntitlementManagementAccessPackageResourceRoleScopeAccessAccessScopeAccessPackageResourceOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateEntitlementManagementAccessPackageResourceRoleScopeAccessAccessScopeAccessPackageResourceOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdateEntitlementManagementAccessPackageResourceRoleScopeAccessAccessScopeAccessPackageResourceOperationOptions() UpdateEntitlementManagementAccessPackageResourceRoleScopeAccessAccessScopeAccessPackageResourceOperationOptions {
	return UpdateEntitlementManagementAccessPackageResourceRoleScopeAccessAccessScopeAccessPackageResourceOperationOptions{}
}

func (o UpdateEntitlementManagementAccessPackageResourceRoleScopeAccessAccessScopeAccessPackageResourceOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateEntitlementManagementAccessPackageResourceRoleScopeAccessAccessScopeAccessPackageResourceOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateEntitlementManagementAccessPackageResourceRoleScopeAccessAccessScopeAccessPackageResourceOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateEntitlementManagementAccessPackageResourceRoleScopeAccessAccessScopeAccessPackageResource - Update the
// navigation property accessPackageResource in identityGovernance
func (c EntitlementManagementAccessPackageResourceRoleScopeAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceScopeAccessPackageResourceClient) UpdateEntitlementManagementAccessPackageResourceRoleScopeAccessAccessScopeAccessPackageResource(ctx context.Context, id beta.IdentityGovernanceEntitlementManagementAccessPackageResourceRoleScopeIdAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceScopeId, input beta.AccessPackageResource, options UpdateEntitlementManagementAccessPackageResourceRoleScopeAccessAccessScopeAccessPackageResourceOperationOptions) (result UpdateEntitlementManagementAccessPackageResourceRoleScopeAccessAccessScopeAccessPackageResourceOperationResponse, err error) {
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

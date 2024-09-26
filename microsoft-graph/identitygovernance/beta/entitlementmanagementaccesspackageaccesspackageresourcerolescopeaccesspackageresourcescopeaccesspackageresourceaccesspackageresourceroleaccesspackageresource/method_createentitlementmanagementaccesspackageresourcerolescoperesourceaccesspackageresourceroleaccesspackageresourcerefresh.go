package entitlementmanagementaccesspackageaccesspackageresourcerolescopeaccesspackageresourcescopeaccesspackageresourceaccesspackageresourceroleaccesspackageresource

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

type CreateEntitlementManagementAccessPackageResourceRoleScopeResourceAccessPackageResourceRoleAccessPackageResourceRefreshOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type CreateEntitlementManagementAccessPackageResourceRoleScopeResourceAccessPackageResourceRoleAccessPackageResourceRefreshOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateEntitlementManagementAccessPackageResourceRoleScopeResourceAccessPackageResourceRoleAccessPackageResourceRefreshOperationOptions() CreateEntitlementManagementAccessPackageResourceRoleScopeResourceAccessPackageResourceRoleAccessPackageResourceRefreshOperationOptions {
	return CreateEntitlementManagementAccessPackageResourceRoleScopeResourceAccessPackageResourceRoleAccessPackageResourceRefreshOperationOptions{}
}

func (o CreateEntitlementManagementAccessPackageResourceRoleScopeResourceAccessPackageResourceRoleAccessPackageResourceRefreshOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateEntitlementManagementAccessPackageResourceRoleScopeResourceAccessPackageResourceRoleAccessPackageResourceRefreshOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateEntitlementManagementAccessPackageResourceRoleScopeResourceAccessPackageResourceRoleAccessPackageResourceRefreshOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateEntitlementManagementAccessPackageResourceRoleScopeResourceAccessPackageResourceRoleAccessPackageResourceRefresh
// - Invoke action refresh. In Microsoft Entra entitlement management, refresh the accessPackageResource object to fetch
// the latest details for displayName, description, and resourceType from the origin system. For the AadApplication
// originSystem, this operation also updates the displayName and description for the accessPackageResourceRole.
func (c EntitlementManagementAccessPackageAccessPackageResourceRoleScopeAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRoleAccessPackageResourceClient) CreateEntitlementManagementAccessPackageResourceRoleScopeResourceAccessPackageResourceRoleAccessPackageResourceRefresh(ctx context.Context, id beta.IdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageResourceRoleScopeIdAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRoleId, options CreateEntitlementManagementAccessPackageResourceRoleScopeResourceAccessPackageResourceRoleAccessPackageResourceRefreshOperationOptions) (result CreateEntitlementManagementAccessPackageResourceRoleScopeResourceAccessPackageResourceRoleAccessPackageResourceRefreshOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusCreated,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/accessPackageResource/refresh", id.ID()),
		RetryFunc:     options.RetryFunc,
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
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

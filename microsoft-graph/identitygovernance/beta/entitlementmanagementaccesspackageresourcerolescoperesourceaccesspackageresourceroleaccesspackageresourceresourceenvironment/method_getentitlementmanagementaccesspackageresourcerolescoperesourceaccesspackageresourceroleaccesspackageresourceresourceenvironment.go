package entitlementmanagementaccesspackageresourcerolescoperesourceaccesspackageresourceroleaccesspackageresourceresourceenvironment

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

type GetEntitlementManagementAccessPackageResourceRoleScopeResourceAccessPackageResourceRoleAccessPackageResourceResourceEnvironmentOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.AccessPackageResourceEnvironment
}

type GetEntitlementManagementAccessPackageResourceRoleScopeResourceAccessPackageResourceRoleAccessPackageResourceResourceEnvironmentOperationOptions struct {
	Expand *odata.Expand
	Select *[]string
}

func DefaultGetEntitlementManagementAccessPackageResourceRoleScopeResourceAccessPackageResourceRoleAccessPackageResourceResourceEnvironmentOperationOptions() GetEntitlementManagementAccessPackageResourceRoleScopeResourceAccessPackageResourceRoleAccessPackageResourceResourceEnvironmentOperationOptions {
	return GetEntitlementManagementAccessPackageResourceRoleScopeResourceAccessPackageResourceRoleAccessPackageResourceResourceEnvironmentOperationOptions{}
}

func (o GetEntitlementManagementAccessPackageResourceRoleScopeResourceAccessPackageResourceRoleAccessPackageResourceResourceEnvironmentOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetEntitlementManagementAccessPackageResourceRoleScopeResourceAccessPackageResourceRoleAccessPackageResourceResourceEnvironmentOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetEntitlementManagementAccessPackageResourceRoleScopeResourceAccessPackageResourceRoleAccessPackageResourceResourceEnvironmentOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetEntitlementManagementAccessPackageResourceRoleScopeResourceAccessPackageResourceRoleAccessPackageResourceResourceEnvironment
// - Get accessPackageResourceEnvironment from identityGovernance. Contains the environment information for the
// resource. This environment can be set using either the @odata.bind annotation or the environment's originId. Supports
// $expand.
func (c EntitlementManagementAccessPackageResourceRoleScopeResourceAccessPackageResourceRoleAccessPackageResourceResourceEnvironmentClient) GetEntitlementManagementAccessPackageResourceRoleScopeResourceAccessPackageResourceRoleAccessPackageResourceResourceEnvironment(ctx context.Context, id beta.IdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageResourceRoleScopeIdAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRoleId, options GetEntitlementManagementAccessPackageResourceRoleScopeResourceAccessPackageResourceRoleAccessPackageResourceResourceEnvironmentOperationOptions) (result GetEntitlementManagementAccessPackageResourceRoleScopeResourceAccessPackageResourceRoleAccessPackageResourceResourceEnvironmentOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/accessPackageResource/accessPackageResourceEnvironment", id.ID()),
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

	var model beta.AccessPackageResourceEnvironment
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}

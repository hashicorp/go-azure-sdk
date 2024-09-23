package entitlementmanagementaccesspackageaccesspackageresourcerolescopeaccesspackageresourceroleaccesspackageresourceaccesspackageresourcescopeaccesspackageresourceaccesspackageresourceenvironment

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

type GetEntitlementManagementAccessPackageResourceRoleScopeResourceRoleAccessPackageResourceResourceAccessPackageResourceEnvironmentOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.AccessPackageResourceEnvironment
}

type GetEntitlementManagementAccessPackageResourceRoleScopeResourceRoleAccessPackageResourceResourceAccessPackageResourceEnvironmentOperationOptions struct {
	Expand    *odata.Expand
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Select    *[]string
}

func DefaultGetEntitlementManagementAccessPackageResourceRoleScopeResourceRoleAccessPackageResourceResourceAccessPackageResourceEnvironmentOperationOptions() GetEntitlementManagementAccessPackageResourceRoleScopeResourceRoleAccessPackageResourceResourceAccessPackageResourceEnvironmentOperationOptions {
	return GetEntitlementManagementAccessPackageResourceRoleScopeResourceRoleAccessPackageResourceResourceAccessPackageResourceEnvironmentOperationOptions{}
}

func (o GetEntitlementManagementAccessPackageResourceRoleScopeResourceRoleAccessPackageResourceResourceAccessPackageResourceEnvironmentOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetEntitlementManagementAccessPackageResourceRoleScopeResourceRoleAccessPackageResourceResourceAccessPackageResourceEnvironmentOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetEntitlementManagementAccessPackageResourceRoleScopeResourceRoleAccessPackageResourceResourceAccessPackageResourceEnvironmentOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetEntitlementManagementAccessPackageResourceRoleScopeResourceRoleAccessPackageResourceResourceAccessPackageResourceEnvironment
// - Get accessPackageResourceEnvironment from identityGovernance. Contains the environment information for the
// resource. This environment can be set using either the @odata.bind annotation or the environment's originId. Supports
// $expand.
func (c EntitlementManagementAccessPackageAccessPackageResourceRoleScopeAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceEnvironmentClient) GetEntitlementManagementAccessPackageResourceRoleScopeResourceRoleAccessPackageResourceResourceAccessPackageResourceEnvironment(ctx context.Context, id beta.IdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageResourceRoleScopeIdAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceScopeId, options GetEntitlementManagementAccessPackageResourceRoleScopeResourceRoleAccessPackageResourceResourceAccessPackageResourceEnvironmentOperationOptions) (result GetEntitlementManagementAccessPackageResourceRoleScopeResourceRoleAccessPackageResourceResourceAccessPackageResourceEnvironmentOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/accessPackageResource/accessPackageResourceEnvironment", id.ID()),
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

	var model beta.AccessPackageResourceEnvironment
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}

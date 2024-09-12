package entitlementmanagementresourceenvironmentresourceroleresourcescoperesourceenvironment

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

type GetEntitlementManagementResourceEnvironmentResourceRoleResourceScopeResourceEnvironmentOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.AccessPackageResourceEnvironment
}

type GetEntitlementManagementResourceEnvironmentResourceRoleResourceScopeResourceEnvironmentOperationOptions struct {
	Expand *odata.Expand
	Select *[]string
}

func DefaultGetEntitlementManagementResourceEnvironmentResourceRoleResourceScopeResourceEnvironmentOperationOptions() GetEntitlementManagementResourceEnvironmentResourceRoleResourceScopeResourceEnvironmentOperationOptions {
	return GetEntitlementManagementResourceEnvironmentResourceRoleResourceScopeResourceEnvironmentOperationOptions{}
}

func (o GetEntitlementManagementResourceEnvironmentResourceRoleResourceScopeResourceEnvironmentOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetEntitlementManagementResourceEnvironmentResourceRoleResourceScopeResourceEnvironmentOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetEntitlementManagementResourceEnvironmentResourceRoleResourceScopeResourceEnvironmentOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetEntitlementManagementResourceEnvironmentResourceRoleResourceScopeResourceEnvironment - Get environment from
// identityGovernance. Contains the environment information for the resource. This can be set using either the
// @odata.bind annotation or the environment's originId.Supports $expand.
func (c EntitlementManagementResourceEnvironmentResourceRoleResourceScopeResourceEnvironmentClient) GetEntitlementManagementResourceEnvironmentResourceRoleResourceScopeResourceEnvironment(ctx context.Context, id stable.IdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdRoleIdResourceScopeId, options GetEntitlementManagementResourceEnvironmentResourceRoleResourceScopeResourceEnvironmentOperationOptions) (result GetEntitlementManagementResourceEnvironmentResourceRoleResourceScopeResourceEnvironmentOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/resource/environment", id.ID()),
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

	var model stable.AccessPackageResourceEnvironment
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}

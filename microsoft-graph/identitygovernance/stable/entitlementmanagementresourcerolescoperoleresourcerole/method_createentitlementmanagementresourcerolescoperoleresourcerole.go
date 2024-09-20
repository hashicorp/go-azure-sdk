package entitlementmanagementresourcerolescoperoleresourcerole

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

type CreateEntitlementManagementResourceRoleScopeRoleResourceRoleOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.AccessPackageResourceRole
}

type CreateEntitlementManagementResourceRoleScopeRoleResourceRoleOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultCreateEntitlementManagementResourceRoleScopeRoleResourceRoleOperationOptions() CreateEntitlementManagementResourceRoleScopeRoleResourceRoleOperationOptions {
	return CreateEntitlementManagementResourceRoleScopeRoleResourceRoleOperationOptions{}
}

func (o CreateEntitlementManagementResourceRoleScopeRoleResourceRoleOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateEntitlementManagementResourceRoleScopeRoleResourceRoleOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateEntitlementManagementResourceRoleScopeRoleResourceRoleOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateEntitlementManagementResourceRoleScopeRoleResourceRole - Create new navigation property to roles for
// identityGovernance
func (c EntitlementManagementResourceRoleScopeRoleResourceRoleClient) CreateEntitlementManagementResourceRoleScopeRoleResourceRole(ctx context.Context, id stable.IdentityGovernanceEntitlementManagementResourceRoleScopeId, input stable.AccessPackageResourceRole, options CreateEntitlementManagementResourceRoleScopeRoleResourceRoleOperationOptions) (result CreateEntitlementManagementResourceRoleScopeRoleResourceRoleOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/role/resource/roles", id.ID()),
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

	var model stable.AccessPackageResourceRole
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}

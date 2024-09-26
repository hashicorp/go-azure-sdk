package entitlementmanagementaccesspackageaccesspackageresourcerolescopeaccesspackageresourceroleaccesspackageresourceaccesspackageresourcerole

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

type CreateEntitlementManagementAccessPackageResourceRoleScopeResourceRoleAccessPackageResourceResourceRoleOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.AccessPackageResourceRole
}

type CreateEntitlementManagementAccessPackageResourceRoleScopeResourceRoleAccessPackageResourceResourceRoleOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateEntitlementManagementAccessPackageResourceRoleScopeResourceRoleAccessPackageResourceResourceRoleOperationOptions() CreateEntitlementManagementAccessPackageResourceRoleScopeResourceRoleAccessPackageResourceResourceRoleOperationOptions {
	return CreateEntitlementManagementAccessPackageResourceRoleScopeResourceRoleAccessPackageResourceResourceRoleOperationOptions{}
}

func (o CreateEntitlementManagementAccessPackageResourceRoleScopeResourceRoleAccessPackageResourceResourceRoleOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateEntitlementManagementAccessPackageResourceRoleScopeResourceRoleAccessPackageResourceResourceRoleOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateEntitlementManagementAccessPackageResourceRoleScopeResourceRoleAccessPackageResourceResourceRoleOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateEntitlementManagementAccessPackageResourceRoleScopeResourceRoleAccessPackageResourceResourceRole - Create new
// navigation property to accessPackageResourceRoles for identityGovernance
func (c EntitlementManagementAccessPackageAccessPackageResourceRoleScopeAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceRoleClient) CreateEntitlementManagementAccessPackageResourceRoleScopeResourceRoleAccessPackageResourceResourceRole(ctx context.Context, id beta.IdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageResourceRoleScopeId, input beta.AccessPackageResourceRole, options CreateEntitlementManagementAccessPackageResourceRoleScopeResourceRoleAccessPackageResourceResourceRoleOperationOptions) (result CreateEntitlementManagementAccessPackageResourceRoleScopeResourceRoleAccessPackageResourceResourceRoleOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/accessPackageResourceRole/accessPackageResource/accessPackageResourceRoles", id.ID()),
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

	var model beta.AccessPackageResourceRole
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}

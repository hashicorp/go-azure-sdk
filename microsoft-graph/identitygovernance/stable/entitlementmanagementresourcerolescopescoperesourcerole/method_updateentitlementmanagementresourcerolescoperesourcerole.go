package entitlementmanagementresourcerolescopescoperesourcerole

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateEntitlementManagementResourceRoleScopeResourceRoleOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateEntitlementManagementResourceRoleScopeResourceRoleOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdateEntitlementManagementResourceRoleScopeResourceRoleOperationOptions() UpdateEntitlementManagementResourceRoleScopeResourceRoleOperationOptions {
	return UpdateEntitlementManagementResourceRoleScopeResourceRoleOperationOptions{}
}

func (o UpdateEntitlementManagementResourceRoleScopeResourceRoleOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateEntitlementManagementResourceRoleScopeResourceRoleOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateEntitlementManagementResourceRoleScopeResourceRoleOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateEntitlementManagementResourceRoleScopeResourceRole - Update the navigation property roles in identityGovernance
func (c EntitlementManagementResourceRoleScopeScopeResourceRoleClient) UpdateEntitlementManagementResourceRoleScopeResourceRole(ctx context.Context, id stable.IdentityGovernanceEntitlementManagementResourceRoleScopeIdScopeResourceRoleId, input stable.AccessPackageResourceRole, options UpdateEntitlementManagementResourceRoleScopeResourceRoleOperationOptions) (result UpdateEntitlementManagementResourceRoleScopeResourceRoleOperationResponse, err error) {
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

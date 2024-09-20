package entitlementmanagementresourcerolescoperoleresourcescoperesource

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

type UpdateEntitlementManagementResourceRoleScopeRoleResourceScopeResourceOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateEntitlementManagementResourceRoleScopeRoleResourceScopeResourceOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultUpdateEntitlementManagementResourceRoleScopeRoleResourceScopeResourceOperationOptions() UpdateEntitlementManagementResourceRoleScopeRoleResourceScopeResourceOperationOptions {
	return UpdateEntitlementManagementResourceRoleScopeRoleResourceScopeResourceOperationOptions{}
}

func (o UpdateEntitlementManagementResourceRoleScopeRoleResourceScopeResourceOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateEntitlementManagementResourceRoleScopeRoleResourceScopeResourceOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateEntitlementManagementResourceRoleScopeRoleResourceScopeResourceOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateEntitlementManagementResourceRoleScopeRoleResourceScopeResource - Update the navigation property resource in
// identityGovernance
func (c EntitlementManagementResourceRoleScopeRoleResourceScopeResourceClient) UpdateEntitlementManagementResourceRoleScopeRoleResourceScopeResource(ctx context.Context, id stable.IdentityGovernanceEntitlementManagementResourceRoleScopeIdRoleResourceScopeId, input stable.AccessPackageResource, options UpdateEntitlementManagementResourceRoleScopeRoleResourceScopeResourceOperationOptions) (result UpdateEntitlementManagementResourceRoleScopeRoleResourceScopeResourceOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/resource", id.ID()),
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

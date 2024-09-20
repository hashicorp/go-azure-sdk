package entitlementmanagementresourceroleresourcescope

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateEntitlementManagementResourceRoleResourceScopeOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateEntitlementManagementResourceRoleResourceScopeOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultUpdateEntitlementManagementResourceRoleResourceScopeOperationOptions() UpdateEntitlementManagementResourceRoleResourceScopeOperationOptions {
	return UpdateEntitlementManagementResourceRoleResourceScopeOperationOptions{}
}

func (o UpdateEntitlementManagementResourceRoleResourceScopeOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateEntitlementManagementResourceRoleResourceScopeOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateEntitlementManagementResourceRoleResourceScopeOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateEntitlementManagementResourceRoleResourceScope - Update the navigation property scopes in identityGovernance
func (c EntitlementManagementResourceRoleResourceScopeClient) UpdateEntitlementManagementResourceRoleResourceScope(ctx context.Context, id stable.IdentityGovernanceEntitlementManagementResourceIdRoleIdResourceScopeId, input stable.AccessPackageResourceScope, options UpdateEntitlementManagementResourceRoleResourceScopeOperationOptions) (result UpdateEntitlementManagementResourceRoleResourceScopeOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          id.ID(),
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

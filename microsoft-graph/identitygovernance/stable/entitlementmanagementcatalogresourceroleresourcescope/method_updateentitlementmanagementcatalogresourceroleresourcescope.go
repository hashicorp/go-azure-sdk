package entitlementmanagementcatalogresourceroleresourcescope

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateEntitlementManagementCatalogResourceRoleResourceScopeOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateEntitlementManagementCatalogResourceRoleResourceScopeOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultUpdateEntitlementManagementCatalogResourceRoleResourceScopeOperationOptions() UpdateEntitlementManagementCatalogResourceRoleResourceScopeOperationOptions {
	return UpdateEntitlementManagementCatalogResourceRoleResourceScopeOperationOptions{}
}

func (o UpdateEntitlementManagementCatalogResourceRoleResourceScopeOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateEntitlementManagementCatalogResourceRoleResourceScopeOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateEntitlementManagementCatalogResourceRoleResourceScopeOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateEntitlementManagementCatalogResourceRoleResourceScope - Update the navigation property scopes in
// identityGovernance
func (c EntitlementManagementCatalogResourceRoleResourceScopeClient) UpdateEntitlementManagementCatalogResourceRoleResourceScope(ctx context.Context, id stable.IdentityGovernanceEntitlementManagementCatalogIdResourceIdRoleIdResourceScopeId, input stable.AccessPackageResourceScope, options UpdateEntitlementManagementCatalogResourceRoleResourceScopeOperationOptions) (result UpdateEntitlementManagementCatalogResourceRoleResourceScopeOperationResponse, err error) {
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

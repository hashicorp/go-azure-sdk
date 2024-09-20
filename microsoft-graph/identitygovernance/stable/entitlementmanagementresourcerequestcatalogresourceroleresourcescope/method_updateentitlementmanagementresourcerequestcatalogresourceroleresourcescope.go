package entitlementmanagementresourcerequestcatalogresourceroleresourcescope

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateEntitlementManagementResourceRequestCatalogResourceRoleResourceScopeOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateEntitlementManagementResourceRequestCatalogResourceRoleResourceScopeOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultUpdateEntitlementManagementResourceRequestCatalogResourceRoleResourceScopeOperationOptions() UpdateEntitlementManagementResourceRequestCatalogResourceRoleResourceScopeOperationOptions {
	return UpdateEntitlementManagementResourceRequestCatalogResourceRoleResourceScopeOperationOptions{}
}

func (o UpdateEntitlementManagementResourceRequestCatalogResourceRoleResourceScopeOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateEntitlementManagementResourceRequestCatalogResourceRoleResourceScopeOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateEntitlementManagementResourceRequestCatalogResourceRoleResourceScopeOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateEntitlementManagementResourceRequestCatalogResourceRoleResourceScope - Update the navigation property scopes in
// identityGovernance
func (c EntitlementManagementResourceRequestCatalogResourceRoleResourceScopeClient) UpdateEntitlementManagementResourceRequestCatalogResourceRoleResourceScope(ctx context.Context, id stable.IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceIdRoleIdResourceScopeId, input stable.AccessPackageResourceScope, options UpdateEntitlementManagementResourceRequestCatalogResourceRoleResourceScopeOperationOptions) (result UpdateEntitlementManagementResourceRequestCatalogResourceRoleResourceScopeOperationResponse, err error) {
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

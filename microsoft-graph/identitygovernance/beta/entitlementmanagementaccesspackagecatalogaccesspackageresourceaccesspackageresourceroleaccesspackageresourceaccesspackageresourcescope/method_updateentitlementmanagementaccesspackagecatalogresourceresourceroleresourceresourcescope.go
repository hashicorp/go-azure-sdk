package entitlementmanagementaccesspackagecatalogaccesspackageresourceaccesspackageresourceroleaccesspackageresourceaccesspackageresourcescope

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateEntitlementManagementAccessPackageCatalogResourceResourceRoleResourceResourceScopeOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateEntitlementManagementAccessPackageCatalogResourceResourceRoleResourceResourceScopeOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdateEntitlementManagementAccessPackageCatalogResourceResourceRoleResourceResourceScopeOperationOptions() UpdateEntitlementManagementAccessPackageCatalogResourceResourceRoleResourceResourceScopeOperationOptions {
	return UpdateEntitlementManagementAccessPackageCatalogResourceResourceRoleResourceResourceScopeOperationOptions{}
}

func (o UpdateEntitlementManagementAccessPackageCatalogResourceResourceRoleResourceResourceScopeOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateEntitlementManagementAccessPackageCatalogResourceResourceRoleResourceResourceScopeOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateEntitlementManagementAccessPackageCatalogResourceResourceRoleResourceResourceScopeOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateEntitlementManagementAccessPackageCatalogResourceResourceRoleResourceResourceScope - Update the navigation
// property accessPackageResourceScopes in identityGovernance
func (c EntitlementManagementAccessPackageCatalogAccessPackageResourceAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceScopeClient) UpdateEntitlementManagementAccessPackageCatalogResourceResourceRoleResourceResourceScope(ctx context.Context, id beta.IdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageResourceIdAccessPackageResourceRoleIdAccessPackageResourceAccessPackageResourceScopeId, input beta.AccessPackageResourceScope, options UpdateEntitlementManagementAccessPackageCatalogResourceResourceRoleResourceResourceScopeOperationOptions) (result UpdateEntitlementManagementAccessPackageCatalogResourceResourceRoleResourceResourceScopeOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusNoContent,
			http.StatusOK,
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

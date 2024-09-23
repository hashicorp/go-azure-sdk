package entitlementmanagementaccesspackagecatalogaccesspackageresourceaccesspackageresourceroleaccesspackageresourceaccesspackageresourcescope

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

type CreateEntitlementManagementAccessPackageCatalogResourceResourceRoleResourceResourceScopeOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.AccessPackageResourceScope
}

type CreateEntitlementManagementAccessPackageCatalogResourceResourceRoleResourceResourceScopeOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateEntitlementManagementAccessPackageCatalogResourceResourceRoleResourceResourceScopeOperationOptions() CreateEntitlementManagementAccessPackageCatalogResourceResourceRoleResourceResourceScopeOperationOptions {
	return CreateEntitlementManagementAccessPackageCatalogResourceResourceRoleResourceResourceScopeOperationOptions{}
}

func (o CreateEntitlementManagementAccessPackageCatalogResourceResourceRoleResourceResourceScopeOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateEntitlementManagementAccessPackageCatalogResourceResourceRoleResourceResourceScopeOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateEntitlementManagementAccessPackageCatalogResourceResourceRoleResourceResourceScopeOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateEntitlementManagementAccessPackageCatalogResourceResourceRoleResourceResourceScope - Create new navigation
// property to accessPackageResourceScopes for identityGovernance
func (c EntitlementManagementAccessPackageCatalogAccessPackageResourceAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceScopeClient) CreateEntitlementManagementAccessPackageCatalogResourceResourceRoleResourceResourceScope(ctx context.Context, id beta.IdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageResourceIdAccessPackageResourceRoleId, input beta.AccessPackageResourceScope, options CreateEntitlementManagementAccessPackageCatalogResourceResourceRoleResourceResourceScopeOperationOptions) (result CreateEntitlementManagementAccessPackageCatalogResourceResourceRoleResourceResourceScopeOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/accessPackageResource/accessPackageResourceScopes", id.ID()),
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

	var model beta.AccessPackageResourceScope
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}

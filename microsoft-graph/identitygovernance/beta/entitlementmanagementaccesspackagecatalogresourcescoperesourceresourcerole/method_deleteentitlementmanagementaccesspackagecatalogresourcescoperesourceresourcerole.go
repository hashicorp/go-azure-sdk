package entitlementmanagementaccesspackagecatalogresourcescoperesourceresourcerole

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

type DeleteEntitlementManagementAccessPackageCatalogResourceScopeResourceResourceRoleOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type DeleteEntitlementManagementAccessPackageCatalogResourceScopeResourceResourceRoleOperationOptions struct {
	IfMatch *string
}

func DefaultDeleteEntitlementManagementAccessPackageCatalogResourceScopeResourceResourceRoleOperationOptions() DeleteEntitlementManagementAccessPackageCatalogResourceScopeResourceResourceRoleOperationOptions {
	return DeleteEntitlementManagementAccessPackageCatalogResourceScopeResourceResourceRoleOperationOptions{}
}

func (o DeleteEntitlementManagementAccessPackageCatalogResourceScopeResourceResourceRoleOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o DeleteEntitlementManagementAccessPackageCatalogResourceScopeResourceResourceRoleOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o DeleteEntitlementManagementAccessPackageCatalogResourceScopeResourceResourceRoleOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// DeleteEntitlementManagementAccessPackageCatalogResourceScopeResourceResourceRole - Delete navigation property
// accessPackageResourceRoles for identityGovernance
func (c EntitlementManagementAccessPackageCatalogResourceScopeResourceResourceRoleClient) DeleteEntitlementManagementAccessPackageCatalogResourceScopeResourceResourceRole(ctx context.Context, id beta.IdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageResourceScopeIdAccessPackageResourceAccessPackageResourceRoleId, options DeleteEntitlementManagementAccessPackageCatalogResourceScopeResourceResourceRoleOperationOptions) (result DeleteEntitlementManagementAccessPackageCatalogResourceScopeResourceResourceRoleOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodDelete,
		OptionsObject: options,
		Path:          id.ID(),
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

	return
}

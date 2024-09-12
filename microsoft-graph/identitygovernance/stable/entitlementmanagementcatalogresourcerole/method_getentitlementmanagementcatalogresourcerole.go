package entitlementmanagementcatalogresourcerole

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetEntitlementManagementCatalogResourceRoleOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.AccessPackageResourceRole
}

type GetEntitlementManagementCatalogResourceRoleOperationOptions struct {
	Expand *odata.Expand
	Select *[]string
}

func DefaultGetEntitlementManagementCatalogResourceRoleOperationOptions() GetEntitlementManagementCatalogResourceRoleOperationOptions {
	return GetEntitlementManagementCatalogResourceRoleOperationOptions{}
}

func (o GetEntitlementManagementCatalogResourceRoleOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetEntitlementManagementCatalogResourceRoleOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetEntitlementManagementCatalogResourceRoleOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetEntitlementManagementCatalogResourceRole - Get resourceRoles from identityGovernance
func (c EntitlementManagementCatalogResourceRoleClient) GetEntitlementManagementCatalogResourceRole(ctx context.Context, id stable.IdentityGovernanceEntitlementManagementCatalogIdResourceRoleId, options GetEntitlementManagementCatalogResourceRoleOperationOptions) (result GetEntitlementManagementCatalogResourceRoleOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
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

	var model stable.AccessPackageResourceRole
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}

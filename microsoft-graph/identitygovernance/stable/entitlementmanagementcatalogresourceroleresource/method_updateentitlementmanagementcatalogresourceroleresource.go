package entitlementmanagementcatalogresourceroleresource

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

type UpdateEntitlementManagementCatalogResourceRoleResourceOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateEntitlementManagementCatalogResourceRoleResourceOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultUpdateEntitlementManagementCatalogResourceRoleResourceOperationOptions() UpdateEntitlementManagementCatalogResourceRoleResourceOperationOptions {
	return UpdateEntitlementManagementCatalogResourceRoleResourceOperationOptions{}
}

func (o UpdateEntitlementManagementCatalogResourceRoleResourceOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateEntitlementManagementCatalogResourceRoleResourceOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateEntitlementManagementCatalogResourceRoleResourceOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateEntitlementManagementCatalogResourceRoleResource - Update the navigation property resource in
// identityGovernance
func (c EntitlementManagementCatalogResourceRoleResourceClient) UpdateEntitlementManagementCatalogResourceRoleResource(ctx context.Context, id stable.IdentityGovernanceEntitlementManagementCatalogIdResourceIdRoleId, input stable.AccessPackageResource, options UpdateEntitlementManagementCatalogResourceRoleResourceOperationOptions) (result UpdateEntitlementManagementCatalogResourceRoleResourceOperationResponse, err error) {
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

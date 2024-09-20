package entitlementmanagementcatalogresourceroleresourcescoperesource

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

type GetEntitlementManagementCatalogResourceRoleResourceScopeResourceOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.AccessPackageResource
}

type GetEntitlementManagementCatalogResourceRoleResourceScopeResourceOperationOptions struct {
	Expand   *odata.Expand
	Metadata *odata.Metadata
	Select   *[]string
}

func DefaultGetEntitlementManagementCatalogResourceRoleResourceScopeResourceOperationOptions() GetEntitlementManagementCatalogResourceRoleResourceScopeResourceOperationOptions {
	return GetEntitlementManagementCatalogResourceRoleResourceScopeResourceOperationOptions{}
}

func (o GetEntitlementManagementCatalogResourceRoleResourceScopeResourceOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetEntitlementManagementCatalogResourceRoleResourceScopeResourceOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetEntitlementManagementCatalogResourceRoleResourceScopeResourceOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetEntitlementManagementCatalogResourceRoleResourceScopeResource - Get resource from identityGovernance
func (c EntitlementManagementCatalogResourceRoleResourceScopeResourceClient) GetEntitlementManagementCatalogResourceRoleResourceScopeResource(ctx context.Context, id stable.IdentityGovernanceEntitlementManagementCatalogIdResourceIdRoleIdResourceScopeId, options GetEntitlementManagementCatalogResourceRoleResourceScopeResourceOperationOptions) (result GetEntitlementManagementCatalogResourceRoleResourceScopeResourceOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/resource", id.ID()),
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

	var model stable.AccessPackageResource
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}

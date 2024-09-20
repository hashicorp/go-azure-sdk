package entitlementmanagementcatalogresourcescoperesourceroleresourcescope

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

type GetEntitlementManagementCatalogResourceScopeResourceRoleCountOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]byte
}

type GetEntitlementManagementCatalogResourceScopeResourceRoleCountOperationOptions struct {
	Filter   *string
	Metadata *odata.Metadata
	Search   *string
}

func DefaultGetEntitlementManagementCatalogResourceScopeResourceRoleCountOperationOptions() GetEntitlementManagementCatalogResourceScopeResourceRoleCountOperationOptions {
	return GetEntitlementManagementCatalogResourceScopeResourceRoleCountOperationOptions{}
}

func (o GetEntitlementManagementCatalogResourceScopeResourceRoleCountOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetEntitlementManagementCatalogResourceScopeResourceRoleCountOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Filter != nil {
		out.Filter = *o.Filter
	}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.Search != nil {
		out.Search = *o.Search
	}
	return &out
}

func (o GetEntitlementManagementCatalogResourceScopeResourceRoleCountOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetEntitlementManagementCatalogResourceScopeResourceRoleCount - Get the number of the resource
func (c EntitlementManagementCatalogResourceScopeResourceRoleResourceScopeClient) GetEntitlementManagementCatalogResourceScopeResourceRoleCount(ctx context.Context, id stable.IdentityGovernanceEntitlementManagementCatalogIdResourceScopeIdResourceRoleId, options GetEntitlementManagementCatalogResourceScopeResourceRoleCountOperationOptions) (result GetEntitlementManagementCatalogResourceScopeResourceRoleCountOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "text/plain",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/resource/scopes/$count", id.ID()),
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

	var model []byte
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}

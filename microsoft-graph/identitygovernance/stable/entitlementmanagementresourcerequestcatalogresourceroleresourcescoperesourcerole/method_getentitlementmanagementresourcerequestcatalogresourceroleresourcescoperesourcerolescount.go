package entitlementmanagementresourcerequestcatalogresourceroleresourcescoperesourcerole

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

type GetEntitlementManagementResourceRequestCatalogResourceRoleResourceScopeResourceRolesCountOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]byte
}

type GetEntitlementManagementResourceRequestCatalogResourceRoleResourceScopeResourceRolesCountOperationOptions struct {
	Filter    *string
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Search    *string
}

func DefaultGetEntitlementManagementResourceRequestCatalogResourceRoleResourceScopeResourceRolesCountOperationOptions() GetEntitlementManagementResourceRequestCatalogResourceRoleResourceScopeResourceRolesCountOperationOptions {
	return GetEntitlementManagementResourceRequestCatalogResourceRoleResourceScopeResourceRolesCountOperationOptions{}
}

func (o GetEntitlementManagementResourceRequestCatalogResourceRoleResourceScopeResourceRolesCountOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetEntitlementManagementResourceRequestCatalogResourceRoleResourceScopeResourceRolesCountOperationOptions) ToOData() *odata.Query {
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

func (o GetEntitlementManagementResourceRequestCatalogResourceRoleResourceScopeResourceRolesCountOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetEntitlementManagementResourceRequestCatalogResourceRoleResourceScopeResourceRolesCount - Get the number of the
// resource
func (c EntitlementManagementResourceRequestCatalogResourceRoleResourceScopeResourceRoleClient) GetEntitlementManagementResourceRequestCatalogResourceRoleResourceScopeResourceRolesCount(ctx context.Context, id stable.IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceRoleIdResourceScopeId, options GetEntitlementManagementResourceRequestCatalogResourceRoleResourceScopeResourceRolesCountOperationOptions) (result GetEntitlementManagementResourceRequestCatalogResourceRoleResourceScopeResourceRolesCountOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "text/plain",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/resource/roles/$count", id.ID()),
		RetryFunc:     options.RetryFunc,
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

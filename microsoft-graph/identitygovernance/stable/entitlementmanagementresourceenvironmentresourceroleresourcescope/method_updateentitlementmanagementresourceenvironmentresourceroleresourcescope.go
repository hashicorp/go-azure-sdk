package entitlementmanagementresourceenvironmentresourceroleresourcescope

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateEntitlementManagementResourceEnvironmentResourceRoleResourceScopeOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateEntitlementManagementResourceEnvironmentResourceRoleResourceScopeOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdateEntitlementManagementResourceEnvironmentResourceRoleResourceScopeOperationOptions() UpdateEntitlementManagementResourceEnvironmentResourceRoleResourceScopeOperationOptions {
	return UpdateEntitlementManagementResourceEnvironmentResourceRoleResourceScopeOperationOptions{}
}

func (o UpdateEntitlementManagementResourceEnvironmentResourceRoleResourceScopeOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateEntitlementManagementResourceEnvironmentResourceRoleResourceScopeOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateEntitlementManagementResourceEnvironmentResourceRoleResourceScopeOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateEntitlementManagementResourceEnvironmentResourceRoleResourceScope - Update the navigation property scopes in
// identityGovernance
func (c EntitlementManagementResourceEnvironmentResourceRoleResourceScopeClient) UpdateEntitlementManagementResourceEnvironmentResourceRoleResourceScope(ctx context.Context, id stable.IdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdRoleIdResourceScopeId, input stable.AccessPackageResourceScope, options UpdateEntitlementManagementResourceEnvironmentResourceRoleResourceScopeOperationOptions) (result UpdateEntitlementManagementResourceEnvironmentResourceRoleResourceScopeOperationResponse, err error) {
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

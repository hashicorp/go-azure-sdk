package entitlementmanagementresourcerolescopescope

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

type UpdateEntitlementManagementResourceRoleScopeScopeOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateEntitlementManagementResourceRoleScopeScopeOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdateEntitlementManagementResourceRoleScopeScopeOperationOptions() UpdateEntitlementManagementResourceRoleScopeScopeOperationOptions {
	return UpdateEntitlementManagementResourceRoleScopeScopeOperationOptions{}
}

func (o UpdateEntitlementManagementResourceRoleScopeScopeOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateEntitlementManagementResourceRoleScopeScopeOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateEntitlementManagementResourceRoleScopeScopeOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateEntitlementManagementResourceRoleScopeScope - Update the navigation property scope in identityGovernance
func (c EntitlementManagementResourceRoleScopeScopeClient) UpdateEntitlementManagementResourceRoleScopeScope(ctx context.Context, id stable.IdentityGovernanceEntitlementManagementResourceRoleScopeId, input stable.AccessPackageResourceScope, options UpdateEntitlementManagementResourceRoleScopeScopeOperationOptions) (result UpdateEntitlementManagementResourceRoleScopeScopeOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/scope", id.ID()),
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

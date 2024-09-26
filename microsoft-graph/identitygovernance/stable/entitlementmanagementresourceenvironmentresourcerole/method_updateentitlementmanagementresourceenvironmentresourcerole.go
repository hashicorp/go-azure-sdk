package entitlementmanagementresourceenvironmentresourcerole

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateEntitlementManagementResourceEnvironmentResourceRoleOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateEntitlementManagementResourceEnvironmentResourceRoleOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdateEntitlementManagementResourceEnvironmentResourceRoleOperationOptions() UpdateEntitlementManagementResourceEnvironmentResourceRoleOperationOptions {
	return UpdateEntitlementManagementResourceEnvironmentResourceRoleOperationOptions{}
}

func (o UpdateEntitlementManagementResourceEnvironmentResourceRoleOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateEntitlementManagementResourceEnvironmentResourceRoleOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateEntitlementManagementResourceEnvironmentResourceRoleOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateEntitlementManagementResourceEnvironmentResourceRole - Update the navigation property roles in
// identityGovernance
func (c EntitlementManagementResourceEnvironmentResourceRoleClient) UpdateEntitlementManagementResourceEnvironmentResourceRole(ctx context.Context, id stable.IdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdRoleId, input stable.AccessPackageResourceRole, options UpdateEntitlementManagementResourceEnvironmentResourceRoleOperationOptions) (result UpdateEntitlementManagementResourceEnvironmentResourceRoleOperationResponse, err error) {
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

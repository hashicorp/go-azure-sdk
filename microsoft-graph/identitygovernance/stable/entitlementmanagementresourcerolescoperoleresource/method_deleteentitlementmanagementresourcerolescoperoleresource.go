package entitlementmanagementresourcerolescoperoleresource

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

type DeleteEntitlementManagementResourceRoleScopeRoleResourceOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type DeleteEntitlementManagementResourceRoleScopeRoleResourceOperationOptions struct {
	IfMatch *string
}

func DefaultDeleteEntitlementManagementResourceRoleScopeRoleResourceOperationOptions() DeleteEntitlementManagementResourceRoleScopeRoleResourceOperationOptions {
	return DeleteEntitlementManagementResourceRoleScopeRoleResourceOperationOptions{}
}

func (o DeleteEntitlementManagementResourceRoleScopeRoleResourceOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o DeleteEntitlementManagementResourceRoleScopeRoleResourceOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o DeleteEntitlementManagementResourceRoleScopeRoleResourceOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// DeleteEntitlementManagementResourceRoleScopeRoleResource - Delete navigation property resource for identityGovernance
func (c EntitlementManagementResourceRoleScopeRoleResourceClient) DeleteEntitlementManagementResourceRoleScopeRoleResource(ctx context.Context, id stable.IdentityGovernanceEntitlementManagementResourceRoleScopeId, options DeleteEntitlementManagementResourceRoleScopeRoleResourceOperationOptions) (result DeleteEntitlementManagementResourceRoleScopeRoleResourceOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodDelete,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/role/resource", id.ID()),
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

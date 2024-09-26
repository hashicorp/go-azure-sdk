package entitlementmanagementaccesspackageaccesspackageresourcerolescopeaccesspackageresourceroleaccesspackageresourceaccesspackageresourcescope

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

type DeleteEntitlementManagementAccessPackageResourceRoleScopeResourceRoleAccessPackageResourceResourceScopeOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type DeleteEntitlementManagementAccessPackageResourceRoleScopeResourceRoleAccessPackageResourceResourceScopeOperationOptions struct {
	IfMatch   *string
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultDeleteEntitlementManagementAccessPackageResourceRoleScopeResourceRoleAccessPackageResourceResourceScopeOperationOptions() DeleteEntitlementManagementAccessPackageResourceRoleScopeResourceRoleAccessPackageResourceResourceScopeOperationOptions {
	return DeleteEntitlementManagementAccessPackageResourceRoleScopeResourceRoleAccessPackageResourceResourceScopeOperationOptions{}
}

func (o DeleteEntitlementManagementAccessPackageResourceRoleScopeResourceRoleAccessPackageResourceResourceScopeOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o DeleteEntitlementManagementAccessPackageResourceRoleScopeResourceRoleAccessPackageResourceResourceScopeOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o DeleteEntitlementManagementAccessPackageResourceRoleScopeResourceRoleAccessPackageResourceResourceScopeOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// DeleteEntitlementManagementAccessPackageResourceRoleScopeResourceRoleAccessPackageResourceResourceScope - Delete
// navigation property accessPackageResourceScopes for identityGovernance
func (c EntitlementManagementAccessPackageAccessPackageResourceRoleScopeAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceScopeClient) DeleteEntitlementManagementAccessPackageResourceRoleScopeResourceRoleAccessPackageResourceResourceScope(ctx context.Context, id beta.IdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageResourceRoleScopeIdAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceScopeId, options DeleteEntitlementManagementAccessPackageResourceRoleScopeResourceRoleAccessPackageResourceResourceScopeOperationOptions) (result DeleteEntitlementManagementAccessPackageResourceRoleScopeResourceRoleAccessPackageResourceResourceScopeOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodDelete,
		OptionsObject: options,
		Path:          id.ID(),
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

	return
}

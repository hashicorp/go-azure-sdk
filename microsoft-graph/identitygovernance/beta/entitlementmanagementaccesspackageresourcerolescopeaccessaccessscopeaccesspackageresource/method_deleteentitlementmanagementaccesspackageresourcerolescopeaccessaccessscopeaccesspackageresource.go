package entitlementmanagementaccesspackageresourcerolescopeaccessaccessscopeaccesspackageresource

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

type DeleteEntitlementManagementAccessPackageResourceRoleScopeAccessAccessScopeAccessPackageResourceOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type DeleteEntitlementManagementAccessPackageResourceRoleScopeAccessAccessScopeAccessPackageResourceOperationOptions struct {
	IfMatch *string
}

func DefaultDeleteEntitlementManagementAccessPackageResourceRoleScopeAccessAccessScopeAccessPackageResourceOperationOptions() DeleteEntitlementManagementAccessPackageResourceRoleScopeAccessAccessScopeAccessPackageResourceOperationOptions {
	return DeleteEntitlementManagementAccessPackageResourceRoleScopeAccessAccessScopeAccessPackageResourceOperationOptions{}
}

func (o DeleteEntitlementManagementAccessPackageResourceRoleScopeAccessAccessScopeAccessPackageResourceOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o DeleteEntitlementManagementAccessPackageResourceRoleScopeAccessAccessScopeAccessPackageResourceOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o DeleteEntitlementManagementAccessPackageResourceRoleScopeAccessAccessScopeAccessPackageResourceOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// DeleteEntitlementManagementAccessPackageResourceRoleScopeAccessAccessScopeAccessPackageResource - Delete navigation
// property accessPackageResource for identityGovernance
func (c EntitlementManagementAccessPackageResourceRoleScopeAccessAccessScopeAccessPackageResourceClient) DeleteEntitlementManagementAccessPackageResourceRoleScopeAccessAccessScopeAccessPackageResource(ctx context.Context, id beta.IdentityGovernanceEntitlementManagementAccessPackageResourceRoleScopeIdAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceScopeId, options DeleteEntitlementManagementAccessPackageResourceRoleScopeAccessAccessScopeAccessPackageResourceOperationOptions) (result DeleteEntitlementManagementAccessPackageResourceRoleScopeAccessAccessScopeAccessPackageResourceOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodDelete,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/accessPackageResource", id.ID()),
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

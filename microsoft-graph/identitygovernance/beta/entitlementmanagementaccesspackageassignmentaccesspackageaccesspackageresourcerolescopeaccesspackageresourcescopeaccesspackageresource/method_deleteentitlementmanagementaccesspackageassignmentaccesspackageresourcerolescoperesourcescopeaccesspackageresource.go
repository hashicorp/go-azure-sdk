package entitlementmanagementaccesspackageassignmentaccesspackageaccesspackageresourcerolescopeaccesspackageresourcescopeaccesspackageresource

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

type DeleteEntitlementManagementAccessPackageAssignmentAccessPackageResourceRoleScopeResourceScopeAccessPackageResourceOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type DeleteEntitlementManagementAccessPackageAssignmentAccessPackageResourceRoleScopeResourceScopeAccessPackageResourceOperationOptions struct {
	IfMatch *string
}

func DefaultDeleteEntitlementManagementAccessPackageAssignmentAccessPackageResourceRoleScopeResourceScopeAccessPackageResourceOperationOptions() DeleteEntitlementManagementAccessPackageAssignmentAccessPackageResourceRoleScopeResourceScopeAccessPackageResourceOperationOptions {
	return DeleteEntitlementManagementAccessPackageAssignmentAccessPackageResourceRoleScopeResourceScopeAccessPackageResourceOperationOptions{}
}

func (o DeleteEntitlementManagementAccessPackageAssignmentAccessPackageResourceRoleScopeResourceScopeAccessPackageResourceOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o DeleteEntitlementManagementAccessPackageAssignmentAccessPackageResourceRoleScopeResourceScopeAccessPackageResourceOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o DeleteEntitlementManagementAccessPackageAssignmentAccessPackageResourceRoleScopeResourceScopeAccessPackageResourceOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// DeleteEntitlementManagementAccessPackageAssignmentAccessPackageResourceRoleScopeResourceScopeAccessPackageResource -
// Delete navigation property accessPackageResource for identityGovernance
func (c EntitlementManagementAccessPackageAssignmentAccessPackageAccessPackageResourceRoleScopeAccessPackageResourceScopeAccessPackageResourceClient) DeleteEntitlementManagementAccessPackageAssignmentAccessPackageResourceRoleScopeResourceScopeAccessPackageResource(ctx context.Context, id beta.IdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAccessPackageResourceRoleScopeId, options DeleteEntitlementManagementAccessPackageAssignmentAccessPackageResourceRoleScopeResourceScopeAccessPackageResourceOperationOptions) (result DeleteEntitlementManagementAccessPackageAssignmentAccessPackageResourceRoleScopeResourceScopeAccessPackageResourceOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodDelete,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/accessPackageResourceScope/accessPackageResource", id.ID()),
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

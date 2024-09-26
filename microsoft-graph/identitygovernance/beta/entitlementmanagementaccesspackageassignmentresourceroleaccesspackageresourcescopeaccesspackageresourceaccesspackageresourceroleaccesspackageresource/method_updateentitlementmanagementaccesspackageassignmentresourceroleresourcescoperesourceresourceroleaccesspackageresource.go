package entitlementmanagementaccesspackageassignmentresourceroleaccesspackageresourcescopeaccesspackageresourceaccesspackageresourceroleaccesspackageresource

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

type UpdateEntitlementManagementAccessPackageAssignmentResourceRoleResourceScopeResourceResourceRoleAccessPackageResourceOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateEntitlementManagementAccessPackageAssignmentResourceRoleResourceScopeResourceResourceRoleAccessPackageResourceOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdateEntitlementManagementAccessPackageAssignmentResourceRoleResourceScopeResourceResourceRoleAccessPackageResourceOperationOptions() UpdateEntitlementManagementAccessPackageAssignmentResourceRoleResourceScopeResourceResourceRoleAccessPackageResourceOperationOptions {
	return UpdateEntitlementManagementAccessPackageAssignmentResourceRoleResourceScopeResourceResourceRoleAccessPackageResourceOperationOptions{}
}

func (o UpdateEntitlementManagementAccessPackageAssignmentResourceRoleResourceScopeResourceResourceRoleAccessPackageResourceOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateEntitlementManagementAccessPackageAssignmentResourceRoleResourceScopeResourceResourceRoleAccessPackageResourceOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateEntitlementManagementAccessPackageAssignmentResourceRoleResourceScopeResourceResourceRoleAccessPackageResourceOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateEntitlementManagementAccessPackageAssignmentResourceRoleResourceScopeResourceResourceRoleAccessPackageResource
// - Update the navigation property accessPackageResource in identityGovernance
func (c EntitlementManagementAccessPackageAssignmentResourceRoleAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRoleAccessPackageResourceClient) UpdateEntitlementManagementAccessPackageAssignmentResourceRoleResourceScopeResourceResourceRoleAccessPackageResource(ctx context.Context, id beta.IdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRoleIdAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRoleId, input beta.AccessPackageResource, options UpdateEntitlementManagementAccessPackageAssignmentResourceRoleResourceScopeResourceResourceRoleAccessPackageResourceOperationOptions) (result UpdateEntitlementManagementAccessPackageAssignmentResourceRoleResourceScopeResourceResourceRoleAccessPackageResourceOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/accessPackageResource", id.ID()),
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

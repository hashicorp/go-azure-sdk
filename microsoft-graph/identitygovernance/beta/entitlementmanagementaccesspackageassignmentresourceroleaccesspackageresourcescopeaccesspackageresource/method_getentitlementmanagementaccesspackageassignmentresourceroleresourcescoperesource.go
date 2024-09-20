package entitlementmanagementaccesspackageassignmentresourceroleaccesspackageresourcescopeaccesspackageresource

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

type GetEntitlementManagementAccessPackageAssignmentResourceRoleResourceScopeResourceOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.AccessPackageResource
}

type GetEntitlementManagementAccessPackageAssignmentResourceRoleResourceScopeResourceOperationOptions struct {
	Expand   *odata.Expand
	Metadata *odata.Metadata
	Select   *[]string
}

func DefaultGetEntitlementManagementAccessPackageAssignmentResourceRoleResourceScopeResourceOperationOptions() GetEntitlementManagementAccessPackageAssignmentResourceRoleResourceScopeResourceOperationOptions {
	return GetEntitlementManagementAccessPackageAssignmentResourceRoleResourceScopeResourceOperationOptions{}
}

func (o GetEntitlementManagementAccessPackageAssignmentResourceRoleResourceScopeResourceOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetEntitlementManagementAccessPackageAssignmentResourceRoleResourceScopeResourceOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetEntitlementManagementAccessPackageAssignmentResourceRoleResourceScopeResourceOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetEntitlementManagementAccessPackageAssignmentResourceRoleResourceScopeResource - Get accessPackageResource from
// identityGovernance
func (c EntitlementManagementAccessPackageAssignmentResourceRoleAccessPackageResourceScopeAccessPackageResourceClient) GetEntitlementManagementAccessPackageAssignmentResourceRoleResourceScopeResource(ctx context.Context, id beta.IdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRoleId, options GetEntitlementManagementAccessPackageAssignmentResourceRoleResourceScopeResourceOperationOptions) (result GetEntitlementManagementAccessPackageAssignmentResourceRoleResourceScopeResourceOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
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

	var model beta.AccessPackageResource
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}

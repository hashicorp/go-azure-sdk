package entitlementmanagementaccesspackageassignmentaccesspackageassignmentresourceroleaccesspackageresourceroleaccesspackageresourceaccesspackageresourcerole

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

type CreateEntitlementManagementAccessPackageAssignmentResourceRoleResourceAccessPackageResourceRoleOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.AccessPackageResourceRole
}

type CreateEntitlementManagementAccessPackageAssignmentResourceRoleResourceAccessPackageResourceRoleOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultCreateEntitlementManagementAccessPackageAssignmentResourceRoleResourceAccessPackageResourceRoleOperationOptions() CreateEntitlementManagementAccessPackageAssignmentResourceRoleResourceAccessPackageResourceRoleOperationOptions {
	return CreateEntitlementManagementAccessPackageAssignmentResourceRoleResourceAccessPackageResourceRoleOperationOptions{}
}

func (o CreateEntitlementManagementAccessPackageAssignmentResourceRoleResourceAccessPackageResourceRoleOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateEntitlementManagementAccessPackageAssignmentResourceRoleResourceAccessPackageResourceRoleOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateEntitlementManagementAccessPackageAssignmentResourceRoleResourceAccessPackageResourceRoleOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateEntitlementManagementAccessPackageAssignmentResourceRoleResourceAccessPackageResourceRole - Create new
// navigation property to accessPackageResourceRoles for identityGovernance
func (c EntitlementManagementAccessPackageAssignmentAccessPackageAssignmentResourceRoleAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceRoleClient) CreateEntitlementManagementAccessPackageAssignmentResourceRoleResourceAccessPackageResourceRole(ctx context.Context, id beta.IdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAssignmentResourceRoleId, input beta.AccessPackageResourceRole, options CreateEntitlementManagementAccessPackageAssignmentResourceRoleResourceAccessPackageResourceRoleOperationOptions) (result CreateEntitlementManagementAccessPackageAssignmentResourceRoleResourceAccessPackageResourceRoleOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/accessPackageResourceRole/accessPackageResource/accessPackageResourceRoles", id.ID()),
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

	var model beta.AccessPackageResourceRole
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}

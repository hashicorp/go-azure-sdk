package entitlementmanagementaccesspackageassignmentaccesspackageassignmentresourceroleaccesspackageresourcescopeaccesspackageresourceaccesspackageresourcerole

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

type CreateEntitlementManagementAccessPackageAssignmentResourceRoleResourceScopeResourceResourceRoleOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.AccessPackageResourceRole
}

type CreateEntitlementManagementAccessPackageAssignmentResourceRoleResourceScopeResourceResourceRoleOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateEntitlementManagementAccessPackageAssignmentResourceRoleResourceScopeResourceResourceRoleOperationOptions() CreateEntitlementManagementAccessPackageAssignmentResourceRoleResourceScopeResourceResourceRoleOperationOptions {
	return CreateEntitlementManagementAccessPackageAssignmentResourceRoleResourceScopeResourceResourceRoleOperationOptions{}
}

func (o CreateEntitlementManagementAccessPackageAssignmentResourceRoleResourceScopeResourceResourceRoleOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateEntitlementManagementAccessPackageAssignmentResourceRoleResourceScopeResourceResourceRoleOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateEntitlementManagementAccessPackageAssignmentResourceRoleResourceScopeResourceResourceRoleOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateEntitlementManagementAccessPackageAssignmentResourceRoleResourceScopeResourceResourceRole - Create new
// navigation property to accessPackageResourceRoles for identityGovernance
func (c EntitlementManagementAccessPackageAssignmentAccessPackageAssignmentResourceRoleAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRoleClient) CreateEntitlementManagementAccessPackageAssignmentResourceRoleResourceScopeResourceResourceRole(ctx context.Context, id beta.IdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAssignmentResourceRoleId, input beta.AccessPackageResourceRole, options CreateEntitlementManagementAccessPackageAssignmentResourceRoleResourceScopeResourceResourceRoleOperationOptions) (result CreateEntitlementManagementAccessPackageAssignmentResourceRoleResourceScopeResourceResourceRoleOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusCreated,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/accessPackageResourceScope/accessPackageResource/accessPackageResourceRoles", id.ID()),
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

	var model beta.AccessPackageResourceRole
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}

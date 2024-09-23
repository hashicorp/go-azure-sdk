package entitlementmanagementroleassignmentapproval

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateEntitlementManagementRoleAssignmentApprovalOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateEntitlementManagementRoleAssignmentApprovalOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdateEntitlementManagementRoleAssignmentApprovalOperationOptions() UpdateEntitlementManagementRoleAssignmentApprovalOperationOptions {
	return UpdateEntitlementManagementRoleAssignmentApprovalOperationOptions{}
}

func (o UpdateEntitlementManagementRoleAssignmentApprovalOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateEntitlementManagementRoleAssignmentApprovalOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateEntitlementManagementRoleAssignmentApprovalOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateEntitlementManagementRoleAssignmentApproval - Update the navigation property roleAssignmentApprovals in
// roleManagement
func (c EntitlementManagementRoleAssignmentApprovalClient) UpdateEntitlementManagementRoleAssignmentApproval(ctx context.Context, id beta.RoleManagementEntitlementManagementRoleAssignmentApprovalId, input beta.Approval, options UpdateEntitlementManagementRoleAssignmentApprovalOperationOptions) (result UpdateEntitlementManagementRoleAssignmentApprovalOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
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

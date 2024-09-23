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

type CreateEntitlementManagementRoleAssignmentApprovalOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.Approval
}

type CreateEntitlementManagementRoleAssignmentApprovalOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateEntitlementManagementRoleAssignmentApprovalOperationOptions() CreateEntitlementManagementRoleAssignmentApprovalOperationOptions {
	return CreateEntitlementManagementRoleAssignmentApprovalOperationOptions{}
}

func (o CreateEntitlementManagementRoleAssignmentApprovalOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateEntitlementManagementRoleAssignmentApprovalOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateEntitlementManagementRoleAssignmentApprovalOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateEntitlementManagementRoleAssignmentApproval - Create new navigation property to roleAssignmentApprovals for
// roleManagement
func (c EntitlementManagementRoleAssignmentApprovalClient) CreateEntitlementManagementRoleAssignmentApproval(ctx context.Context, input beta.Approval, options CreateEntitlementManagementRoleAssignmentApprovalOperationOptions) (result CreateEntitlementManagementRoleAssignmentApprovalOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          "/roleManagement/entitlementManagement/roleAssignmentApprovals",
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

	var model beta.Approval
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}

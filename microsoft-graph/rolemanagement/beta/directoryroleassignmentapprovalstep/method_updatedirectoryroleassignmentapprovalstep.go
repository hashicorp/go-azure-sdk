package directoryroleassignmentapprovalstep

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateDirectoryRoleAssignmentApprovalStepOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateDirectoryRoleAssignmentApprovalStepOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdateDirectoryRoleAssignmentApprovalStepOperationOptions() UpdateDirectoryRoleAssignmentApprovalStepOperationOptions {
	return UpdateDirectoryRoleAssignmentApprovalStepOperationOptions{}
}

func (o UpdateDirectoryRoleAssignmentApprovalStepOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateDirectoryRoleAssignmentApprovalStepOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateDirectoryRoleAssignmentApprovalStepOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateDirectoryRoleAssignmentApprovalStep - Update the navigation property steps in roleManagement
func (c DirectoryRoleAssignmentApprovalStepClient) UpdateDirectoryRoleAssignmentApprovalStep(ctx context.Context, id beta.RoleManagementDirectoryRoleAssignmentApprovalIdStepId, input beta.ApprovalStep, options UpdateDirectoryRoleAssignmentApprovalStepOperationOptions) (result UpdateDirectoryRoleAssignmentApprovalStepOperationResponse, err error) {
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

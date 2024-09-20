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

type GetDirectoryRoleAssignmentApprovalStepOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.ApprovalStep
}

type GetDirectoryRoleAssignmentApprovalStepOperationOptions struct {
	Expand   *odata.Expand
	Metadata *odata.Metadata
	Select   *[]string
}

func DefaultGetDirectoryRoleAssignmentApprovalStepOperationOptions() GetDirectoryRoleAssignmentApprovalStepOperationOptions {
	return GetDirectoryRoleAssignmentApprovalStepOperationOptions{}
}

func (o GetDirectoryRoleAssignmentApprovalStepOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetDirectoryRoleAssignmentApprovalStepOperationOptions) ToOData() *odata.Query {
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

func (o GetDirectoryRoleAssignmentApprovalStepOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetDirectoryRoleAssignmentApprovalStep - Get steps from roleManagement. Used to represent the decision associated
// with a single step in the approval process configured in approvalStage.
func (c DirectoryRoleAssignmentApprovalStepClient) GetDirectoryRoleAssignmentApprovalStep(ctx context.Context, id beta.RoleManagementDirectoryRoleAssignmentApprovalIdStepId, options GetDirectoryRoleAssignmentApprovalStepOperationOptions) (result GetDirectoryRoleAssignmentApprovalStepOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          id.ID(),
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

	var model beta.ApprovalStep
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}

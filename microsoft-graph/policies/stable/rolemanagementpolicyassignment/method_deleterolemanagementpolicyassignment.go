package rolemanagementpolicyassignment

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeleteRoleManagementPolicyAssignmentOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type DeleteRoleManagementPolicyAssignmentOperationOptions struct {
	IfMatch *string
}

func DefaultDeleteRoleManagementPolicyAssignmentOperationOptions() DeleteRoleManagementPolicyAssignmentOperationOptions {
	return DeleteRoleManagementPolicyAssignmentOperationOptions{}
}

func (o DeleteRoleManagementPolicyAssignmentOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o DeleteRoleManagementPolicyAssignmentOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o DeleteRoleManagementPolicyAssignmentOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// DeleteRoleManagementPolicyAssignment - Delete navigation property roleManagementPolicyAssignments for policies
func (c RoleManagementPolicyAssignmentClient) DeleteRoleManagementPolicyAssignment(ctx context.Context, id stable.PolicyRoleManagementPolicyAssignmentId, options DeleteRoleManagementPolicyAssignmentOperationOptions) (result DeleteRoleManagementPolicyAssignmentOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodDelete,
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

	return
}

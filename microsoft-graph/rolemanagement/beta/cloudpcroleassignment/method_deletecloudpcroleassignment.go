package cloudpcroleassignment

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

type DeleteCloudPCRoleAssignmentOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type DeleteCloudPCRoleAssignmentOperationOptions struct {
	IfMatch *string
}

func DefaultDeleteCloudPCRoleAssignmentOperationOptions() DeleteCloudPCRoleAssignmentOperationOptions {
	return DeleteCloudPCRoleAssignmentOperationOptions{}
}

func (o DeleteCloudPCRoleAssignmentOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o DeleteCloudPCRoleAssignmentOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o DeleteCloudPCRoleAssignmentOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// DeleteCloudPCRoleAssignment - Delete unifiedRoleAssignmentMultiple. Delete a unifiedRoleAssignmentMultiple object of
// an RBAC provider. This is applicable for a RBAC application that supports multiple principals and scopes. The
// following RBAC providers are currently supported: - Cloud PC - device management (Intune)
func (c CloudPCRoleAssignmentClient) DeleteCloudPCRoleAssignment(ctx context.Context, id beta.RoleManagementCloudPCRoleAssignmentId, options DeleteCloudPCRoleAssignmentOperationOptions) (result DeleteCloudPCRoleAssignmentOperationResponse, err error) {
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

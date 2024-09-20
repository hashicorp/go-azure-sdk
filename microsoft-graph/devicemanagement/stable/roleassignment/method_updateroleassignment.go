package roleassignment

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateRoleAssignmentOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateRoleAssignmentOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultUpdateRoleAssignmentOperationOptions() UpdateRoleAssignmentOperationOptions {
	return UpdateRoleAssignmentOperationOptions{}
}

func (o UpdateRoleAssignmentOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateRoleAssignmentOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateRoleAssignmentOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateRoleAssignment - Update deviceAndAppManagementRoleAssignment. Update the properties of a
// deviceAndAppManagementRoleAssignment object.
func (c RoleAssignmentClient) UpdateRoleAssignment(ctx context.Context, id stable.DeviceManagementRoleAssignmentId, input stable.DeviceAndAppManagementRoleAssignment, options UpdateRoleAssignmentOperationOptions) (result UpdateRoleAssignmentOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          id.ID(),
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

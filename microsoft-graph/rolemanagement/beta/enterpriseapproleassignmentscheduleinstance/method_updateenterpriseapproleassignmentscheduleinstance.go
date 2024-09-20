package enterpriseapproleassignmentscheduleinstance

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateEnterpriseAppRoleAssignmentScheduleInstanceOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateEnterpriseAppRoleAssignmentScheduleInstanceOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultUpdateEnterpriseAppRoleAssignmentScheduleInstanceOperationOptions() UpdateEnterpriseAppRoleAssignmentScheduleInstanceOperationOptions {
	return UpdateEnterpriseAppRoleAssignmentScheduleInstanceOperationOptions{}
}

func (o UpdateEnterpriseAppRoleAssignmentScheduleInstanceOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateEnterpriseAppRoleAssignmentScheduleInstanceOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateEnterpriseAppRoleAssignmentScheduleInstanceOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateEnterpriseAppRoleAssignmentScheduleInstance - Update the navigation property roleAssignmentScheduleInstances in
// roleManagement
func (c EnterpriseAppRoleAssignmentScheduleInstanceClient) UpdateEnterpriseAppRoleAssignmentScheduleInstance(ctx context.Context, id beta.RoleManagementEnterpriseAppIdRoleAssignmentScheduleInstanceId, input beta.UnifiedRoleAssignmentScheduleInstance, options UpdateEnterpriseAppRoleAssignmentScheduleInstanceOperationOptions) (result UpdateEnterpriseAppRoleAssignmentScheduleInstanceOperationResponse, err error) {
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

package enterpriseapproleassignmentschedulerequest

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateEnterpriseAppRoleAssignmentScheduleRequestOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateEnterpriseAppRoleAssignmentScheduleRequestOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdateEnterpriseAppRoleAssignmentScheduleRequestOperationOptions() UpdateEnterpriseAppRoleAssignmentScheduleRequestOperationOptions {
	return UpdateEnterpriseAppRoleAssignmentScheduleRequestOperationOptions{}
}

func (o UpdateEnterpriseAppRoleAssignmentScheduleRequestOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateEnterpriseAppRoleAssignmentScheduleRequestOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateEnterpriseAppRoleAssignmentScheduleRequestOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateEnterpriseAppRoleAssignmentScheduleRequest - Update the navigation property roleAssignmentScheduleRequests in
// roleManagement
func (c EnterpriseAppRoleAssignmentScheduleRequestClient) UpdateEnterpriseAppRoleAssignmentScheduleRequest(ctx context.Context, id beta.RoleManagementEnterpriseAppIdRoleAssignmentScheduleRequestId, input beta.UnifiedRoleAssignmentScheduleRequest, options UpdateEnterpriseAppRoleAssignmentScheduleRequestOperationOptions) (result UpdateEnterpriseAppRoleAssignmentScheduleRequestOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusNoContent,
			http.StatusOK,
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

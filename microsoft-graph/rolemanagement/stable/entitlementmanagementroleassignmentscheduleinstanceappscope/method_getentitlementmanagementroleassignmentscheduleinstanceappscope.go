package entitlementmanagementroleassignmentscheduleinstanceappscope

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

type GetEntitlementManagementRoleAssignmentScheduleInstanceAppScopeOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.AppScope
}

type GetEntitlementManagementRoleAssignmentScheduleInstanceAppScopeOperationOptions struct {
	Expand    *odata.Expand
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Select    *[]string
}

func DefaultGetEntitlementManagementRoleAssignmentScheduleInstanceAppScopeOperationOptions() GetEntitlementManagementRoleAssignmentScheduleInstanceAppScopeOperationOptions {
	return GetEntitlementManagementRoleAssignmentScheduleInstanceAppScopeOperationOptions{}
}

func (o GetEntitlementManagementRoleAssignmentScheduleInstanceAppScopeOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetEntitlementManagementRoleAssignmentScheduleInstanceAppScopeOperationOptions) ToOData() *odata.Query {
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

func (o GetEntitlementManagementRoleAssignmentScheduleInstanceAppScopeOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetEntitlementManagementRoleAssignmentScheduleInstanceAppScope - Get appScope from roleManagement. Read-only property
// with details of the app-specific scope when the assignment or role eligibility is scoped to an app. Nullable.
func (c EntitlementManagementRoleAssignmentScheduleInstanceAppScopeClient) GetEntitlementManagementRoleAssignmentScheduleInstanceAppScope(ctx context.Context, id stable.RoleManagementEntitlementManagementRoleAssignmentScheduleInstanceId, options GetEntitlementManagementRoleAssignmentScheduleInstanceAppScopeOperationOptions) (result GetEntitlementManagementRoleAssignmentScheduleInstanceAppScopeOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/appScope", id.ID()),
		RetryFunc:     options.RetryFunc,
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

	var model stable.AppScope
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}

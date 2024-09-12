package entitlementmanagementroleassignmentscheduleappscope

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

type GetEntitlementManagementRoleAssignmentScheduleAppScopeOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.AppScope
}

type GetEntitlementManagementRoleAssignmentScheduleAppScopeOperationOptions struct {
	Expand *odata.Expand
	Select *[]string
}

func DefaultGetEntitlementManagementRoleAssignmentScheduleAppScopeOperationOptions() GetEntitlementManagementRoleAssignmentScheduleAppScopeOperationOptions {
	return GetEntitlementManagementRoleAssignmentScheduleAppScopeOperationOptions{}
}

func (o GetEntitlementManagementRoleAssignmentScheduleAppScopeOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetEntitlementManagementRoleAssignmentScheduleAppScopeOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetEntitlementManagementRoleAssignmentScheduleAppScopeOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetEntitlementManagementRoleAssignmentScheduleAppScope - Get appScope from roleManagement. Read-only property with
// details of the app-specific scope when the role eligibility or assignment is scoped to an app. Nullable.
func (c EntitlementManagementRoleAssignmentScheduleAppScopeClient) GetEntitlementManagementRoleAssignmentScheduleAppScope(ctx context.Context, id stable.RoleManagementEntitlementManagementRoleAssignmentScheduleId, options GetEntitlementManagementRoleAssignmentScheduleAppScopeOperationOptions) (result GetEntitlementManagementRoleAssignmentScheduleAppScopeOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/appScope", id.ID()),
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

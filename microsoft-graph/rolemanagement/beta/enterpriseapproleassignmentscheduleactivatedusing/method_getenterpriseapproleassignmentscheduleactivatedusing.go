package enterpriseapproleassignmentscheduleactivatedusing

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

type GetEnterpriseAppRoleAssignmentScheduleActivatedUsingOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.UnifiedRoleEligibilitySchedule
}

type GetEnterpriseAppRoleAssignmentScheduleActivatedUsingOperationOptions struct {
	Expand   *odata.Expand
	Metadata *odata.Metadata
	Select   *[]string
}

func DefaultGetEnterpriseAppRoleAssignmentScheduleActivatedUsingOperationOptions() GetEnterpriseAppRoleAssignmentScheduleActivatedUsingOperationOptions {
	return GetEnterpriseAppRoleAssignmentScheduleActivatedUsingOperationOptions{}
}

func (o GetEnterpriseAppRoleAssignmentScheduleActivatedUsingOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetEnterpriseAppRoleAssignmentScheduleActivatedUsingOperationOptions) ToOData() *odata.Query {
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

func (o GetEnterpriseAppRoleAssignmentScheduleActivatedUsingOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetEnterpriseAppRoleAssignmentScheduleActivatedUsing - Get activatedUsing from roleManagement. If the request is from
// an eligible administrator to activate a role, this parameter shows the related eligible assignment for that
// activation. Otherwise, it's null. Supports $expand.
func (c EnterpriseAppRoleAssignmentScheduleActivatedUsingClient) GetEnterpriseAppRoleAssignmentScheduleActivatedUsing(ctx context.Context, id beta.RoleManagementEnterpriseAppIdRoleAssignmentScheduleId, options GetEnterpriseAppRoleAssignmentScheduleActivatedUsingOperationOptions) (result GetEnterpriseAppRoleAssignmentScheduleActivatedUsingOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/activatedUsing", id.ID()),
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

	var model beta.UnifiedRoleEligibilitySchedule
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}

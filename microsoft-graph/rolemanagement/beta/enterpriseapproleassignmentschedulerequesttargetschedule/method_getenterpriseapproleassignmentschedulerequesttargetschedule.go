package enterpriseapproleassignmentschedulerequesttargetschedule

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

type GetEnterpriseAppRoleAssignmentScheduleRequestTargetScheduleOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.UnifiedRoleAssignmentSchedule
}

type GetEnterpriseAppRoleAssignmentScheduleRequestTargetScheduleOperationOptions struct {
	Expand *odata.Expand
	Select *[]string
}

func DefaultGetEnterpriseAppRoleAssignmentScheduleRequestTargetScheduleOperationOptions() GetEnterpriseAppRoleAssignmentScheduleRequestTargetScheduleOperationOptions {
	return GetEnterpriseAppRoleAssignmentScheduleRequestTargetScheduleOperationOptions{}
}

func (o GetEnterpriseAppRoleAssignmentScheduleRequestTargetScheduleOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetEnterpriseAppRoleAssignmentScheduleRequestTargetScheduleOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetEnterpriseAppRoleAssignmentScheduleRequestTargetScheduleOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetEnterpriseAppRoleAssignmentScheduleRequestTargetSchedule - Get targetSchedule from roleManagement. The schedule
// for an eligible role assignment that is referenced through the targetScheduleId property. Supports $expand and
// $select nested in $expand.
func (c EnterpriseAppRoleAssignmentScheduleRequestTargetScheduleClient) GetEnterpriseAppRoleAssignmentScheduleRequestTargetSchedule(ctx context.Context, id beta.RoleManagementEnterpriseAppIdRoleAssignmentScheduleRequestId, options GetEnterpriseAppRoleAssignmentScheduleRequestTargetScheduleOperationOptions) (result GetEnterpriseAppRoleAssignmentScheduleRequestTargetScheduleOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/targetSchedule", id.ID()),
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

	var model beta.UnifiedRoleAssignmentSchedule
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}

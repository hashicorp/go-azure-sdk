package entitlementmanagementroleassignmentschedulerequesttargetschedule

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetEntitlementManagementRoleAssignmentScheduleRequestTargetScheduleOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.UnifiedRoleAssignmentSchedule
}

type GetEntitlementManagementRoleAssignmentScheduleRequestTargetScheduleOperationOptions struct {
	Expand    *odata.Expand
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Select    *[]string
}

func DefaultGetEntitlementManagementRoleAssignmentScheduleRequestTargetScheduleOperationOptions() GetEntitlementManagementRoleAssignmentScheduleRequestTargetScheduleOperationOptions {
	return GetEntitlementManagementRoleAssignmentScheduleRequestTargetScheduleOperationOptions{}
}

func (o GetEntitlementManagementRoleAssignmentScheduleRequestTargetScheduleOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetEntitlementManagementRoleAssignmentScheduleRequestTargetScheduleOperationOptions) ToOData() *odata.Query {
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

func (o GetEntitlementManagementRoleAssignmentScheduleRequestTargetScheduleOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetEntitlementManagementRoleAssignmentScheduleRequestTargetSchedule - Get targetSchedule from roleManagement. The
// schedule for an eligible role assignment that is referenced through the targetScheduleId property. Supports $expand
// and $select nested in $expand.
func (c EntitlementManagementRoleAssignmentScheduleRequestTargetScheduleClient) GetEntitlementManagementRoleAssignmentScheduleRequestTargetSchedule(ctx context.Context, id stable.RoleManagementEntitlementManagementRoleAssignmentScheduleRequestId, options GetEntitlementManagementRoleAssignmentScheduleRequestTargetScheduleOperationOptions) (result GetEntitlementManagementRoleAssignmentScheduleRequestTargetScheduleOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/targetSchedule", id.ID()),
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

	var model stable.UnifiedRoleAssignmentSchedule
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}

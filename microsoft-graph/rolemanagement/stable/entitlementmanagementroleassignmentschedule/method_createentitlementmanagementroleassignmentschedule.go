package entitlementmanagementroleassignmentschedule

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateEntitlementManagementRoleAssignmentScheduleOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.UnifiedRoleAssignmentSchedule
}

type CreateEntitlementManagementRoleAssignmentScheduleOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateEntitlementManagementRoleAssignmentScheduleOperationOptions() CreateEntitlementManagementRoleAssignmentScheduleOperationOptions {
	return CreateEntitlementManagementRoleAssignmentScheduleOperationOptions{}
}

func (o CreateEntitlementManagementRoleAssignmentScheduleOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateEntitlementManagementRoleAssignmentScheduleOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateEntitlementManagementRoleAssignmentScheduleOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateEntitlementManagementRoleAssignmentSchedule - Create new navigation property to roleAssignmentSchedules for
// roleManagement
func (c EntitlementManagementRoleAssignmentScheduleClient) CreateEntitlementManagementRoleAssignmentSchedule(ctx context.Context, input stable.UnifiedRoleAssignmentSchedule, options CreateEntitlementManagementRoleAssignmentScheduleOperationOptions) (result CreateEntitlementManagementRoleAssignmentScheduleOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusCreated,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          "/roleManagement/entitlementManagement/roleAssignmentSchedules",
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

	var model stable.UnifiedRoleAssignmentSchedule
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}

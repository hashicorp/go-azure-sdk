package entitlementmanagementroleeligibilityschedule

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateEntitlementManagementRoleEligibilityScheduleOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateEntitlementManagementRoleEligibilityScheduleOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultUpdateEntitlementManagementRoleEligibilityScheduleOperationOptions() UpdateEntitlementManagementRoleEligibilityScheduleOperationOptions {
	return UpdateEntitlementManagementRoleEligibilityScheduleOperationOptions{}
}

func (o UpdateEntitlementManagementRoleEligibilityScheduleOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateEntitlementManagementRoleEligibilityScheduleOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateEntitlementManagementRoleEligibilityScheduleOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateEntitlementManagementRoleEligibilitySchedule - Update the navigation property roleEligibilitySchedules in
// roleManagement
func (c EntitlementManagementRoleEligibilityScheduleClient) UpdateEntitlementManagementRoleEligibilitySchedule(ctx context.Context, id stable.RoleManagementEntitlementManagementRoleEligibilityScheduleId, input stable.UnifiedRoleEligibilitySchedule, options UpdateEntitlementManagementRoleEligibilityScheduleOperationOptions) (result UpdateEntitlementManagementRoleEligibilityScheduleOperationResponse, err error) {
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

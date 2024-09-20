package entitlementmanagementroleeligibilityscheduleinstance

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateEntitlementManagementRoleEligibilityScheduleInstanceOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateEntitlementManagementRoleEligibilityScheduleInstanceOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultUpdateEntitlementManagementRoleEligibilityScheduleInstanceOperationOptions() UpdateEntitlementManagementRoleEligibilityScheduleInstanceOperationOptions {
	return UpdateEntitlementManagementRoleEligibilityScheduleInstanceOperationOptions{}
}

func (o UpdateEntitlementManagementRoleEligibilityScheduleInstanceOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateEntitlementManagementRoleEligibilityScheduleInstanceOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateEntitlementManagementRoleEligibilityScheduleInstanceOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateEntitlementManagementRoleEligibilityScheduleInstance - Update the navigation property
// roleEligibilityScheduleInstances in roleManagement
func (c EntitlementManagementRoleEligibilityScheduleInstanceClient) UpdateEntitlementManagementRoleEligibilityScheduleInstance(ctx context.Context, id beta.RoleManagementEntitlementManagementRoleEligibilityScheduleInstanceId, input beta.UnifiedRoleEligibilityScheduleInstance, options UpdateEntitlementManagementRoleEligibilityScheduleInstanceOperationOptions) (result UpdateEntitlementManagementRoleEligibilityScheduleInstanceOperationResponse, err error) {
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

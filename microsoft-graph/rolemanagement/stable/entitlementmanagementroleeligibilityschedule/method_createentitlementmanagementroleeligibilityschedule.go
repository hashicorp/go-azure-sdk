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

type CreateEntitlementManagementRoleEligibilityScheduleOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.UnifiedRoleEligibilitySchedule
}

type CreateEntitlementManagementRoleEligibilityScheduleOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateEntitlementManagementRoleEligibilityScheduleOperationOptions() CreateEntitlementManagementRoleEligibilityScheduleOperationOptions {
	return CreateEntitlementManagementRoleEligibilityScheduleOperationOptions{}
}

func (o CreateEntitlementManagementRoleEligibilityScheduleOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateEntitlementManagementRoleEligibilityScheduleOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateEntitlementManagementRoleEligibilityScheduleOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateEntitlementManagementRoleEligibilitySchedule - Create new navigation property to roleEligibilitySchedules for
// roleManagement
func (c EntitlementManagementRoleEligibilityScheduleClient) CreateEntitlementManagementRoleEligibilitySchedule(ctx context.Context, input stable.UnifiedRoleEligibilitySchedule, options CreateEntitlementManagementRoleEligibilityScheduleOperationOptions) (result CreateEntitlementManagementRoleEligibilityScheduleOperationResponse, err error) {
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
		Path:          "/roleManagement/entitlementManagement/roleEligibilitySchedules",
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

	var model stable.UnifiedRoleEligibilitySchedule
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}

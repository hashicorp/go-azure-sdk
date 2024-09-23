package entitlementmanagementroleeligibilityschedulerequest

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetEntitlementManagementRoleEligibilityScheduleRequestOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.UnifiedRoleEligibilityScheduleRequest
}

type GetEntitlementManagementRoleEligibilityScheduleRequestOperationOptions struct {
	Expand    *odata.Expand
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Select    *[]string
}

func DefaultGetEntitlementManagementRoleEligibilityScheduleRequestOperationOptions() GetEntitlementManagementRoleEligibilityScheduleRequestOperationOptions {
	return GetEntitlementManagementRoleEligibilityScheduleRequestOperationOptions{}
}

func (o GetEntitlementManagementRoleEligibilityScheduleRequestOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetEntitlementManagementRoleEligibilityScheduleRequestOperationOptions) ToOData() *odata.Query {
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

func (o GetEntitlementManagementRoleEligibilityScheduleRequestOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetEntitlementManagementRoleEligibilityScheduleRequest - Get roleEligibilityScheduleRequests from roleManagement.
// Requests for role eligibilities for principals through PIM.
func (c EntitlementManagementRoleEligibilityScheduleRequestClient) GetEntitlementManagementRoleEligibilityScheduleRequest(ctx context.Context, id stable.RoleManagementEntitlementManagementRoleEligibilityScheduleRequestId, options GetEntitlementManagementRoleEligibilityScheduleRequestOperationOptions) (result GetEntitlementManagementRoleEligibilityScheduleRequestOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          id.ID(),
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

	var model stable.UnifiedRoleEligibilityScheduleRequest
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}

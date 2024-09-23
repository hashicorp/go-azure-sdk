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

type CreateEntitlementManagementRoleEligibilityScheduleRequestOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.UnifiedRoleEligibilityScheduleRequest
}

type CreateEntitlementManagementRoleEligibilityScheduleRequestOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateEntitlementManagementRoleEligibilityScheduleRequestOperationOptions() CreateEntitlementManagementRoleEligibilityScheduleRequestOperationOptions {
	return CreateEntitlementManagementRoleEligibilityScheduleRequestOperationOptions{}
}

func (o CreateEntitlementManagementRoleEligibilityScheduleRequestOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateEntitlementManagementRoleEligibilityScheduleRequestOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateEntitlementManagementRoleEligibilityScheduleRequestOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateEntitlementManagementRoleEligibilityScheduleRequest - Create new navigation property to
// roleEligibilityScheduleRequests for roleManagement
func (c EntitlementManagementRoleEligibilityScheduleRequestClient) CreateEntitlementManagementRoleEligibilityScheduleRequest(ctx context.Context, input stable.UnifiedRoleEligibilityScheduleRequest, options CreateEntitlementManagementRoleEligibilityScheduleRequestOperationOptions) (result CreateEntitlementManagementRoleEligibilityScheduleRequestOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          "/roleManagement/entitlementManagement/roleEligibilityScheduleRequests",
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

	var model stable.UnifiedRoleEligibilityScheduleRequest
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}

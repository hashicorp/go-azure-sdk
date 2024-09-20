package entitlementmanagementroleassignmentschedulerequest

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateEntitlementManagementRoleAssignmentScheduleRequestOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.UnifiedRoleAssignmentScheduleRequest
}

type CreateEntitlementManagementRoleAssignmentScheduleRequestOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultCreateEntitlementManagementRoleAssignmentScheduleRequestOperationOptions() CreateEntitlementManagementRoleAssignmentScheduleRequestOperationOptions {
	return CreateEntitlementManagementRoleAssignmentScheduleRequestOperationOptions{}
}

func (o CreateEntitlementManagementRoleAssignmentScheduleRequestOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateEntitlementManagementRoleAssignmentScheduleRequestOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateEntitlementManagementRoleAssignmentScheduleRequestOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateEntitlementManagementRoleAssignmentScheduleRequest - Create new navigation property to
// roleAssignmentScheduleRequests for roleManagement
func (c EntitlementManagementRoleAssignmentScheduleRequestClient) CreateEntitlementManagementRoleAssignmentScheduleRequest(ctx context.Context, input stable.UnifiedRoleAssignmentScheduleRequest, options CreateEntitlementManagementRoleAssignmentScheduleRequestOperationOptions) (result CreateEntitlementManagementRoleAssignmentScheduleRequestOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          "/roleManagement/entitlementManagement/roleAssignmentScheduleRequests",
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

	var model stable.UnifiedRoleAssignmentScheduleRequest
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}

package entitlementmanagementroleassignmentschedulerequestprincipal

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetEntitlementManagementRoleAssignmentScheduleRequestPrincipalOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        stable.DirectoryObject
}

type GetEntitlementManagementRoleAssignmentScheduleRequestPrincipalOperationOptions struct {
	Expand   *odata.Expand
	Metadata *odata.Metadata
	Select   *[]string
}

func DefaultGetEntitlementManagementRoleAssignmentScheduleRequestPrincipalOperationOptions() GetEntitlementManagementRoleAssignmentScheduleRequestPrincipalOperationOptions {
	return GetEntitlementManagementRoleAssignmentScheduleRequestPrincipalOperationOptions{}
}

func (o GetEntitlementManagementRoleAssignmentScheduleRequestPrincipalOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetEntitlementManagementRoleAssignmentScheduleRequestPrincipalOperationOptions) ToOData() *odata.Query {
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

func (o GetEntitlementManagementRoleAssignmentScheduleRequestPrincipalOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetEntitlementManagementRoleAssignmentScheduleRequestPrincipal - Get principal from roleManagement. The principal
// that's getting a role assignment through the request. Supports $expand and $select nested in $expand for id only.
func (c EntitlementManagementRoleAssignmentScheduleRequestPrincipalClient) GetEntitlementManagementRoleAssignmentScheduleRequestPrincipal(ctx context.Context, id stable.RoleManagementEntitlementManagementRoleAssignmentScheduleRequestId, options GetEntitlementManagementRoleAssignmentScheduleRequestPrincipalOperationOptions) (result GetEntitlementManagementRoleAssignmentScheduleRequestPrincipalOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/principal", id.ID()),
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

	var respObj json.RawMessage
	if err = resp.Unmarshal(&respObj); err != nil {
		return
	}
	model, err := stable.UnmarshalDirectoryObjectImplementation(respObj)
	if err != nil {
		return
	}
	result.Model = model

	return
}

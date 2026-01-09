package entitlementmanagementroleeligibilityscheduleroledefinition

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetEntitlementManagementRoleEligibilityScheduleRoleDefinitionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.UnifiedRoleDefinition
}

type GetEntitlementManagementRoleEligibilityScheduleRoleDefinitionOperationOptions struct {
	Expand    *odata.Expand
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Select    *[]string
}

func DefaultGetEntitlementManagementRoleEligibilityScheduleRoleDefinitionOperationOptions() GetEntitlementManagementRoleEligibilityScheduleRoleDefinitionOperationOptions {
	return GetEntitlementManagementRoleEligibilityScheduleRoleDefinitionOperationOptions{}
}

func (o GetEntitlementManagementRoleEligibilityScheduleRoleDefinitionOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetEntitlementManagementRoleEligibilityScheduleRoleDefinitionOperationOptions) ToOData() *odata.Query {
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

func (o GetEntitlementManagementRoleEligibilityScheduleRoleDefinitionOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetEntitlementManagementRoleEligibilityScheduleRoleDefinition - Get roleDefinition from roleManagement. Detailed
// information for the roleDefinition object that is referenced through the roleDefinitionId property.
func (c EntitlementManagementRoleEligibilityScheduleRoleDefinitionClient) GetEntitlementManagementRoleEligibilityScheduleRoleDefinition(ctx context.Context, id beta.RoleManagementEntitlementManagementRoleEligibilityScheduleId, options GetEntitlementManagementRoleEligibilityScheduleRoleDefinitionOperationOptions) (result GetEntitlementManagementRoleEligibilityScheduleRoleDefinitionOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/roleDefinition", id.ID()),
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

	var model beta.UnifiedRoleDefinition
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}

package enterpriseapproleassignmentscheduleinstanceactivatedusing

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetEnterpriseAppRoleAssignmentScheduleInstanceActivatedUsingOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.UnifiedRoleEligibilityScheduleInstance
}

type GetEnterpriseAppRoleAssignmentScheduleInstanceActivatedUsingOperationOptions struct {
	Expand   *odata.Expand
	Metadata *odata.Metadata
	Select   *[]string
}

func DefaultGetEnterpriseAppRoleAssignmentScheduleInstanceActivatedUsingOperationOptions() GetEnterpriseAppRoleAssignmentScheduleInstanceActivatedUsingOperationOptions {
	return GetEnterpriseAppRoleAssignmentScheduleInstanceActivatedUsingOperationOptions{}
}

func (o GetEnterpriseAppRoleAssignmentScheduleInstanceActivatedUsingOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetEnterpriseAppRoleAssignmentScheduleInstanceActivatedUsingOperationOptions) ToOData() *odata.Query {
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

func (o GetEnterpriseAppRoleAssignmentScheduleInstanceActivatedUsingOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetEnterpriseAppRoleAssignmentScheduleInstanceActivatedUsing - Get activatedUsing from roleManagement. If the request
// is from an eligible administrator to activate a role, this parameter shows the related eligible assignment for that
// activation. Otherwise, it's null. Supports $expand and $select nested in $expand.
func (c EnterpriseAppRoleAssignmentScheduleInstanceActivatedUsingClient) GetEnterpriseAppRoleAssignmentScheduleInstanceActivatedUsing(ctx context.Context, id beta.RoleManagementEnterpriseAppIdRoleAssignmentScheduleInstanceId, options GetEnterpriseAppRoleAssignmentScheduleInstanceActivatedUsingOperationOptions) (result GetEnterpriseAppRoleAssignmentScheduleInstanceActivatedUsingOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/activatedUsing", id.ID()),
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

	var model beta.UnifiedRoleEligibilityScheduleInstance
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}

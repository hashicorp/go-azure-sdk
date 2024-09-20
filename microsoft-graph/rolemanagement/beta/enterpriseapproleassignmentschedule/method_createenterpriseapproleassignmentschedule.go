package enterpriseapproleassignmentschedule

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

type CreateEnterpriseAppRoleAssignmentScheduleOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.UnifiedRoleAssignmentSchedule
}

type CreateEnterpriseAppRoleAssignmentScheduleOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultCreateEnterpriseAppRoleAssignmentScheduleOperationOptions() CreateEnterpriseAppRoleAssignmentScheduleOperationOptions {
	return CreateEnterpriseAppRoleAssignmentScheduleOperationOptions{}
}

func (o CreateEnterpriseAppRoleAssignmentScheduleOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateEnterpriseAppRoleAssignmentScheduleOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateEnterpriseAppRoleAssignmentScheduleOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateEnterpriseAppRoleAssignmentSchedule - Create new navigation property to roleAssignmentSchedules for
// roleManagement
func (c EnterpriseAppRoleAssignmentScheduleClient) CreateEnterpriseAppRoleAssignmentSchedule(ctx context.Context, id beta.RoleManagementEnterpriseAppId, input beta.UnifiedRoleAssignmentSchedule, options CreateEnterpriseAppRoleAssignmentScheduleOperationOptions) (result CreateEnterpriseAppRoleAssignmentScheduleOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/roleAssignmentSchedules", id.ID()),
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

	var model beta.UnifiedRoleAssignmentSchedule
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}

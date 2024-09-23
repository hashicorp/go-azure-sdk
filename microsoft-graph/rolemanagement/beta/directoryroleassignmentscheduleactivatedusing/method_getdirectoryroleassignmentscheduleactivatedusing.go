package directoryroleassignmentscheduleactivatedusing

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

type GetDirectoryRoleAssignmentScheduleActivatedUsingOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.UnifiedRoleEligibilitySchedule
}

type GetDirectoryRoleAssignmentScheduleActivatedUsingOperationOptions struct {
	Expand    *odata.Expand
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Select    *[]string
}

func DefaultGetDirectoryRoleAssignmentScheduleActivatedUsingOperationOptions() GetDirectoryRoleAssignmentScheduleActivatedUsingOperationOptions {
	return GetDirectoryRoleAssignmentScheduleActivatedUsingOperationOptions{}
}

func (o GetDirectoryRoleAssignmentScheduleActivatedUsingOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetDirectoryRoleAssignmentScheduleActivatedUsingOperationOptions) ToOData() *odata.Query {
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

func (o GetDirectoryRoleAssignmentScheduleActivatedUsingOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetDirectoryRoleAssignmentScheduleActivatedUsing - Get activatedUsing from roleManagement. If the request is from an
// eligible administrator to activate a role, this parameter shows the related eligible assignment for that activation.
// Otherwise, it's null. Supports $expand.
func (c DirectoryRoleAssignmentScheduleActivatedUsingClient) GetDirectoryRoleAssignmentScheduleActivatedUsing(ctx context.Context, id beta.RoleManagementDirectoryRoleAssignmentScheduleId, options GetDirectoryRoleAssignmentScheduleActivatedUsingOperationOptions) (result GetDirectoryRoleAssignmentScheduleActivatedUsingOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/activatedUsing", id.ID()),
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

	var model beta.UnifiedRoleEligibilitySchedule
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}

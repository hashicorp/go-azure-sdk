package teamscheduleshiftsroledefinition

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateTeamScheduleShiftsRoleDefinitionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateTeamScheduleShiftsRoleDefinitionOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdateTeamScheduleShiftsRoleDefinitionOperationOptions() UpdateTeamScheduleShiftsRoleDefinitionOperationOptions {
	return UpdateTeamScheduleShiftsRoleDefinitionOperationOptions{}
}

func (o UpdateTeamScheduleShiftsRoleDefinitionOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateTeamScheduleShiftsRoleDefinitionOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateTeamScheduleShiftsRoleDefinitionOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateTeamScheduleShiftsRoleDefinition - Update the navigation property shiftsRoleDefinitions in groups
func (c TeamScheduleShiftsRoleDefinitionClient) UpdateTeamScheduleShiftsRoleDefinition(ctx context.Context, id beta.GroupIdTeamScheduleShiftsRoleDefinitionId, input beta.ShiftsRoleDefinition, options UpdateTeamScheduleShiftsRoleDefinitionOperationOptions) (result UpdateTeamScheduleShiftsRoleDefinitionOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          id.ID(),
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

	return
}

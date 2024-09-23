package teamscheduleschedulinggroup

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetTeamScheduleSchedulingGroupOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.SchedulingGroup
}

type GetTeamScheduleSchedulingGroupOperationOptions struct {
	Expand    *odata.Expand
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Select    *[]string
}

func DefaultGetTeamScheduleSchedulingGroupOperationOptions() GetTeamScheduleSchedulingGroupOperationOptions {
	return GetTeamScheduleSchedulingGroupOperationOptions{}
}

func (o GetTeamScheduleSchedulingGroupOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetTeamScheduleSchedulingGroupOperationOptions) ToOData() *odata.Query {
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

func (o GetTeamScheduleSchedulingGroupOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetTeamScheduleSchedulingGroup - Get schedulingGroups from groups. The logical grouping of users in the schedule
// (usually by role).
func (c TeamScheduleSchedulingGroupClient) GetTeamScheduleSchedulingGroup(ctx context.Context, id beta.GroupIdTeamScheduleSchedulingGroupId, options GetTeamScheduleSchedulingGroupOperationOptions) (result GetTeamScheduleSchedulingGroupOperationResponse, err error) {
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

	var model beta.SchedulingGroup
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}

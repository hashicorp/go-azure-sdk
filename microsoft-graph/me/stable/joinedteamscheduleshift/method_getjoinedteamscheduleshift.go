package joinedteamscheduleshift

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetJoinedTeamScheduleShiftOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.Shift
}

type GetJoinedTeamScheduleShiftOperationOptions struct {
	Expand *odata.Expand
	Select *[]string
}

func DefaultGetJoinedTeamScheduleShiftOperationOptions() GetJoinedTeamScheduleShiftOperationOptions {
	return GetJoinedTeamScheduleShiftOperationOptions{}
}

func (o GetJoinedTeamScheduleShiftOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetJoinedTeamScheduleShiftOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetJoinedTeamScheduleShiftOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetJoinedTeamScheduleShift - Get shifts from me. The shifts in the schedule.
func (c JoinedTeamScheduleShiftClient) GetJoinedTeamScheduleShift(ctx context.Context, id stable.MeJoinedTeamIdScheduleShiftId, options GetJoinedTeamScheduleShiftOperationOptions) (result GetJoinedTeamScheduleShiftOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          id.ID(),
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

	var model stable.Shift
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}

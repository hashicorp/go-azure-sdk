package teamscheduletimeoffrequest

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateTeamScheduleTimeOffRequestOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateTeamScheduleTimeOffRequestOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultUpdateTeamScheduleTimeOffRequestOperationOptions() UpdateTeamScheduleTimeOffRequestOperationOptions {
	return UpdateTeamScheduleTimeOffRequestOperationOptions{}
}

func (o UpdateTeamScheduleTimeOffRequestOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateTeamScheduleTimeOffRequestOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateTeamScheduleTimeOffRequestOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateTeamScheduleTimeOffRequest - Update the navigation property timeOffRequests in groups
func (c TeamScheduleTimeOffRequestClient) UpdateTeamScheduleTimeOffRequest(ctx context.Context, id beta.GroupIdTeamScheduleTimeOffRequestId, input beta.TimeOffRequest, options UpdateTeamScheduleTimeOffRequestOperationOptions) (result UpdateTeamScheduleTimeOffRequestOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          id.ID(),
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

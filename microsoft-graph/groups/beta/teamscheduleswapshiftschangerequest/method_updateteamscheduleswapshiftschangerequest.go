package teamscheduleswapshiftschangerequest

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateTeamScheduleSwapShiftsChangeRequestOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateTeamScheduleSwapShiftsChangeRequestOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultUpdateTeamScheduleSwapShiftsChangeRequestOperationOptions() UpdateTeamScheduleSwapShiftsChangeRequestOperationOptions {
	return UpdateTeamScheduleSwapShiftsChangeRequestOperationOptions{}
}

func (o UpdateTeamScheduleSwapShiftsChangeRequestOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateTeamScheduleSwapShiftsChangeRequestOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateTeamScheduleSwapShiftsChangeRequestOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateTeamScheduleSwapShiftsChangeRequest - Update the navigation property swapShiftsChangeRequests in groups
func (c TeamScheduleSwapShiftsChangeRequestClient) UpdateTeamScheduleSwapShiftsChangeRequest(ctx context.Context, id beta.GroupIdTeamScheduleSwapShiftsChangeRequestId, input beta.SwapShiftsChangeRequest, options UpdateTeamScheduleSwapShiftsChangeRequestOperationOptions) (result UpdateTeamScheduleSwapShiftsChangeRequestOperationResponse, err error) {
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

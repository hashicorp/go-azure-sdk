package teamscheduleswapshiftschangerequest

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateTeamScheduleSwapShiftsChangeRequestOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.SwapShiftsChangeRequest
}

type CreateTeamScheduleSwapShiftsChangeRequestOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultCreateTeamScheduleSwapShiftsChangeRequestOperationOptions() CreateTeamScheduleSwapShiftsChangeRequestOperationOptions {
	return CreateTeamScheduleSwapShiftsChangeRequestOperationOptions{}
}

func (o CreateTeamScheduleSwapShiftsChangeRequestOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateTeamScheduleSwapShiftsChangeRequestOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateTeamScheduleSwapShiftsChangeRequestOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateTeamScheduleSwapShiftsChangeRequest - Create new navigation property to swapShiftsChangeRequests for groups
func (c TeamScheduleSwapShiftsChangeRequestClient) CreateTeamScheduleSwapShiftsChangeRequest(ctx context.Context, id stable.GroupId, input stable.SwapShiftsChangeRequest, options CreateTeamScheduleSwapShiftsChangeRequestOperationOptions) (result CreateTeamScheduleSwapShiftsChangeRequestOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/team/schedule/swapShiftsChangeRequests", id.ID()),
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

	var model stable.SwapShiftsChangeRequest
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}

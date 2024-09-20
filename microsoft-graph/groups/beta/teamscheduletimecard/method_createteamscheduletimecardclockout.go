package teamscheduletimecard

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

type CreateTeamScheduleTimeCardClockOutOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.TimeCard
}

type CreateTeamScheduleTimeCardClockOutOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultCreateTeamScheduleTimeCardClockOutOperationOptions() CreateTeamScheduleTimeCardClockOutOperationOptions {
	return CreateTeamScheduleTimeCardClockOutOperationOptions{}
}

func (o CreateTeamScheduleTimeCardClockOutOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateTeamScheduleTimeCardClockOutOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateTeamScheduleTimeCardClockOutOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateTeamScheduleTimeCardClockOut - Invoke action clockOut. Clock out to end an open timeCard.
func (c TeamScheduleTimeCardClient) CreateTeamScheduleTimeCardClockOut(ctx context.Context, id beta.GroupIdTeamScheduleTimeCardId, input CreateTeamScheduleTimeCardClockOutRequest, options CreateTeamScheduleTimeCardClockOutOperationOptions) (result CreateTeamScheduleTimeCardClockOutOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/clockOut", id.ID()),
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

	var model beta.TimeCard
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}

package joinedteamscheduletimecard

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

type StartJoinedTeamScheduleTimeCardBreakOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.TimeCard
}

type StartJoinedTeamScheduleTimeCardBreakOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultStartJoinedTeamScheduleTimeCardBreakOperationOptions() StartJoinedTeamScheduleTimeCardBreakOperationOptions {
	return StartJoinedTeamScheduleTimeCardBreakOperationOptions{}
}

func (o StartJoinedTeamScheduleTimeCardBreakOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o StartJoinedTeamScheduleTimeCardBreakOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o StartJoinedTeamScheduleTimeCardBreakOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// StartJoinedTeamScheduleTimeCardBreak - Invoke action startBreak
func (c JoinedTeamScheduleTimeCardClient) StartJoinedTeamScheduleTimeCardBreak(ctx context.Context, id stable.UserIdJoinedTeamIdScheduleTimeCardId, input StartJoinedTeamScheduleTimeCardBreakRequest, options StartJoinedTeamScheduleTimeCardBreakOperationOptions) (result StartJoinedTeamScheduleTimeCardBreakOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusCreated,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/startBreak", id.ID()),
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

	var model stable.TimeCard
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}

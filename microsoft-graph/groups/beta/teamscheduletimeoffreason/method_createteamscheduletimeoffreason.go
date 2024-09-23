package teamscheduletimeoffreason

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

type CreateTeamScheduleTimeOffReasonOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.TimeOffReason
}

type CreateTeamScheduleTimeOffReasonOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateTeamScheduleTimeOffReasonOperationOptions() CreateTeamScheduleTimeOffReasonOperationOptions {
	return CreateTeamScheduleTimeOffReasonOperationOptions{}
}

func (o CreateTeamScheduleTimeOffReasonOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateTeamScheduleTimeOffReasonOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateTeamScheduleTimeOffReasonOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateTeamScheduleTimeOffReason - Create new navigation property to timeOffReasons for groups
func (c TeamScheduleTimeOffReasonClient) CreateTeamScheduleTimeOffReason(ctx context.Context, id beta.GroupId, input beta.TimeOffReason, options CreateTeamScheduleTimeOffReasonOperationOptions) (result CreateTeamScheduleTimeOffReasonOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/team/schedule/timeOffReasons", id.ID()),
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

	var model beta.TimeOffReason
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}

package joinedteamscheduleopenshift

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

type CreateJoinedTeamScheduleOpenShiftOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.OpenShift
}

type CreateJoinedTeamScheduleOpenShiftOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateJoinedTeamScheduleOpenShiftOperationOptions() CreateJoinedTeamScheduleOpenShiftOperationOptions {
	return CreateJoinedTeamScheduleOpenShiftOperationOptions{}
}

func (o CreateJoinedTeamScheduleOpenShiftOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateJoinedTeamScheduleOpenShiftOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateJoinedTeamScheduleOpenShiftOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateJoinedTeamScheduleOpenShift - Create new navigation property to openShifts for me
func (c JoinedTeamScheduleOpenShiftClient) CreateJoinedTeamScheduleOpenShift(ctx context.Context, id stable.MeJoinedTeamId, input stable.OpenShift, options CreateJoinedTeamScheduleOpenShiftOperationOptions) (result CreateJoinedTeamScheduleOpenShiftOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/schedule/openShifts", id.ID()),
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

	var model stable.OpenShift
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}

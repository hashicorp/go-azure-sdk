package teamchannelplannerplan

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

type CreateTeamChannelPlannerPlanArchiveOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type CreateTeamChannelPlannerPlanArchiveOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateTeamChannelPlannerPlanArchiveOperationOptions() CreateTeamChannelPlannerPlanArchiveOperationOptions {
	return CreateTeamChannelPlannerPlanArchiveOperationOptions{}
}

func (o CreateTeamChannelPlannerPlanArchiveOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateTeamChannelPlannerPlanArchiveOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateTeamChannelPlannerPlanArchiveOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateTeamChannelPlannerPlanArchive - Invoke action archive. Archive a plannerPlan object. Archiving a plan, also
// archives the plannerTasks and plannerBuckets in the plan. An archived entity is read-only. Archived entities cannot
// be updated. An archived plan can be unarchived. All archived entities can be deleted. Archived tasks are not included
// in the response for list of tasks assigned to a user.
func (c TeamChannelPlannerPlanClient) CreateTeamChannelPlannerPlanArchive(ctx context.Context, id beta.GroupIdTeamChannelIdPlannerPlanId, input CreateTeamChannelPlannerPlanArchiveRequest, options CreateTeamChannelPlannerPlanArchiveOperationOptions) (result CreateTeamChannelPlannerPlanArchiveOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/archive", id.ID()),
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

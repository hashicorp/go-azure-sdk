package teamprimarychannelplannerplantask

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateTeamPrimaryChannelPlannerPlanTaskOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateTeamPrimaryChannelPlannerPlanTaskOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdateTeamPrimaryChannelPlannerPlanTaskOperationOptions() UpdateTeamPrimaryChannelPlannerPlanTaskOperationOptions {
	return UpdateTeamPrimaryChannelPlannerPlanTaskOperationOptions{}
}

func (o UpdateTeamPrimaryChannelPlannerPlanTaskOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateTeamPrimaryChannelPlannerPlanTaskOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateTeamPrimaryChannelPlannerPlanTaskOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateTeamPrimaryChannelPlannerPlanTask - Update the navigation property tasks in groups
func (c TeamPrimaryChannelPlannerPlanTaskClient) UpdateTeamPrimaryChannelPlannerPlanTask(ctx context.Context, id beta.GroupIdTeamPrimaryChannelPlannerPlanIdTaskId, input beta.PlannerTask, options UpdateTeamPrimaryChannelPlannerPlanTaskOperationOptions) (result UpdateTeamPrimaryChannelPlannerPlanTaskOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          id.ID(),
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

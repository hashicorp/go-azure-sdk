package teamprimarychannelplannerplan

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

type CreateTeamPrimaryChannelPlannerPlanOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.PlannerPlan
}

type CreateTeamPrimaryChannelPlannerPlanOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateTeamPrimaryChannelPlannerPlanOperationOptions() CreateTeamPrimaryChannelPlannerPlanOperationOptions {
	return CreateTeamPrimaryChannelPlannerPlanOperationOptions{}
}

func (o CreateTeamPrimaryChannelPlannerPlanOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateTeamPrimaryChannelPlannerPlanOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateTeamPrimaryChannelPlannerPlanOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateTeamPrimaryChannelPlannerPlan - Create new navigation property to plans for groups
func (c TeamPrimaryChannelPlannerPlanClient) CreateTeamPrimaryChannelPlannerPlan(ctx context.Context, id beta.GroupId, input beta.PlannerPlan, options CreateTeamPrimaryChannelPlannerPlanOperationOptions) (result CreateTeamPrimaryChannelPlannerPlanOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/team/primaryChannel/planner/plans", id.ID()),
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

	var model beta.PlannerPlan
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}

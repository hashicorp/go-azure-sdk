package plannerplan

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

type MovePlannerPlanToContainerOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.PlannerPlan
}

type MovePlannerPlanToContainerOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultMovePlannerPlanToContainerOperationOptions() MovePlannerPlanToContainerOperationOptions {
	return MovePlannerPlanToContainerOperationOptions{}
}

func (o MovePlannerPlanToContainerOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o MovePlannerPlanToContainerOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o MovePlannerPlanToContainerOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// MovePlannerPlanToContainer - Invoke action moveToContainer. Move a planner plan object from one planner plan
// container to another. Planner plans can only be moved from a user container to a group container.
func (c PlannerPlanClient) MovePlannerPlanToContainer(ctx context.Context, id beta.UserIdPlannerPlanId, input MovePlannerPlanToContainerRequest, options MovePlannerPlanToContainerOperationOptions) (result MovePlannerPlanToContainerOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/moveToContainer", id.ID()),
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

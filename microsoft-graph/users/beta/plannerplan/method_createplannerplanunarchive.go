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

type CreatePlannerPlanUnarchiveOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type CreatePlannerPlanUnarchiveOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultCreatePlannerPlanUnarchiveOperationOptions() CreatePlannerPlanUnarchiveOperationOptions {
	return CreatePlannerPlanUnarchiveOperationOptions{}
}

func (o CreatePlannerPlanUnarchiveOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreatePlannerPlanUnarchiveOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreatePlannerPlanUnarchiveOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreatePlannerPlanUnarchive - Invoke action unarchive. Unarchive a plannerPlan object. Unarchiving a plan, also
// unarchives the plannerTasks and plannerBuckets in the plan. Only a plan that is archived can be unarchived.
func (c PlannerPlanClient) CreatePlannerPlanUnarchive(ctx context.Context, id beta.UserIdPlannerPlanId, input CreatePlannerPlanUnarchiveRequest, options CreatePlannerPlanUnarchiveOperationOptions) (result CreatePlannerPlanUnarchiveOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/unarchive", id.ID()),
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

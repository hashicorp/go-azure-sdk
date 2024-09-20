package plannertaskassignedtotaskboardformat

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

type UpdatePlannerTaskAssignedToTaskBoardFormatOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdatePlannerTaskAssignedToTaskBoardFormatOperationOptions struct {
	IfMatch  *string
	Metadata *odata.Metadata
}

func DefaultUpdatePlannerTaskAssignedToTaskBoardFormatOperationOptions() UpdatePlannerTaskAssignedToTaskBoardFormatOperationOptions {
	return UpdatePlannerTaskAssignedToTaskBoardFormatOperationOptions{}
}

func (o UpdatePlannerTaskAssignedToTaskBoardFormatOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o UpdatePlannerTaskAssignedToTaskBoardFormatOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdatePlannerTaskAssignedToTaskBoardFormatOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdatePlannerTaskAssignedToTaskBoardFormat - Update the navigation property assignedToTaskBoardFormat in me
func (c PlannerTaskAssignedToTaskBoardFormatClient) UpdatePlannerTaskAssignedToTaskBoardFormat(ctx context.Context, id stable.MePlannerTaskId, input stable.PlannerAssignedToTaskBoardTaskFormat, options UpdatePlannerTaskAssignedToTaskBoardFormatOperationOptions) (result UpdatePlannerTaskAssignedToTaskBoardFormatOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/assignedToTaskBoardFormat", id.ID()),
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

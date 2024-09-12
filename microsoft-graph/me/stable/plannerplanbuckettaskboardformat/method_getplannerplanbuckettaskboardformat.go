package plannerplanbuckettaskboardformat

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

type GetPlannerPlanBucketTaskBoardFormatOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.PlannerBucketTaskBoardTaskFormat
}

type GetPlannerPlanBucketTaskBoardFormatOperationOptions struct {
	Expand *odata.Expand
	Select *[]string
}

func DefaultGetPlannerPlanBucketTaskBoardFormatOperationOptions() GetPlannerPlanBucketTaskBoardFormatOperationOptions {
	return GetPlannerPlanBucketTaskBoardFormatOperationOptions{}
}

func (o GetPlannerPlanBucketTaskBoardFormatOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetPlannerPlanBucketTaskBoardFormatOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetPlannerPlanBucketTaskBoardFormatOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetPlannerPlanBucketTaskBoardFormat - Get bucketTaskBoardFormat from me. Read-only. Nullable. Used to render the task
// correctly in the task board view when grouped by bucket.
func (c PlannerPlanBucketTaskBoardFormatClient) GetPlannerPlanBucketTaskBoardFormat(ctx context.Context, id stable.MePlannerPlanIdBucketIdTaskId, options GetPlannerPlanBucketTaskBoardFormatOperationOptions) (result GetPlannerPlanBucketTaskBoardFormatOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/bucketTaskBoardFormat", id.ID()),
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
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

	var model stable.PlannerBucketTaskBoardTaskFormat
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}

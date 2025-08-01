package teamprimarychannelplannerplanbuckettaskprogresstaskboardformat

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

type GetTeamPrimaryChannelPlannerPlanBucketTaskProgressTaskBoardFormatOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.PlannerProgressTaskBoardTaskFormat
}

type GetTeamPrimaryChannelPlannerPlanBucketTaskProgressTaskBoardFormatOperationOptions struct {
	Expand    *odata.Expand
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Select    *[]string
}

func DefaultGetTeamPrimaryChannelPlannerPlanBucketTaskProgressTaskBoardFormatOperationOptions() GetTeamPrimaryChannelPlannerPlanBucketTaskProgressTaskBoardFormatOperationOptions {
	return GetTeamPrimaryChannelPlannerPlanBucketTaskProgressTaskBoardFormatOperationOptions{}
}

func (o GetTeamPrimaryChannelPlannerPlanBucketTaskProgressTaskBoardFormatOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetTeamPrimaryChannelPlannerPlanBucketTaskProgressTaskBoardFormatOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetTeamPrimaryChannelPlannerPlanBucketTaskProgressTaskBoardFormatOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetTeamPrimaryChannelPlannerPlanBucketTaskProgressTaskBoardFormat - Get progressTaskBoardFormat from groups.
// Read-only. Nullable. Used to render the task correctly in the task board view when grouped by progress.
func (c TeamPrimaryChannelPlannerPlanBucketTaskProgressTaskBoardFormatClient) GetTeamPrimaryChannelPlannerPlanBucketTaskProgressTaskBoardFormat(ctx context.Context, id beta.GroupIdTeamPrimaryChannelPlannerPlanIdBucketIdTaskId, options GetTeamPrimaryChannelPlannerPlanBucketTaskProgressTaskBoardFormatOperationOptions) (result GetTeamPrimaryChannelPlannerPlanBucketTaskProgressTaskBoardFormatOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/progressTaskBoardFormat", id.ID()),
		RetryFunc:     options.RetryFunc,
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

	var model beta.PlannerProgressTaskBoardTaskFormat
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}

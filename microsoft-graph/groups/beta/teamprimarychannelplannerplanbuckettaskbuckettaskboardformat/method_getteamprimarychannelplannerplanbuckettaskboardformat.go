package teamprimarychannelplannerplanbuckettaskbuckettaskboardformat

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetTeamPrimaryChannelPlannerPlanBucketTaskBoardFormatOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.PlannerBucketTaskBoardTaskFormat
}

type GetTeamPrimaryChannelPlannerPlanBucketTaskBoardFormatOperationOptions struct {
	Expand    *odata.Expand
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Select    *[]string
}

func DefaultGetTeamPrimaryChannelPlannerPlanBucketTaskBoardFormatOperationOptions() GetTeamPrimaryChannelPlannerPlanBucketTaskBoardFormatOperationOptions {
	return GetTeamPrimaryChannelPlannerPlanBucketTaskBoardFormatOperationOptions{}
}

func (o GetTeamPrimaryChannelPlannerPlanBucketTaskBoardFormatOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetTeamPrimaryChannelPlannerPlanBucketTaskBoardFormatOperationOptions) ToOData() *odata.Query {
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

func (o GetTeamPrimaryChannelPlannerPlanBucketTaskBoardFormatOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetTeamPrimaryChannelPlannerPlanBucketTaskBoardFormat - Get bucketTaskBoardFormat from groups. Read-only. Nullable.
// Used to render the task correctly in the task board view when grouped by bucket.
func (c TeamPrimaryChannelPlannerPlanBucketTaskBucketTaskBoardFormatClient) GetTeamPrimaryChannelPlannerPlanBucketTaskBoardFormat(ctx context.Context, id beta.GroupIdTeamPrimaryChannelPlannerPlanIdBucketIdTaskId, options GetTeamPrimaryChannelPlannerPlanBucketTaskBoardFormatOperationOptions) (result GetTeamPrimaryChannelPlannerPlanBucketTaskBoardFormatOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/bucketTaskBoardFormat", id.ID()),
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

	var model beta.PlannerBucketTaskBoardTaskFormat
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}

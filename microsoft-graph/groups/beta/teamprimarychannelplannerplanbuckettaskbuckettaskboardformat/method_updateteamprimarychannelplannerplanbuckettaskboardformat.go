package teamprimarychannelplannerplanbuckettaskbuckettaskboardformat

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

type UpdateTeamPrimaryChannelPlannerPlanBucketTaskBoardFormatOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateTeamPrimaryChannelPlannerPlanBucketTaskBoardFormatOperationOptions struct {
	IfMatch   *string
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdateTeamPrimaryChannelPlannerPlanBucketTaskBoardFormatOperationOptions() UpdateTeamPrimaryChannelPlannerPlanBucketTaskBoardFormatOperationOptions {
	return UpdateTeamPrimaryChannelPlannerPlanBucketTaskBoardFormatOperationOptions{}
}

func (o UpdateTeamPrimaryChannelPlannerPlanBucketTaskBoardFormatOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o UpdateTeamPrimaryChannelPlannerPlanBucketTaskBoardFormatOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateTeamPrimaryChannelPlannerPlanBucketTaskBoardFormatOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateTeamPrimaryChannelPlannerPlanBucketTaskBoardFormat - Update the navigation property bucketTaskBoardFormat in
// groups
func (c TeamPrimaryChannelPlannerPlanBucketTaskBucketTaskBoardFormatClient) UpdateTeamPrimaryChannelPlannerPlanBucketTaskBoardFormat(ctx context.Context, id beta.GroupIdTeamPrimaryChannelPlannerPlanIdBucketIdTaskId, input beta.PlannerBucketTaskBoardTaskFormat, options UpdateTeamPrimaryChannelPlannerPlanBucketTaskBoardFormatOperationOptions) (result UpdateTeamPrimaryChannelPlannerPlanBucketTaskBoardFormatOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/bucketTaskBoardFormat", id.ID()),
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

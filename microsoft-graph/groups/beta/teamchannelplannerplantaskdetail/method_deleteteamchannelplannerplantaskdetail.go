package teamchannelplannerplantaskdetail

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

type DeleteTeamChannelPlannerPlanTaskDetailOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type DeleteTeamChannelPlannerPlanTaskDetailOperationOptions struct {
	IfMatch   *string
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultDeleteTeamChannelPlannerPlanTaskDetailOperationOptions() DeleteTeamChannelPlannerPlanTaskDetailOperationOptions {
	return DeleteTeamChannelPlannerPlanTaskDetailOperationOptions{}
}

func (o DeleteTeamChannelPlannerPlanTaskDetailOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o DeleteTeamChannelPlannerPlanTaskDetailOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o DeleteTeamChannelPlannerPlanTaskDetailOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// DeleteTeamChannelPlannerPlanTaskDetail - Delete navigation property details for groups
func (c TeamChannelPlannerPlanTaskDetailClient) DeleteTeamChannelPlannerPlanTaskDetail(ctx context.Context, id beta.GroupIdTeamChannelIdPlannerPlanIdTaskId, options DeleteTeamChannelPlannerPlanTaskDetailOperationOptions) (result DeleteTeamChannelPlannerPlanTaskDetailOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodDelete,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/details", id.ID()),
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

	return
}

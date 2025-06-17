package teamchannelplannerplandetail

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

type UpdateTeamChannelPlannerPlanDetailOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateTeamChannelPlannerPlanDetailOperationOptions struct {
	IfMatch   *string
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdateTeamChannelPlannerPlanDetailOperationOptions() UpdateTeamChannelPlannerPlanDetailOperationOptions {
	return UpdateTeamChannelPlannerPlanDetailOperationOptions{}
}

func (o UpdateTeamChannelPlannerPlanDetailOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o UpdateTeamChannelPlannerPlanDetailOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateTeamChannelPlannerPlanDetailOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateTeamChannelPlannerPlanDetail - Update the navigation property details in groups
func (c TeamChannelPlannerPlanDetailClient) UpdateTeamChannelPlannerPlanDetail(ctx context.Context, id beta.GroupIdTeamChannelIdPlannerPlanId, input beta.PlannerPlanDetails, options UpdateTeamChannelPlannerPlanDetailOperationOptions) (result UpdateTeamChannelPlannerPlanDetailOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/details", id.ID()),
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

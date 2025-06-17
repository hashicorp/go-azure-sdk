package teamprimarychannelplannerplanbucket

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateTeamPrimaryChannelPlannerPlanBucketOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateTeamPrimaryChannelPlannerPlanBucketOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdateTeamPrimaryChannelPlannerPlanBucketOperationOptions() UpdateTeamPrimaryChannelPlannerPlanBucketOperationOptions {
	return UpdateTeamPrimaryChannelPlannerPlanBucketOperationOptions{}
}

func (o UpdateTeamPrimaryChannelPlannerPlanBucketOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateTeamPrimaryChannelPlannerPlanBucketOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateTeamPrimaryChannelPlannerPlanBucketOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateTeamPrimaryChannelPlannerPlanBucket - Update the navigation property buckets in groups
func (c TeamPrimaryChannelPlannerPlanBucketClient) UpdateTeamPrimaryChannelPlannerPlanBucket(ctx context.Context, id beta.GroupIdTeamPrimaryChannelPlannerPlanIdBucketId, input beta.PlannerBucket, options UpdateTeamPrimaryChannelPlannerPlanBucketOperationOptions) (result UpdateTeamPrimaryChannelPlannerPlanBucketOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          id.ID(),
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

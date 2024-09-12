package teamincomingchannel

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetTeamIncomingChannelOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.Channel
}

type GetTeamIncomingChannelOperationOptions struct {
	Expand *odata.Expand
	Select *[]string
}

func DefaultGetTeamIncomingChannelOperationOptions() GetTeamIncomingChannelOperationOptions {
	return GetTeamIncomingChannelOperationOptions{}
}

func (o GetTeamIncomingChannelOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetTeamIncomingChannelOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetTeamIncomingChannelOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetTeamIncomingChannel - Get incomingChannels from groups. List of channels shared with the team.
func (c TeamIncomingChannelClient) GetTeamIncomingChannel(ctx context.Context, id beta.GroupIdTeamIncomingChannelId, options GetTeamIncomingChannelOperationOptions) (result GetTeamIncomingChannelOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          id.ID(),
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

	var model beta.Channel
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}

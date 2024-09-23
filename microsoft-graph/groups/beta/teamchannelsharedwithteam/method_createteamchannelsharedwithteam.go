package teamchannelsharedwithteam

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

type CreateTeamChannelSharedWithTeamOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.SharedWithChannelTeamInfo
}

type CreateTeamChannelSharedWithTeamOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateTeamChannelSharedWithTeamOperationOptions() CreateTeamChannelSharedWithTeamOperationOptions {
	return CreateTeamChannelSharedWithTeamOperationOptions{}
}

func (o CreateTeamChannelSharedWithTeamOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateTeamChannelSharedWithTeamOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateTeamChannelSharedWithTeamOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateTeamChannelSharedWithTeam - Create new navigation property to sharedWithTeams for groups
func (c TeamChannelSharedWithTeamClient) CreateTeamChannelSharedWithTeam(ctx context.Context, id beta.GroupIdTeamChannelId, input beta.SharedWithChannelTeamInfo, options CreateTeamChannelSharedWithTeamOperationOptions) (result CreateTeamChannelSharedWithTeamOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/sharedWithTeams", id.ID()),
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

	var model beta.SharedWithChannelTeamInfo
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}

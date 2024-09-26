package teamprimarychannelsharedwithteam

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

type CreateTeamPrimaryChannelSharedWithTeamOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.SharedWithChannelTeamInfo
}

type CreateTeamPrimaryChannelSharedWithTeamOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateTeamPrimaryChannelSharedWithTeamOperationOptions() CreateTeamPrimaryChannelSharedWithTeamOperationOptions {
	return CreateTeamPrimaryChannelSharedWithTeamOperationOptions{}
}

func (o CreateTeamPrimaryChannelSharedWithTeamOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateTeamPrimaryChannelSharedWithTeamOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateTeamPrimaryChannelSharedWithTeamOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateTeamPrimaryChannelSharedWithTeam - Create new navigation property to sharedWithTeams for groups
func (c TeamPrimaryChannelSharedWithTeamClient) CreateTeamPrimaryChannelSharedWithTeam(ctx context.Context, id stable.GroupId, input stable.SharedWithChannelTeamInfo, options CreateTeamPrimaryChannelSharedWithTeamOperationOptions) (result CreateTeamPrimaryChannelSharedWithTeamOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusCreated,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/team/primaryChannel/sharedWithTeams", id.ID()),
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

	var model stable.SharedWithChannelTeamInfo
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}

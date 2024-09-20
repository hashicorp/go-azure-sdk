package joinedteamprimarychannelsharedwithteamteam

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

type GetJoinedTeamPrimaryChannelSharedWithTeamTeamOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.Team
}

type GetJoinedTeamPrimaryChannelSharedWithTeamTeamOperationOptions struct {
	Expand   *odata.Expand
	Metadata *odata.Metadata
	Select   *[]string
}

func DefaultGetJoinedTeamPrimaryChannelSharedWithTeamTeamOperationOptions() GetJoinedTeamPrimaryChannelSharedWithTeamTeamOperationOptions {
	return GetJoinedTeamPrimaryChannelSharedWithTeamTeamOperationOptions{}
}

func (o GetJoinedTeamPrimaryChannelSharedWithTeamTeamOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetJoinedTeamPrimaryChannelSharedWithTeamTeamOperationOptions) ToOData() *odata.Query {
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

func (o GetJoinedTeamPrimaryChannelSharedWithTeamTeamOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetJoinedTeamPrimaryChannelSharedWithTeamTeam - Get team from me
func (c JoinedTeamPrimaryChannelSharedWithTeamTeamClient) GetJoinedTeamPrimaryChannelSharedWithTeamTeam(ctx context.Context, id stable.MeJoinedTeamIdPrimaryChannelSharedWithTeamId, options GetJoinedTeamPrimaryChannelSharedWithTeamTeamOperationOptions) (result GetJoinedTeamPrimaryChannelSharedWithTeamTeamOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/team", id.ID()),
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

	var model stable.Team
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}

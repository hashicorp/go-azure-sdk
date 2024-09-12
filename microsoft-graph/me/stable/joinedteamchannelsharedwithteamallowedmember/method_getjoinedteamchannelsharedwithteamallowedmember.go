package joinedteamchannelsharedwithteamallowedmember

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetJoinedTeamChannelSharedWithTeamAllowedMemberOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        stable.ConversationMember
}

type GetJoinedTeamChannelSharedWithTeamAllowedMemberOperationOptions struct {
	Expand *odata.Expand
	Select *[]string
}

func DefaultGetJoinedTeamChannelSharedWithTeamAllowedMemberOperationOptions() GetJoinedTeamChannelSharedWithTeamAllowedMemberOperationOptions {
	return GetJoinedTeamChannelSharedWithTeamAllowedMemberOperationOptions{}
}

func (o GetJoinedTeamChannelSharedWithTeamAllowedMemberOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetJoinedTeamChannelSharedWithTeamAllowedMemberOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetJoinedTeamChannelSharedWithTeamAllowedMemberOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetJoinedTeamChannelSharedWithTeamAllowedMember - Get allowedMembers from me. A collection of team members who have
// access to the shared channel.
func (c JoinedTeamChannelSharedWithTeamAllowedMemberClient) GetJoinedTeamChannelSharedWithTeamAllowedMember(ctx context.Context, id stable.MeJoinedTeamIdChannelIdSharedWithTeamIdAllowedMemberId, options GetJoinedTeamChannelSharedWithTeamAllowedMemberOperationOptions) (result GetJoinedTeamChannelSharedWithTeamAllowedMemberOperationResponse, err error) {
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

	var respObj json.RawMessage
	if err = resp.Unmarshal(&respObj); err != nil {
		return
	}
	model, err := stable.UnmarshalConversationMemberImplementation(respObj)
	if err != nil {
		return
	}
	result.Model = model

	return
}

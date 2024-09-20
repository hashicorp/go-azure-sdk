package joinedteamchannelmessagereply

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

type UnsetJoinedTeamChannelMessageReplyReactionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UnsetJoinedTeamChannelMessageReplyReactionOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultUnsetJoinedTeamChannelMessageReplyReactionOperationOptions() UnsetJoinedTeamChannelMessageReplyReactionOperationOptions {
	return UnsetJoinedTeamChannelMessageReplyReactionOperationOptions{}
}

func (o UnsetJoinedTeamChannelMessageReplyReactionOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UnsetJoinedTeamChannelMessageReplyReactionOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UnsetJoinedTeamChannelMessageReplyReactionOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UnsetJoinedTeamChannelMessageReplyReaction - Invoke action unsetReaction
func (c JoinedTeamChannelMessageReplyClient) UnsetJoinedTeamChannelMessageReplyReaction(ctx context.Context, id stable.MeJoinedTeamIdChannelIdMessageIdReplyId, input UnsetJoinedTeamChannelMessageReplyReactionRequest, options UnsetJoinedTeamChannelMessageReplyReactionOperationOptions) (result UnsetJoinedTeamChannelMessageReplyReactionOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/unsetReaction", id.ID()),
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

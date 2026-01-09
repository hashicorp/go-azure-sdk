package joinedteamprimarychannelmessagereply

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateJoinedTeamPrimaryChannelMessageReplyOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateJoinedTeamPrimaryChannelMessageReplyOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdateJoinedTeamPrimaryChannelMessageReplyOperationOptions() UpdateJoinedTeamPrimaryChannelMessageReplyOperationOptions {
	return UpdateJoinedTeamPrimaryChannelMessageReplyOperationOptions{}
}

func (o UpdateJoinedTeamPrimaryChannelMessageReplyOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateJoinedTeamPrimaryChannelMessageReplyOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateJoinedTeamPrimaryChannelMessageReplyOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateJoinedTeamPrimaryChannelMessageReply - Update the navigation property replies in me
func (c JoinedTeamPrimaryChannelMessageReplyClient) UpdateJoinedTeamPrimaryChannelMessageReply(ctx context.Context, id stable.MeJoinedTeamIdPrimaryChannelMessageIdReplyId, input stable.ChatMessage, options UpdateJoinedTeamPrimaryChannelMessageReplyOperationOptions) (result UpdateJoinedTeamPrimaryChannelMessageReplyOperationResponse, err error) {
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

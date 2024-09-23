package joinedteamprimarychannelmessagereply

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetJoinedTeamPrimaryChannelMessageReplyOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.ChatMessage
}

type GetJoinedTeamPrimaryChannelMessageReplyOperationOptions struct {
	Expand    *odata.Expand
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Select    *[]string
}

func DefaultGetJoinedTeamPrimaryChannelMessageReplyOperationOptions() GetJoinedTeamPrimaryChannelMessageReplyOperationOptions {
	return GetJoinedTeamPrimaryChannelMessageReplyOperationOptions{}
}

func (o GetJoinedTeamPrimaryChannelMessageReplyOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetJoinedTeamPrimaryChannelMessageReplyOperationOptions) ToOData() *odata.Query {
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

func (o GetJoinedTeamPrimaryChannelMessageReplyOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetJoinedTeamPrimaryChannelMessageReply - Get replies from me. Replies for a specified message. Supports $expand for
// channel messages.
func (c JoinedTeamPrimaryChannelMessageReplyClient) GetJoinedTeamPrimaryChannelMessageReply(ctx context.Context, id stable.MeJoinedTeamIdPrimaryChannelMessageIdReplyId, options GetJoinedTeamPrimaryChannelMessageReplyOperationOptions) (result GetJoinedTeamPrimaryChannelMessageReplyOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          id.ID(),
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

	var model stable.ChatMessage
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}

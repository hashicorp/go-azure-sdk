package teamchannelmessagereply

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

type ReplyTeamChannelMessageRepliesWithQuoteOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.ChatMessage
}

type ReplyTeamChannelMessageRepliesWithQuoteOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultReplyTeamChannelMessageRepliesWithQuoteOperationOptions() ReplyTeamChannelMessageRepliesWithQuoteOperationOptions {
	return ReplyTeamChannelMessageRepliesWithQuoteOperationOptions{}
}

func (o ReplyTeamChannelMessageRepliesWithQuoteOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ReplyTeamChannelMessageRepliesWithQuoteOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o ReplyTeamChannelMessageRepliesWithQuoteOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// ReplyTeamChannelMessageRepliesWithQuote - Invoke action replyWithQuote. Reply with quote to a single chat message or
// multiple chat messages in a chat.
func (c TeamChannelMessageReplyClient) ReplyTeamChannelMessageRepliesWithQuote(ctx context.Context, id beta.GroupIdTeamChannelIdMessageId, input ReplyTeamChannelMessageRepliesWithQuoteRequest, options ReplyTeamChannelMessageRepliesWithQuoteOperationOptions) (result ReplyTeamChannelMessageRepliesWithQuoteOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/replies/replyWithQuote", id.ID()),
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

	var model beta.ChatMessage
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}

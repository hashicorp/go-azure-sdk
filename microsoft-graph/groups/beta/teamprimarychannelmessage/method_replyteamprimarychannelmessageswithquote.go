package teamprimarychannelmessage

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

type ReplyTeamPrimaryChannelMessagesWithQuoteOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.ChatMessage
}

type ReplyTeamPrimaryChannelMessagesWithQuoteOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultReplyTeamPrimaryChannelMessagesWithQuoteOperationOptions() ReplyTeamPrimaryChannelMessagesWithQuoteOperationOptions {
	return ReplyTeamPrimaryChannelMessagesWithQuoteOperationOptions{}
}

func (o ReplyTeamPrimaryChannelMessagesWithQuoteOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ReplyTeamPrimaryChannelMessagesWithQuoteOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o ReplyTeamPrimaryChannelMessagesWithQuoteOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// ReplyTeamPrimaryChannelMessagesWithQuote - Invoke action replyWithQuote. Reply with quote to a single chat message or
// multiple chat messages in a chat.
func (c TeamPrimaryChannelMessageClient) ReplyTeamPrimaryChannelMessagesWithQuote(ctx context.Context, id beta.GroupId, input ReplyTeamPrimaryChannelMessagesWithQuoteRequest, options ReplyTeamPrimaryChannelMessagesWithQuoteOperationOptions) (result ReplyTeamPrimaryChannelMessagesWithQuoteOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/team/primaryChannel/messages/replyWithQuote", id.ID()),
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

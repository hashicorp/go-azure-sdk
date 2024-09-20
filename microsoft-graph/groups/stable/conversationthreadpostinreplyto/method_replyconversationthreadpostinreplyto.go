package conversationthreadpostinreplyto

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

type ReplyConversationThreadPostInReplyToOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type ReplyConversationThreadPostInReplyToOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultReplyConversationThreadPostInReplyToOperationOptions() ReplyConversationThreadPostInReplyToOperationOptions {
	return ReplyConversationThreadPostInReplyToOperationOptions{}
}

func (o ReplyConversationThreadPostInReplyToOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ReplyConversationThreadPostInReplyToOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o ReplyConversationThreadPostInReplyToOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// ReplyConversationThreadPostInReplyTo - Invoke action reply
func (c ConversationThreadPostInReplyToClient) ReplyConversationThreadPostInReplyTo(ctx context.Context, id stable.GroupIdConversationIdThreadIdPostId, input ReplyConversationThreadPostInReplyToRequest, options ReplyConversationThreadPostInReplyToOperationOptions) (result ReplyConversationThreadPostInReplyToOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/inReplyTo/reply", id.ID()),
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

package conversationthreadpostinreplytomention

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

type DeleteConversationThreadPostInReplyToMentionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type DeleteConversationThreadPostInReplyToMentionOperationOptions struct {
	IfMatch  *string
	Metadata *odata.Metadata
}

func DefaultDeleteConversationThreadPostInReplyToMentionOperationOptions() DeleteConversationThreadPostInReplyToMentionOperationOptions {
	return DeleteConversationThreadPostInReplyToMentionOperationOptions{}
}

func (o DeleteConversationThreadPostInReplyToMentionOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o DeleteConversationThreadPostInReplyToMentionOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o DeleteConversationThreadPostInReplyToMentionOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// DeleteConversationThreadPostInReplyToMention - Delete navigation property mentions for groups
func (c ConversationThreadPostInReplyToMentionClient) DeleteConversationThreadPostInReplyToMention(ctx context.Context, id beta.GroupIdConversationIdThreadIdPostIdInReplyToMentionId, options DeleteConversationThreadPostInReplyToMentionOperationOptions) (result DeleteConversationThreadPostInReplyToMentionOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodDelete,
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

	return
}

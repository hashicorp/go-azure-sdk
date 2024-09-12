package conversationthreadpostmention

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

type DeleteConversationThreadPostMentionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type DeleteConversationThreadPostMentionOperationOptions struct {
	IfMatch *string
}

func DefaultDeleteConversationThreadPostMentionOperationOptions() DeleteConversationThreadPostMentionOperationOptions {
	return DeleteConversationThreadPostMentionOperationOptions{}
}

func (o DeleteConversationThreadPostMentionOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o DeleteConversationThreadPostMentionOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o DeleteConversationThreadPostMentionOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// DeleteConversationThreadPostMention - Delete navigation property mentions for groups
func (c ConversationThreadPostMentionClient) DeleteConversationThreadPostMention(ctx context.Context, id beta.GroupIdConversationIdThreadIdPostIdMentionId, options DeleteConversationThreadPostMentionOperationOptions) (result DeleteConversationThreadPostMentionOperationResponse, err error) {
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

package conversationthreadpostinreplytoattachment

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

type DeleteConversationThreadPostInReplyToAttachmentOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type DeleteConversationThreadPostInReplyToAttachmentOperationOptions struct {
	IfMatch *string
}

func DefaultDeleteConversationThreadPostInReplyToAttachmentOperationOptions() DeleteConversationThreadPostInReplyToAttachmentOperationOptions {
	return DeleteConversationThreadPostInReplyToAttachmentOperationOptions{}
}

func (o DeleteConversationThreadPostInReplyToAttachmentOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o DeleteConversationThreadPostInReplyToAttachmentOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o DeleteConversationThreadPostInReplyToAttachmentOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// DeleteConversationThreadPostInReplyToAttachment - Delete navigation property attachments for groups
func (c ConversationThreadPostInReplyToAttachmentClient) DeleteConversationThreadPostInReplyToAttachment(ctx context.Context, id stable.GroupIdConversationIdThreadIdPostIdInReplyToAttachmentId, options DeleteConversationThreadPostInReplyToAttachmentOperationOptions) (result DeleteConversationThreadPostInReplyToAttachmentOperationResponse, err error) {
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

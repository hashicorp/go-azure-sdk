package message

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

type ReplyMessageOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type ReplyMessageOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultReplyMessageOperationOptions() ReplyMessageOperationOptions {
	return ReplyMessageOperationOptions{}
}

func (o ReplyMessageOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ReplyMessageOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o ReplyMessageOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// ReplyMessage - Invoke action reply. Reply to the sender of a message using either JSON or MIME format. When using
// JSON format: * Specify either a comment or the body property of the message parameter. Specifying both will return an
// HTTP 400 Bad Request error. * If the original message specifies a recipient in the replyTo property, per Internet
// Message Format (RFC 2822), send the reply to the recipients in replyTo and not the recipient in the from property.
// When using MIME format: - Provide the applicable Internet message headers and the MIME content, all encoded in base64
// format in the request body. - Add any attachments and S/MIME properties to the MIME content. This method saves the
// message in the Sent Items folder. Alternatively, create a draft to reply to an existing message and send it later.
func (c MessageClient) ReplyMessage(ctx context.Context, id stable.UserIdMessageId, input ReplyMessageRequest, options ReplyMessageOperationOptions) (result ReplyMessageOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/reply", id.ID()),
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

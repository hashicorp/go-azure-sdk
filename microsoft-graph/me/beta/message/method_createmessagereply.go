package message

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateMessageReplyOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        beta.Message
}

type CreateMessageReplyOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultCreateMessageReplyOperationOptions() CreateMessageReplyOperationOptions {
	return CreateMessageReplyOperationOptions{}
}

func (o CreateMessageReplyOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateMessageReplyOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateMessageReplyOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateMessageReply - Invoke action createReply. Create a draft to reply to the sender of a message in either JSON or
// MIME format. When using JSON format: - Specify either a comment or the body property of the message parameter.
// Specifying both will return an HTTP 400 Bad Request error. - If replyTo is specified in the original message, per
// Internet Message Format (RFC 2822), you should send the reply to the recipients in replyTo, and not the recipients in
// from. - You can update the draft later to add reply content to the body or change other message properties. When
// using MIME format: - Provide the applicable Internet message headers and the MIME content, all encoded in base64
// format in the request body. - Add any attachments and S/MIME properties to the MIME content. Send the draft message
// in a subsequent operation. Alternatively, reply to a message in a single operation.
func (c MessageClient) CreateMessageReply(ctx context.Context, id beta.MeMessageId, input CreateMessageReplyRequest, options CreateMessageReplyOperationOptions) (result CreateMessageReplyOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/createReply", id.ID()),
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

	var respObj json.RawMessage
	if err = resp.Unmarshal(&respObj); err != nil {
		return
	}
	model, err := beta.UnmarshalMessageImplementation(respObj)
	if err != nil {
		return
	}
	result.Model = model

	return
}

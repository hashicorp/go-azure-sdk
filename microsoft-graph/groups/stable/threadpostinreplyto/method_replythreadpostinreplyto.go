package threadpostinreplyto

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

type ReplyThreadPostInReplyToOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type ReplyThreadPostInReplyToOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultReplyThreadPostInReplyToOperationOptions() ReplyThreadPostInReplyToOperationOptions {
	return ReplyThreadPostInReplyToOperationOptions{}
}

func (o ReplyThreadPostInReplyToOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ReplyThreadPostInReplyToOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o ReplyThreadPostInReplyToOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// ReplyThreadPostInReplyTo - Invoke action reply
func (c ThreadPostInReplyToClient) ReplyThreadPostInReplyTo(ctx context.Context, id stable.GroupIdThreadIdPostId, input ReplyThreadPostInReplyToRequest, options ReplyThreadPostInReplyToOperationOptions) (result ReplyThreadPostInReplyToOperationResponse, err error) {
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

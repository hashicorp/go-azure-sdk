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

type ForwardThreadPostInReplyToOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type ForwardThreadPostInReplyToOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultForwardThreadPostInReplyToOperationOptions() ForwardThreadPostInReplyToOperationOptions {
	return ForwardThreadPostInReplyToOperationOptions{}
}

func (o ForwardThreadPostInReplyToOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ForwardThreadPostInReplyToOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o ForwardThreadPostInReplyToOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// ForwardThreadPostInReplyTo - Invoke action forward. Forward a post to a recipient. You can specify both the parent
// conversation and thread in the request, or, you can specify just the parent thread without the parent conversation.
func (c ThreadPostInReplyToClient) ForwardThreadPostInReplyTo(ctx context.Context, id stable.GroupIdThreadIdPostId, input ForwardThreadPostInReplyToRequest, options ForwardThreadPostInReplyToOperationOptions) (result ForwardThreadPostInReplyToOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/inReplyTo/forward", id.ID()),
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

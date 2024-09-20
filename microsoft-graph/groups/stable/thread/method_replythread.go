package thread

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

type ReplyThreadOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type ReplyThreadOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultReplyThreadOperationOptions() ReplyThreadOperationOptions {
	return ReplyThreadOperationOptions{}
}

func (o ReplyThreadOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ReplyThreadOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o ReplyThreadOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// ReplyThread - Invoke action reply. Reply to a post and add a new post to the specified thread in a group
// conversation. You can specify both the parent conversation and thread in the request, or, you can specify just the
// parent thread without the parent conversation.
func (c ThreadClient) ReplyThread(ctx context.Context, id stable.GroupIdThreadId, input ReplyThreadRequest, options ReplyThreadOperationOptions) (result ReplyThreadOperationResponse, err error) {
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

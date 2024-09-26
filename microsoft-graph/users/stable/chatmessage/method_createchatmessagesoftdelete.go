package chatmessage

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

type CreateChatMessageSoftDeleteOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type CreateChatMessageSoftDeleteOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateChatMessageSoftDeleteOperationOptions() CreateChatMessageSoftDeleteOperationOptions {
	return CreateChatMessageSoftDeleteOperationOptions{}
}

func (o CreateChatMessageSoftDeleteOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateChatMessageSoftDeleteOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateChatMessageSoftDeleteOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateChatMessageSoftDelete - Invoke action softDelete. Delete a single chatMessage or a chat message reply in a
// channel or a chat.
func (c ChatMessageClient) CreateChatMessageSoftDelete(ctx context.Context, id stable.UserIdChatIdMessageId, options CreateChatMessageSoftDeleteOperationOptions) (result CreateChatMessageSoftDeleteOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/softDelete", id.ID()),
		RetryFunc:     options.RetryFunc,
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

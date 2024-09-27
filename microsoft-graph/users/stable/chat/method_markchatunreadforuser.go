package chat

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

type MarkChatUnreadForUserOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type MarkChatUnreadForUserOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultMarkChatUnreadForUserOperationOptions() MarkChatUnreadForUserOperationOptions {
	return MarkChatUnreadForUserOperationOptions{}
}

func (o MarkChatUnreadForUserOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o MarkChatUnreadForUserOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o MarkChatUnreadForUserOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// MarkChatUnreadForUser - Invoke action markChatUnreadForUser. Mark a chat as unread for a user.
func (c ChatClient) MarkChatUnreadForUser(ctx context.Context, id stable.UserIdChatId, input MarkChatUnreadForUserRequest, options MarkChatUnreadForUserOperationOptions) (result MarkChatUnreadForUserOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/markChatUnreadForUser", id.ID()),
		RetryFunc:     options.RetryFunc,
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

package chatpinnedmessage

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

type CreateChatPinnedMessageOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.PinnedChatMessageInfo
}

type CreateChatPinnedMessageOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultCreateChatPinnedMessageOperationOptions() CreateChatPinnedMessageOperationOptions {
	return CreateChatPinnedMessageOperationOptions{}
}

func (o CreateChatPinnedMessageOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateChatPinnedMessageOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateChatPinnedMessageOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateChatPinnedMessage - Create new navigation property to pinnedMessages for users
func (c ChatPinnedMessageClient) CreateChatPinnedMessage(ctx context.Context, id beta.UserIdChatId, input beta.PinnedChatMessageInfo, options CreateChatPinnedMessageOperationOptions) (result CreateChatPinnedMessageOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/pinnedMessages", id.ID()),
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

	var model beta.PinnedChatMessageInfo
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}

package chatlastmessagepreview

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

type UpdateChatLastMessagePreviewOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateChatLastMessagePreviewOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultUpdateChatLastMessagePreviewOperationOptions() UpdateChatLastMessagePreviewOperationOptions {
	return UpdateChatLastMessagePreviewOperationOptions{}
}

func (o UpdateChatLastMessagePreviewOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateChatLastMessagePreviewOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateChatLastMessagePreviewOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateChatLastMessagePreview - Update the navigation property lastMessagePreview in users
func (c ChatLastMessagePreviewClient) UpdateChatLastMessagePreview(ctx context.Context, id beta.UserIdChatId, input beta.ChatMessageInfo, options UpdateChatLastMessagePreviewOperationOptions) (result UpdateChatLastMessagePreviewOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/lastMessagePreview", id.ID()),
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

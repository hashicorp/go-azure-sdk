package conversation

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

type CreateConversationOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.Conversation
}

type CreateConversationOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultCreateConversationOperationOptions() CreateConversationOperationOptions {
	return CreateConversationOperationOptions{}
}

func (o CreateConversationOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateConversationOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateConversationOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateConversation - Create openTypeExtension. Create an open extension (openTypeExtension object) and add custom
// properties in a new or existing instance of a resource. You can create an open extension in a resource instance and
// store custom data to it all in the same operation, except for specific resources. The table in the Permissions
// section lists the resources that support open extensions.
func (c ConversationClient) CreateConversation(ctx context.Context, id beta.GroupId, input beta.Conversation, options CreateConversationOperationOptions) (result CreateConversationOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/conversations", id.ID()),
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

	var model beta.Conversation
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}

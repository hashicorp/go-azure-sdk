package threadpostinreplytoextension

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateThreadPostInReplyToExtensionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        stable.Extension
}

type CreateThreadPostInReplyToExtensionOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateThreadPostInReplyToExtensionOperationOptions() CreateThreadPostInReplyToExtensionOperationOptions {
	return CreateThreadPostInReplyToExtensionOperationOptions{}
}

func (o CreateThreadPostInReplyToExtensionOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateThreadPostInReplyToExtensionOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateThreadPostInReplyToExtensionOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateThreadPostInReplyToExtension - Create new navigation property to extensions for groups
func (c ThreadPostInReplyToExtensionClient) CreateThreadPostInReplyToExtension(ctx context.Context, id stable.GroupIdThreadIdPostId, input stable.Extension, options CreateThreadPostInReplyToExtensionOperationOptions) (result CreateThreadPostInReplyToExtensionOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/inReplyTo/extensions", id.ID()),
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

	var respObj json.RawMessage
	if err = resp.Unmarshal(&respObj); err != nil {
		return
	}
	model, err := stable.UnmarshalExtensionImplementation(respObj)
	if err != nil {
		return
	}
	result.Model = model

	return
}

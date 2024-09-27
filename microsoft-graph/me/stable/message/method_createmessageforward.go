package message

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

type CreateMessageForwardOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        stable.Message
}

type CreateMessageForwardOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateMessageForwardOperationOptions() CreateMessageForwardOperationOptions {
	return CreateMessageForwardOperationOptions{}
}

func (o CreateMessageForwardOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateMessageForwardOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateMessageForwardOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateMessageForward - Invoke action createForward. Create a draft to forward an existing message, in either JSON or
// MIME format. When using JSON format, you can: - Specify either a comment or the body property of the message
// parameter. Specifying both will return an HTTP 400 Bad Request error. - Specify either the toRecipients parameter or
// the toRecipients property of the message parameter. Specifying both or specifying neither will return an HTTP 400 Bad
// Request error. - Update the draft later to add content to the body or change other message properties. When using
// MIME format: - Provide the applicable Internet message headers and the MIME content, all encoded in base64 format in
// the request body. - Add any attachments and S/MIME properties to the MIME content. Send the draft message in a
// subsequent operation. Alternatively, forward a message in a single operation.
func (c MessageClient) CreateMessageForward(ctx context.Context, id stable.MeMessageId, input CreateMessageForwardRequest, options CreateMessageForwardOperationOptions) (result CreateMessageForwardOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/createForward", id.ID()),
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
	model, err := stable.UnmarshalMessageImplementation(respObj)
	if err != nil {
		return
	}
	result.Model = model

	return
}

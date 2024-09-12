package message

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreatessageOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        stable.Message
}

// Createssage - Create message. Create a draft of a new message in either JSON or MIME format. When using JSON format,
// you can: - Include an attachment to the message. - Update the draft later to add content to the body or change other
// message properties. When using MIME format: - Provide the applicable Internet message headers and the MIME content,
// all encoded in base64 format in the request body. - /* Add any attachments and S/MIME properties to the MIME content.
// By default, this operation saves the draft in the Drafts folder. Send the draft message in a subsequent operation.
// Alternatively, send a new message in a single operation, or create a draft to forward, reply and reply-all to an
// existing message.
func (c MessageClient) Createssage(ctx context.Context, input stable.Message) (result CreatessageOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
		},
		HttpMethod: http.MethodPost,
		Path:       "/me/messages",
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

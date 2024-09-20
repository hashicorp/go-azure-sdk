package mailfoldermessage

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MarkMailFolderMessageAsJunkOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        beta.Message
}

type MarkMailFolderMessageAsJunkOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultMarkMailFolderMessageAsJunkOperationOptions() MarkMailFolderMessageAsJunkOperationOptions {
	return MarkMailFolderMessageAsJunkOperationOptions{}
}

func (o MarkMailFolderMessageAsJunkOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o MarkMailFolderMessageAsJunkOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o MarkMailFolderMessageAsJunkOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// MarkMailFolderMessageAsJunk - Invoke action markAsJunk. Mark a message as junk. This API adds the sender to the list
// of blocked senders and moves the message to the Junk Email folder, when moveToJunk is true.
func (c MailFolderMessageClient) MarkMailFolderMessageAsJunk(ctx context.Context, id beta.MeMailFolderIdMessageId, input MarkMailFolderMessageAsJunkRequest, options MarkMailFolderMessageAsJunkOperationOptions) (result MarkMailFolderMessageAsJunkOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/markAsJunk", id.ID()),
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
	model, err := beta.UnmarshalMessageImplementation(respObj)
	if err != nil {
		return
	}
	result.Model = model

	return
}

package mailfoldermessage

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

type SendMailFolderMessageOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type SendMailFolderMessageOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultSendMailFolderMessageOperationOptions() SendMailFolderMessageOperationOptions {
	return SendMailFolderMessageOperationOptions{}
}

func (o SendMailFolderMessageOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o SendMailFolderMessageOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o SendMailFolderMessageOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// SendMailFolderMessage - Invoke action send. Send an existing draft message. The draft message can be a new message
// draft, reply draft, reply-all draft, or a forward draft. This method saves the message in the Sent Items folder.
// Alternatively, send a new message in a single operation.
func (c MailFolderMessageClient) SendMailFolderMessage(ctx context.Context, id stable.MeMailFolderIdMessageId, options SendMailFolderMessageOperationOptions) (result SendMailFolderMessageOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/send", id.ID()),
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

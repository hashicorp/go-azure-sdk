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

type ForwardMailFolderMessageOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type ForwardMailFolderMessageOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultForwardMailFolderMessageOperationOptions() ForwardMailFolderMessageOperationOptions {
	return ForwardMailFolderMessageOperationOptions{}
}

func (o ForwardMailFolderMessageOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ForwardMailFolderMessageOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o ForwardMailFolderMessageOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// ForwardMailFolderMessage - Invoke action forward. Forward a message using either JSON or MIME format. When using JSON
// format, you can: - Specify either a comment or the body property of the message parameter. Specifying both will
// return an HTTP 400 Bad Request error. - Specify either the toRecipients parameter or the toRecipients property of the
// message parameter. Specifying both or specifying neither will return an HTTP 400 Bad Request error. When using MIME
// format: - Provide the applicable Internet message headers and the MIME content, all encoded in base64 format in the
// request body. - Add any attachments and S/MIME properties to the MIME content. This method saves the message in the
// Sent Items folder. Alternatively, create a draft to forward a message, and send it later.
func (c MailFolderMessageClient) ForwardMailFolderMessage(ctx context.Context, id stable.MeMailFolderIdMessageId, input ForwardMailFolderMessageRequest, options ForwardMailFolderMessageOperationOptions) (result ForwardMailFolderMessageOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/forward", id.ID()),
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

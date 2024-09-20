package mailfolderchildfoldermessage

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

type SendMailFolderChildFolderMessageOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type SendMailFolderChildFolderMessageOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultSendMailFolderChildFolderMessageOperationOptions() SendMailFolderChildFolderMessageOperationOptions {
	return SendMailFolderChildFolderMessageOperationOptions{}
}

func (o SendMailFolderChildFolderMessageOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o SendMailFolderChildFolderMessageOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o SendMailFolderChildFolderMessageOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// SendMailFolderChildFolderMessage - Invoke action send. Send an existing draft message. The draft message can be a new
// message draft, reply draft, reply-all draft, or a forward draft. This method saves the message in the Sent Items
// folder. Alternatively, send a new message in a single operation.
func (c MailFolderChildFolderMessageClient) SendMailFolderChildFolderMessage(ctx context.Context, id beta.UserIdMailFolderIdChildFolderIdMessageId, options SendMailFolderChildFolderMessageOperationOptions) (result SendMailFolderChildFolderMessageOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/send", id.ID()),
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

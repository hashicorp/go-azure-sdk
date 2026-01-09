package mailfolderchildfoldermessage

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateMailFolderChildFolderMessageUnsubscribeOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type CreateMailFolderChildFolderMessageUnsubscribeOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateMailFolderChildFolderMessageUnsubscribeOperationOptions() CreateMailFolderChildFolderMessageUnsubscribeOperationOptions {
	return CreateMailFolderChildFolderMessageUnsubscribeOperationOptions{}
}

func (o CreateMailFolderChildFolderMessageUnsubscribeOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateMailFolderChildFolderMessageUnsubscribeOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateMailFolderChildFolderMessageUnsubscribeOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateMailFolderChildFolderMessageUnsubscribe - Invoke action unsubscribe. Submits an email request on behalf of the
// signed-in user to unsubscribe from an email distribution list. Uses the information in the List-Unsubscribe header.
// Message senders can use mailing lists in a user-friendly way by including an option for recipients to opt out. They
// can do so by specifying the List-Unsubscribe header in each message following RFC-2369. Note In particular, for the
// unsubscribe action to work, the sender must specify mailto: and not URL-based unsubscribe information. Setting that
// header would also set the unsubscribeEnabled property of the message instance to true, and the unsubscribeData
// property to the header data. If the unsubscribeEnabled property of a message is true, you can use the unsubscribe
// action to unsubscribe the user from similar future messages as managed by the message sender. A successful
// unsubscribe action moves the message to the Deleted Items folder. The actual exclusion of the user from future mail
// distribution is managed by the sender.
func (c MailFolderChildFolderMessageClient) CreateMailFolderChildFolderMessageUnsubscribe(ctx context.Context, id beta.MeMailFolderIdChildFolderIdMessageId, options CreateMailFolderChildFolderMessageUnsubscribeOperationOptions) (result CreateMailFolderChildFolderMessageUnsubscribeOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/unsubscribe", id.ID()),
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

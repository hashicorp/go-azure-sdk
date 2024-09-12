package mailfolderchildfoldermessage

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

type ReplyMailFolderChildFolderMessageOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

// ReplyMailFolderChildFolderMessage - Invoke action reply. Reply to the sender of a message using either JSON or MIME
// format. When using JSON format: * Specify either a comment or the body property of the message parameter. Specifying
// both will return an HTTP 400 Bad Request error. * If the original message specifies a recipient in the replyTo
// property, per Internet Message Format (RFC 2822), send the reply to the recipients in replyTo and not the recipient
// in the from property. When using MIME format: - Provide the applicable Internet message headers and the MIME content,
// all encoded in base64 format in the request body. - Add any attachments and S/MIME properties to the MIME content.
// This method saves the message in the Sent Items folder. Alternatively, create a draft to reply to an existing message
// and send it later.
func (c MailFolderChildFolderMessageClient) ReplyMailFolderChildFolderMessage(ctx context.Context, id stable.MeMailFolderIdChildFolderIdMessageId, input ReplyMailFolderChildFolderMessageRequest) (result ReplyMailFolderChildFolderMessageOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod: http.MethodPost,
		Path:       fmt.Sprintf("%s/reply", id.ID()),
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

package mailfoldermessagemention

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetMailFolderMessageMentionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.Mention
}

type GetMailFolderMessageMentionOperationOptions struct {
	Expand *odata.Expand
	Select *[]string
}

func DefaultGetMailFolderMessageMentionOperationOptions() GetMailFolderMessageMentionOperationOptions {
	return GetMailFolderMessageMentionOperationOptions{}
}

func (o GetMailFolderMessageMentionOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetMailFolderMessageMentionOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetMailFolderMessageMentionOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetMailFolderMessageMention - Get mentions from users. A collection of mentions in the message, ordered by the
// createdDateTime from the newest to the oldest. By default, a GET /messages does not return this property unless you
// apply $expand on the property.
func (c MailFolderMessageMentionClient) GetMailFolderMessageMention(ctx context.Context, id beta.UserIdMailFolderIdMessageIdMentionId, options GetMailFolderMessageMentionOperationOptions) (result GetMailFolderMessageMentionOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          id.ID(),
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

	var model beta.Mention
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}

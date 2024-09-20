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

type SetMailFolderChildFolderMessageValueOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type SetMailFolderChildFolderMessageValueOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultSetMailFolderChildFolderMessageValueOperationOptions() SetMailFolderChildFolderMessageValueOperationOptions {
	return SetMailFolderChildFolderMessageValueOperationOptions{}
}

func (o SetMailFolderChildFolderMessageValueOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o SetMailFolderChildFolderMessageValueOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o SetMailFolderChildFolderMessageValueOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// SetMailFolderChildFolderMessageValue - Update media content for the navigation property messages in users. The unique
// identifier for an entity. Read-only.
func (c MailFolderChildFolderMessageClient) SetMailFolderChildFolderMessageValue(ctx context.Context, id stable.UserIdMailFolderIdChildFolderIdMessageId, input []byte, options SetMailFolderChildFolderMessageValueOperationOptions) (result SetMailFolderChildFolderMessageValueOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPut,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/$value", id.ID()),
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

package mailfolder

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

type GetMailFolderOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        stable.MailFolder
}

type GetMailFolderOperationOptions struct {
	Expand               *odata.Expand
	IncludeHiddenFolders *string
	Select               *[]string
}

func DefaultGetMailFolderOperationOptions() GetMailFolderOperationOptions {
	return GetMailFolderOperationOptions{}
}

func (o GetMailFolderOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetMailFolderOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetMailFolderOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.IncludeHiddenFolders != nil {
		out.Append("includeHiddenFolders", fmt.Sprintf("%v", *o.IncludeHiddenFolders))
	}
	return &out
}

// GetMailFolder - Get mailFolders from users. The user's mail folders. Read-only. Nullable.
func (c MailFolderClient) GetMailFolder(ctx context.Context, id stable.UserIdMailFolderId, options GetMailFolderOperationOptions) (result GetMailFolderOperationResponse, err error) {
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

	var respObj json.RawMessage
	if err = resp.Unmarshal(&respObj); err != nil {
		return
	}
	model, err := stable.UnmarshalMailFolderImplementation(respObj)
	if err != nil {
		return
	}
	result.Model = model

	return
}

package mailfolderchildfolder

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

type GetMailFolderChildFolderOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        beta.MailFolder
}

type GetMailFolderChildFolderOperationOptions struct {
	Expand               *odata.Expand
	IncludeHiddenFolders *string
	Metadata             *odata.Metadata
	Select               *[]string
}

func DefaultGetMailFolderChildFolderOperationOptions() GetMailFolderChildFolderOperationOptions {
	return GetMailFolderChildFolderOperationOptions{}
}

func (o GetMailFolderChildFolderOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetMailFolderChildFolderOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetMailFolderChildFolderOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.IncludeHiddenFolders != nil {
		out.Append("includeHiddenFolders", fmt.Sprintf("%v", *o.IncludeHiddenFolders))
	}
	return &out
}

// GetMailFolderChildFolder - Get childFolders from users. The collection of child folders in the mailFolder.
func (c MailFolderChildFolderClient) GetMailFolderChildFolder(ctx context.Context, id beta.UserIdMailFolderIdChildFolderId, options GetMailFolderChildFolderOperationOptions) (result GetMailFolderChildFolderOperationResponse, err error) {
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
	model, err := beta.UnmarshalMailFolderImplementation(respObj)
	if err != nil {
		return
	}
	result.Model = model

	return
}

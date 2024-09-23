package mailfolderchildfolder

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

type MoveMailFolderChildFolderOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        stable.MailFolder
}

type MoveMailFolderChildFolderOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultMoveMailFolderChildFolderOperationOptions() MoveMailFolderChildFolderOperationOptions {
	return MoveMailFolderChildFolderOperationOptions{}
}

func (o MoveMailFolderChildFolderOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o MoveMailFolderChildFolderOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o MoveMailFolderChildFolderOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// MoveMailFolderChildFolder - Invoke action move. Move a mailfolder and its contents to another mailfolder.
func (c MailFolderChildFolderClient) MoveMailFolderChildFolder(ctx context.Context, id stable.MeMailFolderIdChildFolderId, input MoveMailFolderChildFolderRequest, options MoveMailFolderChildFolderOperationOptions) (result MoveMailFolderChildFolderOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/move", id.ID()),
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

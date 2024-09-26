package contactfolderchildfolder

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

type CreateContactFolderChildFolderOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.ContactFolder
}

type CreateContactFolderChildFolderOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateContactFolderChildFolderOperationOptions() CreateContactFolderChildFolderOperationOptions {
	return CreateContactFolderChildFolderOperationOptions{}
}

func (o CreateContactFolderChildFolderOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateContactFolderChildFolderOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateContactFolderChildFolderOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateContactFolderChildFolder - Create ContactFolder. Create a new contactFolder as a child of a specified folder.
// You can also create a new contactFolder under the user's default contact folder.
func (c ContactFolderChildFolderClient) CreateContactFolderChildFolder(ctx context.Context, id stable.MeContactFolderId, input stable.ContactFolder, options CreateContactFolderChildFolderOperationOptions) (result CreateContactFolderChildFolderOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/childFolders", id.ID()),
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

	var model stable.ContactFolder
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}

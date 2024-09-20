package onenotepage

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateOnenotePageOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.OnenotePage
}

type CreateOnenotePageOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultCreateOnenotePageOperationOptions() CreateOnenotePageOperationOptions {
	return CreateOnenotePageOperationOptions{}
}

func (o CreateOnenotePageOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateOnenotePageOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateOnenotePageOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateOnenotePage - Create onenotePage. Create a new OneNote page in the default section of the default notebook. To
// create a page in a different section in the default notebook, you can use the sectionName query parameter. Example:
// ../onenote/pages?sectionName=My%20section The POST /onenote/pages operation is used only to create pages in the
// current user's default notebook. If you're targeting other notebooks, you can create pages in a specified section.
func (c OnenotePageClient) CreateOnenotePage(ctx context.Context, input stable.OnenotePage, options CreateOnenotePageOperationOptions) (result CreateOnenotePageOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          "/me/onenote/pages",
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

	var model stable.OnenotePage
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}

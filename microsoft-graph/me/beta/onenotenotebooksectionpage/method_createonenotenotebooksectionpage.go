package onenotenotebooksectionpage

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

type CreateOnenoteNotebookSectionPageOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.OnenotePage
}

type CreateOnenoteNotebookSectionPageOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultCreateOnenoteNotebookSectionPageOperationOptions() CreateOnenoteNotebookSectionPageOperationOptions {
	return CreateOnenoteNotebookSectionPageOperationOptions{}
}

func (o CreateOnenoteNotebookSectionPageOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateOnenoteNotebookSectionPageOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateOnenoteNotebookSectionPageOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateOnenoteNotebookSectionPage - Create new navigation property to pages for me
func (c OnenoteNotebookSectionPageClient) CreateOnenoteNotebookSectionPage(ctx context.Context, id beta.MeOnenoteNotebookIdSectionId, input beta.OnenotePage, options CreateOnenoteNotebookSectionPageOperationOptions) (result CreateOnenoteNotebookSectionPageOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/pages", id.ID()),
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

	var model beta.OnenotePage
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}

package onenotenotebook

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

type GetOnenoteNotebooksFromWebUrlOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.CopyNotebookModel
}

type GetOnenoteNotebooksFromWebUrlOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultGetOnenoteNotebooksFromWebUrlOperationOptions() GetOnenoteNotebooksFromWebUrlOperationOptions {
	return GetOnenoteNotebooksFromWebUrlOperationOptions{}
}

func (o GetOnenoteNotebooksFromWebUrlOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetOnenoteNotebooksFromWebUrlOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o GetOnenoteNotebooksFromWebUrlOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetOnenoteNotebooksFromWebUrl - Invoke action getNotebookFromWebUrl. Retrieve the properties and relationships of a
// notebook object by using its URL path. The location can be user notebooks on Microsoft 365, group notebooks, or
// SharePoint site-hosted team notebooks on Microsoft 365.
func (c OnenoteNotebookClient) GetOnenoteNotebooksFromWebUrl(ctx context.Context, id beta.GroupId, input GetOnenoteNotebooksFromWebUrlRequest, options GetOnenoteNotebooksFromWebUrlOperationOptions) (result GetOnenoteNotebooksFromWebUrlOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/onenote/notebooks/getNotebookFromWebUrl", id.ID()),
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

	var model beta.CopyNotebookModel
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}

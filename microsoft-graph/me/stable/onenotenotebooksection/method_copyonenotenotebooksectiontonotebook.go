package onenotenotebooksection

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

type CopyOnenoteNotebookSectionToNotebookOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.OnenoteOperation
}

type CopyOnenoteNotebookSectionToNotebookOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCopyOnenoteNotebookSectionToNotebookOperationOptions() CopyOnenoteNotebookSectionToNotebookOperationOptions {
	return CopyOnenoteNotebookSectionToNotebookOperationOptions{}
}

func (o CopyOnenoteNotebookSectionToNotebookOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CopyOnenoteNotebookSectionToNotebookOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CopyOnenoteNotebookSectionToNotebookOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CopyOnenoteNotebookSectionToNotebook - Invoke action copyToNotebook. For Copy operations, you follow an asynchronous
// calling pattern: First call the Copy action, and then poll the operation endpoint for the result.
func (c OnenoteNotebookSectionClient) CopyOnenoteNotebookSectionToNotebook(ctx context.Context, id stable.MeOnenoteNotebookIdSectionId, input CopyOnenoteNotebookSectionToNotebookRequest, options CopyOnenoteNotebookSectionToNotebookOperationOptions) (result CopyOnenoteNotebookSectionToNotebookOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/copyToNotebook", id.ID()),
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

	var model stable.OnenoteOperation
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}

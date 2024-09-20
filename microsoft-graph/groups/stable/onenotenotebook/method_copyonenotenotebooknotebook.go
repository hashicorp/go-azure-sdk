package onenotenotebook

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

type CopyOnenoteNotebookNotebookOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.OnenoteOperation
}

type CopyOnenoteNotebookNotebookOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultCopyOnenoteNotebookNotebookOperationOptions() CopyOnenoteNotebookNotebookOperationOptions {
	return CopyOnenoteNotebookNotebookOperationOptions{}
}

func (o CopyOnenoteNotebookNotebookOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CopyOnenoteNotebookNotebookOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CopyOnenoteNotebookNotebookOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CopyOnenoteNotebookNotebook - Invoke action copyNotebook. For Copy operations, you follow an asynchronous calling
// pattern: First call the Copy action, and then poll the operation endpoint for the result.
func (c OnenoteNotebookClient) CopyOnenoteNotebookNotebook(ctx context.Context, id stable.GroupIdOnenoteNotebookId, input CopyOnenoteNotebookNotebookRequest, options CopyOnenoteNotebookNotebookOperationOptions) (result CopyOnenoteNotebookNotebookOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/copyNotebook", id.ID()),
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

package notebooks

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/client/dataplane"
	"github.com/hashicorp/go-azure-sdk/sdk/client/pollers"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NotebookCreateOrUpdateNotebookOperationResponse struct {
	Poller       pollers.Poller
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *NotebookResource
}

type NotebookCreateOrUpdateNotebookOperationOptions struct {
	IfMatch *string
}

func DefaultNotebookCreateOrUpdateNotebookOperationOptions() NotebookCreateOrUpdateNotebookOperationOptions {
	return NotebookCreateOrUpdateNotebookOperationOptions{}
}

func (o NotebookCreateOrUpdateNotebookOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o NotebookCreateOrUpdateNotebookOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o NotebookCreateOrUpdateNotebookOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// NotebookCreateOrUpdateNotebook ...
func (c NotebooksClient) NotebookCreateOrUpdateNotebook(ctx context.Context, id NotebookId, input NotebookResource, options NotebookCreateOrUpdateNotebookOperationOptions) (result NotebookCreateOrUpdateNotebookOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPut,
		OptionsObject: options,
		Path:          id.Path(),
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

	result.Poller, err = dataplane.PollerFromResponse(resp, c.Client)
	if err != nil {
		return
	}

	return
}

// NotebookCreateOrUpdateNotebookThenPoll performs NotebookCreateOrUpdateNotebook then polls until it's completed
func (c NotebooksClient) NotebookCreateOrUpdateNotebookThenPoll(ctx context.Context, id NotebookId, input NotebookResource, options NotebookCreateOrUpdateNotebookOperationOptions) error {
	result, err := c.NotebookCreateOrUpdateNotebook(ctx, id, input, options)
	if err != nil {
		return fmt.Errorf("performing NotebookCreateOrUpdateNotebook: %+v", err)
	}

	if err := result.Poller.PollUntilDone(ctx); err != nil {
		return fmt.Errorf("polling after NotebookCreateOrUpdateNotebook: %+v", err)
	}

	return nil
}

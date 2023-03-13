package notebookworkspacesresource

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/hashicorp/go-azure-helpers/polling"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NotebookWorkspacesDeleteOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// NotebookWorkspacesDelete ...
func (c NotebookWorkspacesResourceClient) NotebookWorkspacesDelete(ctx context.Context, id DatabaseAccountId) (result NotebookWorkspacesDeleteOperationResponse, err error) {
	req, err := c.preparerForNotebookWorkspacesDelete(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "notebookworkspacesresource.NotebookWorkspacesResourceClient", "NotebookWorkspacesDelete", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForNotebookWorkspacesDelete(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "notebookworkspacesresource.NotebookWorkspacesResourceClient", "NotebookWorkspacesDelete", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// NotebookWorkspacesDeleteThenPoll performs NotebookWorkspacesDelete then polls until it's completed
func (c NotebookWorkspacesResourceClient) NotebookWorkspacesDeleteThenPoll(ctx context.Context, id DatabaseAccountId) error {
	result, err := c.NotebookWorkspacesDelete(ctx, id)
	if err != nil {
		return fmt.Errorf("performing NotebookWorkspacesDelete: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after NotebookWorkspacesDelete: %+v", err)
	}

	return nil
}

// preparerForNotebookWorkspacesDelete prepares the NotebookWorkspacesDelete request.
func (c NotebookWorkspacesResourceClient) preparerForNotebookWorkspacesDelete(ctx context.Context, id DatabaseAccountId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsDelete(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/notebookWorkspaces/default", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForNotebookWorkspacesDelete sends the NotebookWorkspacesDelete request. The method will close the
// http.Response Body if it receives an error.
func (c NotebookWorkspacesResourceClient) senderForNotebookWorkspacesDelete(ctx context.Context, req *http.Request) (future NotebookWorkspacesDeleteOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}

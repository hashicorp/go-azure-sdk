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

type NotebookWorkspacesRegenerateAuthTokenOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// NotebookWorkspacesRegenerateAuthToken ...
func (c NotebookWorkspacesResourceClient) NotebookWorkspacesRegenerateAuthToken(ctx context.Context, id DatabaseAccountId) (result NotebookWorkspacesRegenerateAuthTokenOperationResponse, err error) {
	req, err := c.preparerForNotebookWorkspacesRegenerateAuthToken(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "notebookworkspacesresource.NotebookWorkspacesResourceClient", "NotebookWorkspacesRegenerateAuthToken", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForNotebookWorkspacesRegenerateAuthToken(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "notebookworkspacesresource.NotebookWorkspacesResourceClient", "NotebookWorkspacesRegenerateAuthToken", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// NotebookWorkspacesRegenerateAuthTokenThenPoll performs NotebookWorkspacesRegenerateAuthToken then polls until it's completed
func (c NotebookWorkspacesResourceClient) NotebookWorkspacesRegenerateAuthTokenThenPoll(ctx context.Context, id DatabaseAccountId) error {
	result, err := c.NotebookWorkspacesRegenerateAuthToken(ctx, id)
	if err != nil {
		return fmt.Errorf("performing NotebookWorkspacesRegenerateAuthToken: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after NotebookWorkspacesRegenerateAuthToken: %+v", err)
	}

	return nil
}

// preparerForNotebookWorkspacesRegenerateAuthToken prepares the NotebookWorkspacesRegenerateAuthToken request.
func (c NotebookWorkspacesResourceClient) preparerForNotebookWorkspacesRegenerateAuthToken(ctx context.Context, id DatabaseAccountId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/notebookWorkspaces/default/regenerateAuthToken", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForNotebookWorkspacesRegenerateAuthToken sends the NotebookWorkspacesRegenerateAuthToken request. The method will close the
// http.Response Body if it receives an error.
func (c NotebookWorkspacesResourceClient) senderForNotebookWorkspacesRegenerateAuthToken(ctx context.Context, req *http.Request) (future NotebookWorkspacesRegenerateAuthTokenOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}

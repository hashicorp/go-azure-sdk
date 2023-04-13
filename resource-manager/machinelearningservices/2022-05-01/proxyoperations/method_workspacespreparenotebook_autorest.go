package proxyoperations

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

type WorkspacesPrepareNotebookOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// WorkspacesPrepareNotebook ...
func (c ProxyOperationsClient) WorkspacesPrepareNotebook(ctx context.Context, id WorkspaceId) (result WorkspacesPrepareNotebookOperationResponse, err error) {
	req, err := c.preparerForWorkspacesPrepareNotebook(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "proxyoperations.ProxyOperationsClient", "WorkspacesPrepareNotebook", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForWorkspacesPrepareNotebook(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "proxyoperations.ProxyOperationsClient", "WorkspacesPrepareNotebook", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// WorkspacesPrepareNotebookThenPoll performs WorkspacesPrepareNotebook then polls until it's completed
func (c ProxyOperationsClient) WorkspacesPrepareNotebookThenPoll(ctx context.Context, id WorkspaceId) error {
	result, err := c.WorkspacesPrepareNotebook(ctx, id)
	if err != nil {
		return fmt.Errorf("performing WorkspacesPrepareNotebook: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after WorkspacesPrepareNotebook: %+v", err)
	}

	return nil
}

// preparerForWorkspacesPrepareNotebook prepares the WorkspacesPrepareNotebook request.
func (c ProxyOperationsClient) preparerForWorkspacesPrepareNotebook(ctx context.Context, id WorkspaceId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/prepareNotebook", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForWorkspacesPrepareNotebook sends the WorkspacesPrepareNotebook request. The method will close the
// http.Response Body if it receives an error.
func (c ProxyOperationsClient) senderForWorkspacesPrepareNotebook(ctx context.Context, req *http.Request) (future WorkspacesPrepareNotebookOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}

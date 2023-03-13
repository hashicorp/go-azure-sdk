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

type NotebookWorkspacesCreateOrUpdateOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// NotebookWorkspacesCreateOrUpdate ...
func (c NotebookWorkspacesResourceClient) NotebookWorkspacesCreateOrUpdate(ctx context.Context, id DatabaseAccountId, input ARMProxyResource) (result NotebookWorkspacesCreateOrUpdateOperationResponse, err error) {
	req, err := c.preparerForNotebookWorkspacesCreateOrUpdate(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "notebookworkspacesresource.NotebookWorkspacesResourceClient", "NotebookWorkspacesCreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForNotebookWorkspacesCreateOrUpdate(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "notebookworkspacesresource.NotebookWorkspacesResourceClient", "NotebookWorkspacesCreateOrUpdate", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// NotebookWorkspacesCreateOrUpdateThenPoll performs NotebookWorkspacesCreateOrUpdate then polls until it's completed
func (c NotebookWorkspacesResourceClient) NotebookWorkspacesCreateOrUpdateThenPoll(ctx context.Context, id DatabaseAccountId, input ARMProxyResource) error {
	result, err := c.NotebookWorkspacesCreateOrUpdate(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing NotebookWorkspacesCreateOrUpdate: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after NotebookWorkspacesCreateOrUpdate: %+v", err)
	}

	return nil
}

// preparerForNotebookWorkspacesCreateOrUpdate prepares the NotebookWorkspacesCreateOrUpdate request.
func (c NotebookWorkspacesResourceClient) preparerForNotebookWorkspacesCreateOrUpdate(ctx context.Context, id DatabaseAccountId, input ARMProxyResource) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/notebookWorkspaces/default", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForNotebookWorkspacesCreateOrUpdate sends the NotebookWorkspacesCreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (c NotebookWorkspacesResourceClient) senderForNotebookWorkspacesCreateOrUpdate(ctx context.Context, req *http.Request) (future NotebookWorkspacesCreateOrUpdateOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}

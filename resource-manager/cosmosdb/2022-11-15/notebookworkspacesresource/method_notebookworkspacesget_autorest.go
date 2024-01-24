package notebookworkspacesresource

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NotebookWorkspacesGetOperationResponse struct {
	HttpResponse *http.Response
	Model        *NotebookWorkspace
}

// NotebookWorkspacesGet ...
func (c NotebookWorkspacesResourceClient) NotebookWorkspacesGet(ctx context.Context, id DatabaseAccountId) (result NotebookWorkspacesGetOperationResponse, err error) {
	req, err := c.preparerForNotebookWorkspacesGet(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "notebookworkspacesresource.NotebookWorkspacesResourceClient", "NotebookWorkspacesGet", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "notebookworkspacesresource.NotebookWorkspacesResourceClient", "NotebookWorkspacesGet", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForNotebookWorkspacesGet(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "notebookworkspacesresource.NotebookWorkspacesResourceClient", "NotebookWorkspacesGet", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForNotebookWorkspacesGet prepares the NotebookWorkspacesGet request.
func (c NotebookWorkspacesResourceClient) preparerForNotebookWorkspacesGet(ctx context.Context, id DatabaseAccountId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/notebookWorkspaces/default", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForNotebookWorkspacesGet handles the response to the NotebookWorkspacesGet request. The method always
// closes the http.Response Body.
func (c NotebookWorkspacesResourceClient) responderForNotebookWorkspacesGet(resp *http.Response) (result NotebookWorkspacesGetOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

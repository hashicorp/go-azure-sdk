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

type NotebookWorkspacesListByDatabaseAccountOperationResponse struct {
	HttpResponse *http.Response
	Model        *NotebookWorkspaceListResult
}

// NotebookWorkspacesListByDatabaseAccount ...
func (c NotebookWorkspacesResourceClient) NotebookWorkspacesListByDatabaseAccount(ctx context.Context, id DatabaseAccountId) (result NotebookWorkspacesListByDatabaseAccountOperationResponse, err error) {
	req, err := c.preparerForNotebookWorkspacesListByDatabaseAccount(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "notebookworkspacesresource.NotebookWorkspacesResourceClient", "NotebookWorkspacesListByDatabaseAccount", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "notebookworkspacesresource.NotebookWorkspacesResourceClient", "NotebookWorkspacesListByDatabaseAccount", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForNotebookWorkspacesListByDatabaseAccount(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "notebookworkspacesresource.NotebookWorkspacesResourceClient", "NotebookWorkspacesListByDatabaseAccount", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForNotebookWorkspacesListByDatabaseAccount prepares the NotebookWorkspacesListByDatabaseAccount request.
func (c NotebookWorkspacesResourceClient) preparerForNotebookWorkspacesListByDatabaseAccount(ctx context.Context, id DatabaseAccountId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/notebookWorkspaces", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForNotebookWorkspacesListByDatabaseAccount handles the response to the NotebookWorkspacesListByDatabaseAccount request. The method always
// closes the http.Response Body.
func (c NotebookWorkspacesResourceClient) responderForNotebookWorkspacesListByDatabaseAccount(resp *http.Response) (result NotebookWorkspacesListByDatabaseAccountOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

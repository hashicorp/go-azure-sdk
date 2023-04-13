package proxyoperations

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkspacesListNotebookKeysOperationResponse struct {
	HttpResponse *http.Response
	Model        *ListNotebookKeysResult
}

// WorkspacesListNotebookKeys ...
func (c ProxyOperationsClient) WorkspacesListNotebookKeys(ctx context.Context, id WorkspaceId) (result WorkspacesListNotebookKeysOperationResponse, err error) {
	req, err := c.preparerForWorkspacesListNotebookKeys(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "proxyoperations.ProxyOperationsClient", "WorkspacesListNotebookKeys", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "proxyoperations.ProxyOperationsClient", "WorkspacesListNotebookKeys", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForWorkspacesListNotebookKeys(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "proxyoperations.ProxyOperationsClient", "WorkspacesListNotebookKeys", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForWorkspacesListNotebookKeys prepares the WorkspacesListNotebookKeys request.
func (c ProxyOperationsClient) preparerForWorkspacesListNotebookKeys(ctx context.Context, id WorkspaceId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/listNotebookKeys", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForWorkspacesListNotebookKeys handles the response to the WorkspacesListNotebookKeys request. The method always
// closes the http.Response Body.
func (c ProxyOperationsClient) responderForWorkspacesListNotebookKeys(resp *http.Response) (result WorkspacesListNotebookKeysOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

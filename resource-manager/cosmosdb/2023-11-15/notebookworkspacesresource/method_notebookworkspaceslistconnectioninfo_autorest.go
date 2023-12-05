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

type NotebookWorkspacesListConnectionInfoOperationResponse struct {
	HttpResponse *http.Response
	Model        *NotebookWorkspaceConnectionInfoResult
}

// NotebookWorkspacesListConnectionInfo ...
func (c NotebookWorkspacesResourceClient) NotebookWorkspacesListConnectionInfo(ctx context.Context, id DatabaseAccountId) (result NotebookWorkspacesListConnectionInfoOperationResponse, err error) {
	req, err := c.preparerForNotebookWorkspacesListConnectionInfo(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "notebookworkspacesresource.NotebookWorkspacesResourceClient", "NotebookWorkspacesListConnectionInfo", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "notebookworkspacesresource.NotebookWorkspacesResourceClient", "NotebookWorkspacesListConnectionInfo", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForNotebookWorkspacesListConnectionInfo(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "notebookworkspacesresource.NotebookWorkspacesResourceClient", "NotebookWorkspacesListConnectionInfo", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForNotebookWorkspacesListConnectionInfo prepares the NotebookWorkspacesListConnectionInfo request.
func (c NotebookWorkspacesResourceClient) preparerForNotebookWorkspacesListConnectionInfo(ctx context.Context, id DatabaseAccountId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/notebookWorkspaces/default/listConnectionInfo", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForNotebookWorkspacesListConnectionInfo handles the response to the NotebookWorkspacesListConnectionInfo request. The method always
// closes the http.Response Body.
func (c NotebookWorkspacesResourceClient) responderForNotebookWorkspacesListConnectionInfo(resp *http.Response) (result NotebookWorkspacesListConnectionInfoOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

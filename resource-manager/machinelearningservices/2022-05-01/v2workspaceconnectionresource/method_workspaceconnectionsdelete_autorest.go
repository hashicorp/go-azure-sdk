package v2workspaceconnectionresource

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkspaceConnectionsDeleteOperationResponse struct {
	HttpResponse *http.Response
}

// WorkspaceConnectionsDelete ...
func (c V2WorkspaceConnectionResourceClient) WorkspaceConnectionsDelete(ctx context.Context, id ConnectionId) (result WorkspaceConnectionsDeleteOperationResponse, err error) {
	req, err := c.preparerForWorkspaceConnectionsDelete(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "v2workspaceconnectionresource.V2WorkspaceConnectionResourceClient", "WorkspaceConnectionsDelete", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "v2workspaceconnectionresource.V2WorkspaceConnectionResourceClient", "WorkspaceConnectionsDelete", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForWorkspaceConnectionsDelete(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "v2workspaceconnectionresource.V2WorkspaceConnectionResourceClient", "WorkspaceConnectionsDelete", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForWorkspaceConnectionsDelete prepares the WorkspaceConnectionsDelete request.
func (c V2WorkspaceConnectionResourceClient) preparerForWorkspaceConnectionsDelete(ctx context.Context, id ConnectionId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsDelete(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForWorkspaceConnectionsDelete handles the response to the WorkspaceConnectionsDelete request. The method always
// closes the http.Response Body.
func (c V2WorkspaceConnectionResourceClient) responderForWorkspaceConnectionsDelete(resp *http.Response) (result WorkspaceConnectionsDeleteOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusNoContent, http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

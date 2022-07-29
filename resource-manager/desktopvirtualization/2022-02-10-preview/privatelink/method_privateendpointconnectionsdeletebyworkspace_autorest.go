package privatelink

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivateEndpointConnectionsDeleteByWorkspaceOperationResponse struct {
	HttpResponse *http.Response
}

// PrivateEndpointConnectionsDeleteByWorkspace ...
func (c PrivateLinkClient) PrivateEndpointConnectionsDeleteByWorkspace(ctx context.Context, id WorkspacePrivateEndpointConnectionId) (result PrivateEndpointConnectionsDeleteByWorkspaceOperationResponse, err error) {
	req, err := c.preparerForPrivateEndpointConnectionsDeleteByWorkspace(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "privatelink.PrivateLinkClient", "PrivateEndpointConnectionsDeleteByWorkspace", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "privatelink.PrivateLinkClient", "PrivateEndpointConnectionsDeleteByWorkspace", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForPrivateEndpointConnectionsDeleteByWorkspace(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "privatelink.PrivateLinkClient", "PrivateEndpointConnectionsDeleteByWorkspace", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForPrivateEndpointConnectionsDeleteByWorkspace prepares the PrivateEndpointConnectionsDeleteByWorkspace request.
func (c PrivateLinkClient) preparerForPrivateEndpointConnectionsDeleteByWorkspace(ctx context.Context, id WorkspacePrivateEndpointConnectionId) (*http.Request, error) {
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

// responderForPrivateEndpointConnectionsDeleteByWorkspace handles the response to the PrivateEndpointConnectionsDeleteByWorkspace request. The method always
// closes the http.Response Body.
func (c PrivateLinkClient) responderForPrivateEndpointConnectionsDeleteByWorkspace(resp *http.Response) (result PrivateEndpointConnectionsDeleteByWorkspaceOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusNoContent, http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

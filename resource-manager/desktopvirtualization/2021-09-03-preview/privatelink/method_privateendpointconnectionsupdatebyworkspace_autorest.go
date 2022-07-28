package privatelink

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivateEndpointConnectionsUpdateByWorkspaceOperationResponse struct {
	HttpResponse *http.Response
	Model        *PrivateEndpointConnectionWithSystemData
}

// PrivateEndpointConnectionsUpdateByWorkspace ...
func (c PrivateLinkClient) PrivateEndpointConnectionsUpdateByWorkspace(ctx context.Context, id WorkspacePrivateEndpointConnectionId, input PrivateEndpointConnection) (result PrivateEndpointConnectionsUpdateByWorkspaceOperationResponse, err error) {
	req, err := c.preparerForPrivateEndpointConnectionsUpdateByWorkspace(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "privatelink.PrivateLinkClient", "PrivateEndpointConnectionsUpdateByWorkspace", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "privatelink.PrivateLinkClient", "PrivateEndpointConnectionsUpdateByWorkspace", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForPrivateEndpointConnectionsUpdateByWorkspace(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "privatelink.PrivateLinkClient", "PrivateEndpointConnectionsUpdateByWorkspace", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForPrivateEndpointConnectionsUpdateByWorkspace prepares the PrivateEndpointConnectionsUpdateByWorkspace request.
func (c PrivateLinkClient) preparerForPrivateEndpointConnectionsUpdateByWorkspace(ctx context.Context, id WorkspacePrivateEndpointConnectionId, input PrivateEndpointConnection) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForPrivateEndpointConnectionsUpdateByWorkspace handles the response to the PrivateEndpointConnectionsUpdateByWorkspace request. The method always
// closes the http.Response Body.
func (c PrivateLinkClient) responderForPrivateEndpointConnectionsUpdateByWorkspace(resp *http.Response) (result PrivateEndpointConnectionsUpdateByWorkspaceOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

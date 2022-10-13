package privateendpointconnections

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivateEndpointConnectionProxiesGetOperationResponse struct {
	HttpResponse *http.Response
	Model        *PrivateEndpointConnectionProxy
}

// PrivateEndpointConnectionProxiesGet ...
func (c PrivateEndpointConnectionsClient) PrivateEndpointConnectionProxiesGet(ctx context.Context, id PrivateEndpointConnectionProxyId) (result PrivateEndpointConnectionProxiesGetOperationResponse, err error) {
	req, err := c.preparerForPrivateEndpointConnectionProxiesGet(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "privateendpointconnections.PrivateEndpointConnectionsClient", "PrivateEndpointConnectionProxiesGet", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "privateendpointconnections.PrivateEndpointConnectionsClient", "PrivateEndpointConnectionProxiesGet", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForPrivateEndpointConnectionProxiesGet(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "privateendpointconnections.PrivateEndpointConnectionsClient", "PrivateEndpointConnectionProxiesGet", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForPrivateEndpointConnectionProxiesGet prepares the PrivateEndpointConnectionProxiesGet request.
func (c PrivateEndpointConnectionsClient) preparerForPrivateEndpointConnectionProxiesGet(ctx context.Context, id PrivateEndpointConnectionProxyId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForPrivateEndpointConnectionProxiesGet handles the response to the PrivateEndpointConnectionProxiesGet request. The method always
// closes the http.Response Body.
func (c PrivateEndpointConnectionsClient) responderForPrivateEndpointConnectionProxiesGet(resp *http.Response) (result PrivateEndpointConnectionProxiesGetOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

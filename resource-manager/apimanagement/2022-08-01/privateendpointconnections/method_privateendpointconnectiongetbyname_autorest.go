package privateendpointconnections

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivateEndpointConnectionGetByNameOperationResponse struct {
	HttpResponse *http.Response
	Model        *PrivateEndpointConnection
}

// PrivateEndpointConnectionGetByName ...
func (c PrivateEndpointConnectionsClient) PrivateEndpointConnectionGetByName(ctx context.Context, id PrivateEndpointConnectionId) (result PrivateEndpointConnectionGetByNameOperationResponse, err error) {
	req, err := c.preparerForPrivateEndpointConnectionGetByName(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "privateendpointconnections.PrivateEndpointConnectionsClient", "PrivateEndpointConnectionGetByName", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "privateendpointconnections.PrivateEndpointConnectionsClient", "PrivateEndpointConnectionGetByName", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForPrivateEndpointConnectionGetByName(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "privateendpointconnections.PrivateEndpointConnectionsClient", "PrivateEndpointConnectionGetByName", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForPrivateEndpointConnectionGetByName prepares the PrivateEndpointConnectionGetByName request.
func (c PrivateEndpointConnectionsClient) preparerForPrivateEndpointConnectionGetByName(ctx context.Context, id PrivateEndpointConnectionId) (*http.Request, error) {
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

// responderForPrivateEndpointConnectionGetByName handles the response to the PrivateEndpointConnectionGetByName request. The method always
// closes the http.Response Body.
func (c PrivateEndpointConnectionsClient) responderForPrivateEndpointConnectionGetByName(resp *http.Response) (result PrivateEndpointConnectionGetByNameOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

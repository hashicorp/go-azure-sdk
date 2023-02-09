package privateendpointconnectionresource

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivateEndpointConnectionsListByStorageSyncServiceOperationResponse struct {
	HttpResponse *http.Response
	Model        *PrivateEndpointConnectionListResult
}

// PrivateEndpointConnectionsListByStorageSyncService ...
func (c PrivateEndpointConnectionResourceClient) PrivateEndpointConnectionsListByStorageSyncService(ctx context.Context, id StorageSyncServiceId) (result PrivateEndpointConnectionsListByStorageSyncServiceOperationResponse, err error) {
	req, err := c.preparerForPrivateEndpointConnectionsListByStorageSyncService(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "privateendpointconnectionresource.PrivateEndpointConnectionResourceClient", "PrivateEndpointConnectionsListByStorageSyncService", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "privateendpointconnectionresource.PrivateEndpointConnectionResourceClient", "PrivateEndpointConnectionsListByStorageSyncService", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForPrivateEndpointConnectionsListByStorageSyncService(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "privateendpointconnectionresource.PrivateEndpointConnectionResourceClient", "PrivateEndpointConnectionsListByStorageSyncService", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForPrivateEndpointConnectionsListByStorageSyncService prepares the PrivateEndpointConnectionsListByStorageSyncService request.
func (c PrivateEndpointConnectionResourceClient) preparerForPrivateEndpointConnectionsListByStorageSyncService(ctx context.Context, id StorageSyncServiceId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/privateEndpointConnections", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForPrivateEndpointConnectionsListByStorageSyncService handles the response to the PrivateEndpointConnectionsListByStorageSyncService request. The method always
// closes the http.Response Body.
func (c PrivateEndpointConnectionResourceClient) responderForPrivateEndpointConnectionsListByStorageSyncService(resp *http.Response) (result PrivateEndpointConnectionsListByStorageSyncServiceOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

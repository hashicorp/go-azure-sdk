package privateendpointconnections

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivateLinkHubGetOperationResponse struct {
	HttpResponse *http.Response
	Model        *PrivateEndpointConnectionForPrivateLinkHub
}

// PrivateLinkHubGet ...
func (c PrivateEndpointConnectionsClient) PrivateLinkHubGet(ctx context.Context, id PrivateEndpointConnectionId) (result PrivateLinkHubGetOperationResponse, err error) {
	req, err := c.preparerForPrivateLinkHubGet(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "privateendpointconnections.PrivateEndpointConnectionsClient", "PrivateLinkHubGet", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "privateendpointconnections.PrivateEndpointConnectionsClient", "PrivateLinkHubGet", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForPrivateLinkHubGet(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "privateendpointconnections.PrivateEndpointConnectionsClient", "PrivateLinkHubGet", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForPrivateLinkHubGet prepares the PrivateLinkHubGet request.
func (c PrivateEndpointConnectionsClient) preparerForPrivateLinkHubGet(ctx context.Context, id PrivateEndpointConnectionId) (*http.Request, error) {
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

// responderForPrivateLinkHubGet handles the response to the PrivateLinkHubGet request. The method always
// closes the http.Response Body.
func (c PrivateEndpointConnectionsClient) responderForPrivateLinkHubGet(resp *http.Response) (result PrivateLinkHubGetOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

package privatelink

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivateEndpointConnectionsDeleteByHostPoolOperationResponse struct {
	HttpResponse *http.Response
}

// PrivateEndpointConnectionsDeleteByHostPool ...
func (c PrivateLinkClient) PrivateEndpointConnectionsDeleteByHostPool(ctx context.Context, id PrivateEndpointConnectionId) (result PrivateEndpointConnectionsDeleteByHostPoolOperationResponse, err error) {
	req, err := c.preparerForPrivateEndpointConnectionsDeleteByHostPool(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "privatelink.PrivateLinkClient", "PrivateEndpointConnectionsDeleteByHostPool", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "privatelink.PrivateLinkClient", "PrivateEndpointConnectionsDeleteByHostPool", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForPrivateEndpointConnectionsDeleteByHostPool(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "privatelink.PrivateLinkClient", "PrivateEndpointConnectionsDeleteByHostPool", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForPrivateEndpointConnectionsDeleteByHostPool prepares the PrivateEndpointConnectionsDeleteByHostPool request.
func (c PrivateLinkClient) preparerForPrivateEndpointConnectionsDeleteByHostPool(ctx context.Context, id PrivateEndpointConnectionId) (*http.Request, error) {
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

// responderForPrivateEndpointConnectionsDeleteByHostPool handles the response to the PrivateEndpointConnectionsDeleteByHostPool request. The method always
// closes the http.Response Body.
func (c PrivateLinkClient) responderForPrivateEndpointConnectionsDeleteByHostPool(resp *http.Response) (result PrivateEndpointConnectionsDeleteByHostPoolOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusNoContent, http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp
	return
}

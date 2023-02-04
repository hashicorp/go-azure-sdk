package privateendpointconnection

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeletePrivateEndpointConnectionOperationResponse struct {
	HttpResponse *http.Response
}

// DeletePrivateEndpointConnection ...
func (c PrivateEndpointConnectionClient) DeletePrivateEndpointConnection(ctx context.Context, id PrivateEndpointConnectionId) (result DeletePrivateEndpointConnectionOperationResponse, err error) {
	req, err := c.preparerForDeletePrivateEndpointConnection(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "privateendpointconnection.PrivateEndpointConnectionClient", "DeletePrivateEndpointConnection", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "privateendpointconnection.PrivateEndpointConnectionClient", "DeletePrivateEndpointConnection", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForDeletePrivateEndpointConnection(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "privateendpointconnection.PrivateEndpointConnectionClient", "DeletePrivateEndpointConnection", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForDeletePrivateEndpointConnection prepares the DeletePrivateEndpointConnection request.
func (c PrivateEndpointConnectionClient) preparerForDeletePrivateEndpointConnection(ctx context.Context, id PrivateEndpointConnectionId) (*http.Request, error) {
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

// responderForDeletePrivateEndpointConnection handles the response to the DeletePrivateEndpointConnection request. The method always
// closes the http.Response Body.
func (c PrivateEndpointConnectionClient) responderForDeletePrivateEndpointConnection(resp *http.Response) (result DeletePrivateEndpointConnectionOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusNoContent, http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

package privateendpointconnection

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PutPrivateEndpointConnectionOperationResponse struct {
	HttpResponse *http.Response
	Model        *PrivateEndpointConnection
}

// PutPrivateEndpointConnection ...
func (c PrivateEndpointConnectionClient) PutPrivateEndpointConnection(ctx context.Context, id PrivateEndpointConnectionId, input PrivateEndpointConnection) (result PutPrivateEndpointConnectionOperationResponse, err error) {
	req, err := c.preparerForPutPrivateEndpointConnection(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "privateendpointconnection.PrivateEndpointConnectionClient", "PutPrivateEndpointConnection", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "privateendpointconnection.PrivateEndpointConnectionClient", "PutPrivateEndpointConnection", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForPutPrivateEndpointConnection(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "privateendpointconnection.PrivateEndpointConnectionClient", "PutPrivateEndpointConnection", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForPutPrivateEndpointConnection prepares the PutPrivateEndpointConnection request.
func (c PrivateEndpointConnectionClient) preparerForPutPrivateEndpointConnection(ctx context.Context, id PrivateEndpointConnectionId, input PrivateEndpointConnection) (*http.Request, error) {
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

// responderForPutPrivateEndpointConnection handles the response to the PutPrivateEndpointConnection request. The method always
// closes the http.Response Body.
func (c PrivateEndpointConnectionClient) responderForPutPrivateEndpointConnection(resp *http.Response) (result PutPrivateEndpointConnectionOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

package privateendpointconnectionresources

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivateEndpointConnectionGetOperationResponse struct {
	HttpResponse *http.Response
	Model        *PrivateEndpointConnectionResource
}

// PrivateEndpointConnectionGet ...
func (c PrivateEndpointConnectionResourcesClient) PrivateEndpointConnectionGet(ctx context.Context, id PrivateEndpointConnectionId) (result PrivateEndpointConnectionGetOperationResponse, err error) {
	req, err := c.preparerForPrivateEndpointConnectionGet(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "privateendpointconnectionresources.PrivateEndpointConnectionResourcesClient", "PrivateEndpointConnectionGet", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "privateendpointconnectionresources.PrivateEndpointConnectionResourcesClient", "PrivateEndpointConnectionGet", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForPrivateEndpointConnectionGet(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "privateendpointconnectionresources.PrivateEndpointConnectionResourcesClient", "PrivateEndpointConnectionGet", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForPrivateEndpointConnectionGet prepares the PrivateEndpointConnectionGet request.
func (c PrivateEndpointConnectionResourcesClient) preparerForPrivateEndpointConnectionGet(ctx context.Context, id PrivateEndpointConnectionId) (*http.Request, error) {
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

// responderForPrivateEndpointConnectionGet handles the response to the PrivateEndpointConnectionGet request. The method always
// closes the http.Response Body.
func (c PrivateEndpointConnectionResourcesClient) responderForPrivateEndpointConnectionGet(resp *http.Response) (result PrivateEndpointConnectionGetOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

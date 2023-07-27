package privateendpointconnections

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivateEndpointConnectionListByServiceOperationResponse struct {
	HttpResponse *http.Response
	Model        *PrivateEndpointConnectionListResult
}

// PrivateEndpointConnectionListByService ...
func (c PrivateEndpointConnectionsClient) PrivateEndpointConnectionListByService(ctx context.Context, id ServiceId) (result PrivateEndpointConnectionListByServiceOperationResponse, err error) {
	req, err := c.preparerForPrivateEndpointConnectionListByService(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "privateendpointconnections.PrivateEndpointConnectionsClient", "PrivateEndpointConnectionListByService", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "privateendpointconnections.PrivateEndpointConnectionsClient", "PrivateEndpointConnectionListByService", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForPrivateEndpointConnectionListByService(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "privateendpointconnections.PrivateEndpointConnectionsClient", "PrivateEndpointConnectionListByService", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForPrivateEndpointConnectionListByService prepares the PrivateEndpointConnectionListByService request.
func (c PrivateEndpointConnectionsClient) preparerForPrivateEndpointConnectionListByService(ctx context.Context, id ServiceId) (*http.Request, error) {
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

// responderForPrivateEndpointConnectionListByService handles the response to the PrivateEndpointConnectionListByService request. The method always
// closes the http.Response Body.
func (c PrivateEndpointConnectionsClient) responderForPrivateEndpointConnectionListByService(resp *http.Response) (result PrivateEndpointConnectionListByServiceOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

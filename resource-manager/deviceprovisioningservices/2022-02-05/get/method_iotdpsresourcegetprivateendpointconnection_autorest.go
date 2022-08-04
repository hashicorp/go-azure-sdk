package get

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IotDpsResourceGetPrivateEndpointConnectionOperationResponse struct {
	HttpResponse *http.Response
	Model        *PrivateEndpointConnection
}

// IotDpsResourceGetPrivateEndpointConnection ...
func (c GETClient) IotDpsResourceGetPrivateEndpointConnection(ctx context.Context, id PrivateEndpointConnectionId) (result IotDpsResourceGetPrivateEndpointConnectionOperationResponse, err error) {
	req, err := c.preparerForIotDpsResourceGetPrivateEndpointConnection(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "get.GETClient", "IotDpsResourceGetPrivateEndpointConnection", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "get.GETClient", "IotDpsResourceGetPrivateEndpointConnection", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForIotDpsResourceGetPrivateEndpointConnection(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "get.GETClient", "IotDpsResourceGetPrivateEndpointConnection", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForIotDpsResourceGetPrivateEndpointConnection prepares the IotDpsResourceGetPrivateEndpointConnection request.
func (c GETClient) preparerForIotDpsResourceGetPrivateEndpointConnection(ctx context.Context, id PrivateEndpointConnectionId) (*http.Request, error) {
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

// responderForIotDpsResourceGetPrivateEndpointConnection handles the response to the IotDpsResourceGetPrivateEndpointConnection request. The method always
// closes the http.Response Body.
func (c GETClient) responderForIotDpsResourceGetPrivateEndpointConnection(resp *http.Response) (result IotDpsResourceGetPrivateEndpointConnectionOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

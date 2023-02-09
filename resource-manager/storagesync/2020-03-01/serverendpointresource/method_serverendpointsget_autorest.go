package serverendpointresource

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServerEndpointsGetOperationResponse struct {
	HttpResponse *http.Response
	Model        *ServerEndpoint
}

// ServerEndpointsGet ...
func (c ServerEndpointResourceClient) ServerEndpointsGet(ctx context.Context, id ServerEndpointId) (result ServerEndpointsGetOperationResponse, err error) {
	req, err := c.preparerForServerEndpointsGet(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "serverendpointresource.ServerEndpointResourceClient", "ServerEndpointsGet", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "serverendpointresource.ServerEndpointResourceClient", "ServerEndpointsGet", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForServerEndpointsGet(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "serverendpointresource.ServerEndpointResourceClient", "ServerEndpointsGet", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForServerEndpointsGet prepares the ServerEndpointsGet request.
func (c ServerEndpointResourceClient) preparerForServerEndpointsGet(ctx context.Context, id ServerEndpointId) (*http.Request, error) {
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

// responderForServerEndpointsGet handles the response to the ServerEndpointsGet request. The method always
// closes the http.Response Body.
func (c ServerEndpointResourceClient) responderForServerEndpointsGet(resp *http.Response) (result ServerEndpointsGetOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

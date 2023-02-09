package serverendpointresource

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServerEndpointsListBySyncGroupOperationResponse struct {
	HttpResponse *http.Response
	Model        *ServerEndpointArray
}

// ServerEndpointsListBySyncGroup ...
func (c ServerEndpointResourceClient) ServerEndpointsListBySyncGroup(ctx context.Context, id SyncGroupId) (result ServerEndpointsListBySyncGroupOperationResponse, err error) {
	req, err := c.preparerForServerEndpointsListBySyncGroup(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "serverendpointresource.ServerEndpointResourceClient", "ServerEndpointsListBySyncGroup", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "serverendpointresource.ServerEndpointResourceClient", "ServerEndpointsListBySyncGroup", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForServerEndpointsListBySyncGroup(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "serverendpointresource.ServerEndpointResourceClient", "ServerEndpointsListBySyncGroup", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForServerEndpointsListBySyncGroup prepares the ServerEndpointsListBySyncGroup request.
func (c ServerEndpointResourceClient) preparerForServerEndpointsListBySyncGroup(ctx context.Context, id SyncGroupId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/serverEndpoints", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForServerEndpointsListBySyncGroup handles the response to the ServerEndpointsListBySyncGroup request. The method always
// closes the http.Response Body.
func (c ServerEndpointResourceClient) responderForServerEndpointsListBySyncGroup(resp *http.Response) (result ServerEndpointsListBySyncGroupOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

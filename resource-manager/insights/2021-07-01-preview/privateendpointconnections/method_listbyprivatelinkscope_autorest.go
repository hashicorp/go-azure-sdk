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

type ListByPrivateLinkScopeOperationResponse struct {
	HttpResponse *http.Response
	Model        *PrivateEndpointConnectionListResult
}

// ListByPrivateLinkScope ...
func (c PrivateEndpointConnectionsClient) ListByPrivateLinkScope(ctx context.Context, id PrivateLinkScopeId) (result ListByPrivateLinkScopeOperationResponse, err error) {
	req, err := c.preparerForListByPrivateLinkScope(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "privateendpointconnections.PrivateEndpointConnectionsClient", "ListByPrivateLinkScope", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "privateendpointconnections.PrivateEndpointConnectionsClient", "ListByPrivateLinkScope", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForListByPrivateLinkScope(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "privateendpointconnections.PrivateEndpointConnectionsClient", "ListByPrivateLinkScope", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForListByPrivateLinkScope prepares the ListByPrivateLinkScope request.
func (c PrivateEndpointConnectionsClient) preparerForListByPrivateLinkScope(ctx context.Context, id PrivateLinkScopeId) (*http.Request, error) {
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

// responderForListByPrivateLinkScope handles the response to the ListByPrivateLinkScope request. The method always
// closes the http.Response Body.
func (c PrivateEndpointConnectionsClient) responderForListByPrivateLinkScope(resp *http.Response) (result ListByPrivateLinkScopeOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

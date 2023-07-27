package privateendpointconnections

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivateEndpointConnectionGetPrivateLinkResourceOperationResponse struct {
	HttpResponse *http.Response
	Model        *PrivateLinkResource
}

// PrivateEndpointConnectionGetPrivateLinkResource ...
func (c PrivateEndpointConnectionsClient) PrivateEndpointConnectionGetPrivateLinkResource(ctx context.Context, id PrivateLinkResourceId) (result PrivateEndpointConnectionGetPrivateLinkResourceOperationResponse, err error) {
	req, err := c.preparerForPrivateEndpointConnectionGetPrivateLinkResource(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "privateendpointconnections.PrivateEndpointConnectionsClient", "PrivateEndpointConnectionGetPrivateLinkResource", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "privateendpointconnections.PrivateEndpointConnectionsClient", "PrivateEndpointConnectionGetPrivateLinkResource", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForPrivateEndpointConnectionGetPrivateLinkResource(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "privateendpointconnections.PrivateEndpointConnectionsClient", "PrivateEndpointConnectionGetPrivateLinkResource", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForPrivateEndpointConnectionGetPrivateLinkResource prepares the PrivateEndpointConnectionGetPrivateLinkResource request.
func (c PrivateEndpointConnectionsClient) preparerForPrivateEndpointConnectionGetPrivateLinkResource(ctx context.Context, id PrivateLinkResourceId) (*http.Request, error) {
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

// responderForPrivateEndpointConnectionGetPrivateLinkResource handles the response to the PrivateEndpointConnectionGetPrivateLinkResource request. The method always
// closes the http.Response Body.
func (c PrivateEndpointConnectionsClient) responderForPrivateEndpointConnectionGetPrivateLinkResource(resp *http.Response) (result PrivateEndpointConnectionGetPrivateLinkResourceOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

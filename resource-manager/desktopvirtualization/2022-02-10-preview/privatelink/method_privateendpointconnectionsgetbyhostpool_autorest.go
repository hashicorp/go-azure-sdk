package privatelink

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivateEndpointConnectionsGetByHostPoolOperationResponse struct {
	HttpResponse *http.Response
	Model        *PrivateEndpointConnectionWithSystemData
}

// PrivateEndpointConnectionsGetByHostPool ...
func (c PrivateLinkClient) PrivateEndpointConnectionsGetByHostPool(ctx context.Context, id PrivateEndpointConnectionId) (result PrivateEndpointConnectionsGetByHostPoolOperationResponse, err error) {
	req, err := c.preparerForPrivateEndpointConnectionsGetByHostPool(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "privatelink.PrivateLinkClient", "PrivateEndpointConnectionsGetByHostPool", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "privatelink.PrivateLinkClient", "PrivateEndpointConnectionsGetByHostPool", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForPrivateEndpointConnectionsGetByHostPool(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "privatelink.PrivateLinkClient", "PrivateEndpointConnectionsGetByHostPool", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForPrivateEndpointConnectionsGetByHostPool prepares the PrivateEndpointConnectionsGetByHostPool request.
func (c PrivateLinkClient) preparerForPrivateEndpointConnectionsGetByHostPool(ctx context.Context, id PrivateEndpointConnectionId) (*http.Request, error) {
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

// responderForPrivateEndpointConnectionsGetByHostPool handles the response to the PrivateEndpointConnectionsGetByHostPool request. The method always
// closes the http.Response Body.
func (c PrivateLinkClient) responderForPrivateEndpointConnectionsGetByHostPool(resp *http.Response) (result PrivateEndpointConnectionsGetByHostPoolOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

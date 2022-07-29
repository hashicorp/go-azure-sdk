package privatelink

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivateEndpointConnectionsUpdateByHostPoolOperationResponse struct {
	HttpResponse *http.Response
	Model        *PrivateEndpointConnectionWithSystemData
}

// PrivateEndpointConnectionsUpdateByHostPool ...
func (c PrivateLinkClient) PrivateEndpointConnectionsUpdateByHostPool(ctx context.Context, id PrivateEndpointConnectionId, input PrivateEndpointConnection) (result PrivateEndpointConnectionsUpdateByHostPoolOperationResponse, err error) {
	req, err := c.preparerForPrivateEndpointConnectionsUpdateByHostPool(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "privatelink.PrivateLinkClient", "PrivateEndpointConnectionsUpdateByHostPool", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "privatelink.PrivateLinkClient", "PrivateEndpointConnectionsUpdateByHostPool", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForPrivateEndpointConnectionsUpdateByHostPool(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "privatelink.PrivateLinkClient", "PrivateEndpointConnectionsUpdateByHostPool", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForPrivateEndpointConnectionsUpdateByHostPool prepares the PrivateEndpointConnectionsUpdateByHostPool request.
func (c PrivateLinkClient) preparerForPrivateEndpointConnectionsUpdateByHostPool(ctx context.Context, id PrivateEndpointConnectionId, input PrivateEndpointConnection) (*http.Request, error) {
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

// responderForPrivateEndpointConnectionsUpdateByHostPool handles the response to the PrivateEndpointConnectionsUpdateByHostPool request. The method always
// closes the http.Response Body.
func (c PrivateLinkClient) responderForPrivateEndpointConnectionsUpdateByHostPool(resp *http.Response) (result PrivateEndpointConnectionsUpdateByHostPoolOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

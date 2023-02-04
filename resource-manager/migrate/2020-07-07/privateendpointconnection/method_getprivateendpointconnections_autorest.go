package privateendpointconnection

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetPrivateEndpointConnectionsOperationResponse struct {
	HttpResponse *http.Response
	Model        *PrivateEndpointConnectionCollection
}

// GetPrivateEndpointConnections ...
func (c PrivateEndpointConnectionClient) GetPrivateEndpointConnections(ctx context.Context, id MasterSiteId) (result GetPrivateEndpointConnectionsOperationResponse, err error) {
	req, err := c.preparerForGetPrivateEndpointConnections(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "privateendpointconnection.PrivateEndpointConnectionClient", "GetPrivateEndpointConnections", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "privateendpointconnection.PrivateEndpointConnectionClient", "GetPrivateEndpointConnections", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetPrivateEndpointConnections(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "privateendpointconnection.PrivateEndpointConnectionClient", "GetPrivateEndpointConnections", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetPrivateEndpointConnections prepares the GetPrivateEndpointConnections request.
func (c PrivateEndpointConnectionClient) preparerForGetPrivateEndpointConnections(ctx context.Context, id MasterSiteId) (*http.Request, error) {
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

// responderForGetPrivateEndpointConnections handles the response to the GetPrivateEndpointConnections request. The method always
// closes the http.Response Body.
func (c PrivateEndpointConnectionClient) responderForGetPrivateEndpointConnections(resp *http.Response) (result GetPrivateEndpointConnectionsOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

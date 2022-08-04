package get

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IotDpsResourceListPrivateEndpointConnectionsOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]PrivateEndpointConnection
}

// IotDpsResourceListPrivateEndpointConnections ...
func (c GETClient) IotDpsResourceListPrivateEndpointConnections(ctx context.Context, id ProvisioningServiceId) (result IotDpsResourceListPrivateEndpointConnectionsOperationResponse, err error) {
	req, err := c.preparerForIotDpsResourceListPrivateEndpointConnections(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "get.GETClient", "IotDpsResourceListPrivateEndpointConnections", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "get.GETClient", "IotDpsResourceListPrivateEndpointConnections", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForIotDpsResourceListPrivateEndpointConnections(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "get.GETClient", "IotDpsResourceListPrivateEndpointConnections", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForIotDpsResourceListPrivateEndpointConnections prepares the IotDpsResourceListPrivateEndpointConnections request.
func (c GETClient) preparerForIotDpsResourceListPrivateEndpointConnections(ctx context.Context, id ProvisioningServiceId) (*http.Request, error) {
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

// responderForIotDpsResourceListPrivateEndpointConnections handles the response to the IotDpsResourceListPrivateEndpointConnections request. The method always
// closes the http.Response Body.
func (c GETClient) responderForIotDpsResourceListPrivateEndpointConnections(resp *http.Response) (result IotDpsResourceListPrivateEndpointConnectionsOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

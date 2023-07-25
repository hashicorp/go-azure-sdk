package webapps

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetRelayServiceConnectionOperationResponse struct {
	HttpResponse *http.Response
	Model        *RelayServiceConnectionEntity
}

// GetRelayServiceConnection ...
func (c WebAppsClient) GetRelayServiceConnection(ctx context.Context, id HybridConnectionId) (result GetRelayServiceConnectionOperationResponse, err error) {
	req, err := c.preparerForGetRelayServiceConnection(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetRelayServiceConnection", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetRelayServiceConnection", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetRelayServiceConnection(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetRelayServiceConnection", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetRelayServiceConnection prepares the GetRelayServiceConnection request.
func (c WebAppsClient) preparerForGetRelayServiceConnection(ctx context.Context, id HybridConnectionId) (*http.Request, error) {
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

// responderForGetRelayServiceConnection handles the response to the GetRelayServiceConnection request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForGetRelayServiceConnection(resp *http.Response) (result GetRelayServiceConnectionOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

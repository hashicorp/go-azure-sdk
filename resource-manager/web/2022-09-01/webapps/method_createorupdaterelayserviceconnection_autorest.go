package webapps

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateOrUpdateRelayServiceConnectionOperationResponse struct {
	HttpResponse *http.Response
	Model        *RelayServiceConnectionEntity
}

// CreateOrUpdateRelayServiceConnection ...
func (c WebAppsClient) CreateOrUpdateRelayServiceConnection(ctx context.Context, id HybridConnectionId, input RelayServiceConnectionEntity) (result CreateOrUpdateRelayServiceConnectionOperationResponse, err error) {
	req, err := c.preparerForCreateOrUpdateRelayServiceConnection(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "CreateOrUpdateRelayServiceConnection", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "CreateOrUpdateRelayServiceConnection", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForCreateOrUpdateRelayServiceConnection(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "CreateOrUpdateRelayServiceConnection", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForCreateOrUpdateRelayServiceConnection prepares the CreateOrUpdateRelayServiceConnection request.
func (c WebAppsClient) preparerForCreateOrUpdateRelayServiceConnection(ctx context.Context, id HybridConnectionId, input RelayServiceConnectionEntity) (*http.Request, error) {
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

// responderForCreateOrUpdateRelayServiceConnection handles the response to the CreateOrUpdateRelayServiceConnection request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForCreateOrUpdateRelayServiceConnection(resp *http.Response) (result CreateOrUpdateRelayServiceConnectionOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

package webapps

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateRelayServiceConnectionOperationResponse struct {
	HttpResponse *http.Response
	Model        *RelayServiceConnectionEntity
}

// UpdateRelayServiceConnection ...
func (c WebAppsClient) UpdateRelayServiceConnection(ctx context.Context, id HybridConnectionId, input RelayServiceConnectionEntity) (result UpdateRelayServiceConnectionOperationResponse, err error) {
	req, err := c.preparerForUpdateRelayServiceConnection(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "UpdateRelayServiceConnection", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "UpdateRelayServiceConnection", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForUpdateRelayServiceConnection(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "UpdateRelayServiceConnection", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForUpdateRelayServiceConnection prepares the UpdateRelayServiceConnection request.
func (c WebAppsClient) preparerForUpdateRelayServiceConnection(ctx context.Context, id HybridConnectionId, input RelayServiceConnectionEntity) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForUpdateRelayServiceConnection handles the response to the UpdateRelayServiceConnection request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForUpdateRelayServiceConnection(resp *http.Response) (result UpdateRelayServiceConnectionOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

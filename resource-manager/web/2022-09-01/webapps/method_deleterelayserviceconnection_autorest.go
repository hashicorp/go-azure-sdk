package webapps

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeleteRelayServiceConnectionOperationResponse struct {
	HttpResponse *http.Response
}

// DeleteRelayServiceConnection ...
func (c WebAppsClient) DeleteRelayServiceConnection(ctx context.Context, id HybridConnectionId) (result DeleteRelayServiceConnectionOperationResponse, err error) {
	req, err := c.preparerForDeleteRelayServiceConnection(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "DeleteRelayServiceConnection", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "DeleteRelayServiceConnection", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForDeleteRelayServiceConnection(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "DeleteRelayServiceConnection", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForDeleteRelayServiceConnection prepares the DeleteRelayServiceConnection request.
func (c WebAppsClient) preparerForDeleteRelayServiceConnection(ctx context.Context, id HybridConnectionId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsDelete(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForDeleteRelayServiceConnection handles the response to the DeleteRelayServiceConnection request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForDeleteRelayServiceConnection(resp *http.Response) (result DeleteRelayServiceConnectionOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

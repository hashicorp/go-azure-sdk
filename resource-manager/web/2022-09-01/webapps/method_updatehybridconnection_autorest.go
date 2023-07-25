package webapps

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateHybridConnectionOperationResponse struct {
	HttpResponse *http.Response
	Model        *HybridConnection
}

// UpdateHybridConnection ...
func (c WebAppsClient) UpdateHybridConnection(ctx context.Context, id RelayId, input HybridConnection) (result UpdateHybridConnectionOperationResponse, err error) {
	req, err := c.preparerForUpdateHybridConnection(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "UpdateHybridConnection", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "UpdateHybridConnection", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForUpdateHybridConnection(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "UpdateHybridConnection", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForUpdateHybridConnection prepares the UpdateHybridConnection request.
func (c WebAppsClient) preparerForUpdateHybridConnection(ctx context.Context, id RelayId, input HybridConnection) (*http.Request, error) {
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

// responderForUpdateHybridConnection handles the response to the UpdateHybridConnection request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForUpdateHybridConnection(resp *http.Response) (result UpdateHybridConnectionOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

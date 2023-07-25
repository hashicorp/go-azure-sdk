package webapps

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateOrUpdateHybridConnectionOperationResponse struct {
	HttpResponse *http.Response
	Model        *HybridConnection
}

// CreateOrUpdateHybridConnection ...
func (c WebAppsClient) CreateOrUpdateHybridConnection(ctx context.Context, id RelayId, input HybridConnection) (result CreateOrUpdateHybridConnectionOperationResponse, err error) {
	req, err := c.preparerForCreateOrUpdateHybridConnection(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "CreateOrUpdateHybridConnection", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "CreateOrUpdateHybridConnection", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForCreateOrUpdateHybridConnection(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "CreateOrUpdateHybridConnection", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForCreateOrUpdateHybridConnection prepares the CreateOrUpdateHybridConnection request.
func (c WebAppsClient) preparerForCreateOrUpdateHybridConnection(ctx context.Context, id RelayId, input HybridConnection) (*http.Request, error) {
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

// responderForCreateOrUpdateHybridConnection handles the response to the CreateOrUpdateHybridConnection request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForCreateOrUpdateHybridConnection(resp *http.Response) (result CreateOrUpdateHybridConnectionOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

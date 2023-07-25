package webapps

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListHybridConnectionsOperationResponse struct {
	HttpResponse *http.Response
	Model        *HybridConnection
}

// ListHybridConnections ...
func (c WebAppsClient) ListHybridConnections(ctx context.Context, id SiteId) (result ListHybridConnectionsOperationResponse, err error) {
	req, err := c.preparerForListHybridConnections(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListHybridConnections", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListHybridConnections", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForListHybridConnections(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListHybridConnections", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForListHybridConnections prepares the ListHybridConnections request.
func (c WebAppsClient) preparerForListHybridConnections(ctx context.Context, id SiteId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/hybridConnectionRelays", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForListHybridConnections handles the response to the ListHybridConnections request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForListHybridConnections(resp *http.Response) (result ListHybridConnectionsOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

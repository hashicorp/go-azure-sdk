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

type ListRelayServiceConnectionsOperationResponse struct {
	HttpResponse *http.Response
	Model        *RelayServiceConnectionEntity
}

// ListRelayServiceConnections ...
func (c WebAppsClient) ListRelayServiceConnections(ctx context.Context, id SiteId) (result ListRelayServiceConnectionsOperationResponse, err error) {
	req, err := c.preparerForListRelayServiceConnections(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListRelayServiceConnections", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListRelayServiceConnections", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForListRelayServiceConnections(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListRelayServiceConnections", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForListRelayServiceConnections prepares the ListRelayServiceConnections request.
func (c WebAppsClient) preparerForListRelayServiceConnections(ctx context.Context, id SiteId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/hybridConnection", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForListRelayServiceConnections handles the response to the ListRelayServiceConnections request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForListRelayServiceConnections(resp *http.Response) (result ListRelayServiceConnectionsOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

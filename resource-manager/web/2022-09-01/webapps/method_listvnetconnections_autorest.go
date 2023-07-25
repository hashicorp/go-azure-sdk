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

type ListVnetConnectionsOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]VnetInfoResource
}

// ListVnetConnections ...
func (c WebAppsClient) ListVnetConnections(ctx context.Context, id SiteId) (result ListVnetConnectionsOperationResponse, err error) {
	req, err := c.preparerForListVnetConnections(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListVnetConnections", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListVnetConnections", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForListVnetConnections(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListVnetConnections", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForListVnetConnections prepares the ListVnetConnections request.
func (c WebAppsClient) preparerForListVnetConnections(ctx context.Context, id SiteId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/virtualNetworkConnections", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForListVnetConnections handles the response to the ListVnetConnections request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForListVnetConnections(resp *http.Response) (result ListVnetConnectionsOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

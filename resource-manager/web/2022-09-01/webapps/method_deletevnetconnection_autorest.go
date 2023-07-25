package webapps

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeleteVnetConnectionOperationResponse struct {
	HttpResponse *http.Response
}

// DeleteVnetConnection ...
func (c WebAppsClient) DeleteVnetConnection(ctx context.Context, id VirtualNetworkConnectionId) (result DeleteVnetConnectionOperationResponse, err error) {
	req, err := c.preparerForDeleteVnetConnection(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "DeleteVnetConnection", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "DeleteVnetConnection", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForDeleteVnetConnection(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "DeleteVnetConnection", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForDeleteVnetConnection prepares the DeleteVnetConnection request.
func (c WebAppsClient) preparerForDeleteVnetConnection(ctx context.Context, id VirtualNetworkConnectionId) (*http.Request, error) {
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

// responderForDeleteVnetConnection handles the response to the DeleteVnetConnection request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForDeleteVnetConnection(resp *http.Response) (result DeleteVnetConnectionOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

package webapps

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateVnetConnectionOperationResponse struct {
	HttpResponse *http.Response
	Model        *VnetInfoResource
}

// UpdateVnetConnection ...
func (c WebAppsClient) UpdateVnetConnection(ctx context.Context, id VirtualNetworkConnectionId, input VnetInfoResource) (result UpdateVnetConnectionOperationResponse, err error) {
	req, err := c.preparerForUpdateVnetConnection(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "UpdateVnetConnection", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "UpdateVnetConnection", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForUpdateVnetConnection(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "UpdateVnetConnection", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForUpdateVnetConnection prepares the UpdateVnetConnection request.
func (c WebAppsClient) preparerForUpdateVnetConnection(ctx context.Context, id VirtualNetworkConnectionId, input VnetInfoResource) (*http.Request, error) {
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

// responderForUpdateVnetConnection handles the response to the UpdateVnetConnection request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForUpdateVnetConnection(resp *http.Response) (result UpdateVnetConnectionOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

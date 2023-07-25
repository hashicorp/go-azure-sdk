package webapps

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateOrUpdateVnetConnectionOperationResponse struct {
	HttpResponse *http.Response
	Model        *VnetInfoResource
}

// CreateOrUpdateVnetConnection ...
func (c WebAppsClient) CreateOrUpdateVnetConnection(ctx context.Context, id VirtualNetworkConnectionId, input VnetInfoResource) (result CreateOrUpdateVnetConnectionOperationResponse, err error) {
	req, err := c.preparerForCreateOrUpdateVnetConnection(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "CreateOrUpdateVnetConnection", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "CreateOrUpdateVnetConnection", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForCreateOrUpdateVnetConnection(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "CreateOrUpdateVnetConnection", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForCreateOrUpdateVnetConnection prepares the CreateOrUpdateVnetConnection request.
func (c WebAppsClient) preparerForCreateOrUpdateVnetConnection(ctx context.Context, id VirtualNetworkConnectionId, input VnetInfoResource) (*http.Request, error) {
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

// responderForCreateOrUpdateVnetConnection handles the response to the CreateOrUpdateVnetConnection request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForCreateOrUpdateVnetConnection(resp *http.Response) (result CreateOrUpdateVnetConnectionOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

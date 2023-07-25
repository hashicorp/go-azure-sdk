package webapps

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListNetworkFeaturesOperationResponse struct {
	HttpResponse *http.Response
	Model        *NetworkFeatures
}

// ListNetworkFeatures ...
func (c WebAppsClient) ListNetworkFeatures(ctx context.Context, id NetworkFeatureId) (result ListNetworkFeaturesOperationResponse, err error) {
	req, err := c.preparerForListNetworkFeatures(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListNetworkFeatures", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListNetworkFeatures", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForListNetworkFeatures(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListNetworkFeatures", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForListNetworkFeatures prepares the ListNetworkFeatures request.
func (c WebAppsClient) preparerForListNetworkFeatures(ctx context.Context, id NetworkFeatureId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForListNetworkFeatures handles the response to the ListNetworkFeatures request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForListNetworkFeatures(resp *http.Response) (result ListNetworkFeaturesOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

package webapps

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListNetworkFeaturesSlotOperationResponse struct {
	HttpResponse *http.Response
	Model        *NetworkFeatures
}

// ListNetworkFeaturesSlot ...
func (c WebAppsClient) ListNetworkFeaturesSlot(ctx context.Context, id SlotNetworkFeatureId) (result ListNetworkFeaturesSlotOperationResponse, err error) {
	req, err := c.preparerForListNetworkFeaturesSlot(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListNetworkFeaturesSlot", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListNetworkFeaturesSlot", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForListNetworkFeaturesSlot(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListNetworkFeaturesSlot", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForListNetworkFeaturesSlot prepares the ListNetworkFeaturesSlot request.
func (c WebAppsClient) preparerForListNetworkFeaturesSlot(ctx context.Context, id SlotNetworkFeatureId) (*http.Request, error) {
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

// responderForListNetworkFeaturesSlot handles the response to the ListNetworkFeaturesSlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForListNetworkFeaturesSlot(resp *http.Response) (result ListNetworkFeaturesSlotOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

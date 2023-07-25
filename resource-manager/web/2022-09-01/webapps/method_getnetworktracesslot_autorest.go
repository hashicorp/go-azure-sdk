package webapps

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetNetworkTracesSlotOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]NetworkTrace
}

// GetNetworkTracesSlot ...
func (c WebAppsClient) GetNetworkTracesSlot(ctx context.Context, id SlotNetworkTraceId) (result GetNetworkTracesSlotOperationResponse, err error) {
	req, err := c.preparerForGetNetworkTracesSlot(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetNetworkTracesSlot", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetNetworkTracesSlot", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetNetworkTracesSlot(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetNetworkTracesSlot", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetNetworkTracesSlot prepares the GetNetworkTracesSlot request.
func (c WebAppsClient) preparerForGetNetworkTracesSlot(ctx context.Context, id SlotNetworkTraceId) (*http.Request, error) {
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

// responderForGetNetworkTracesSlot handles the response to the GetNetworkTracesSlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForGetNetworkTracesSlot(resp *http.Response) (result GetNetworkTracesSlotOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

package webapps

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetNetworkTracesSlotV2OperationResponse struct {
	HttpResponse *http.Response
	Model        *[]NetworkTrace
}

// GetNetworkTracesSlotV2 ...
func (c WebAppsClient) GetNetworkTracesSlotV2(ctx context.Context, id SiteSlotNetworkTraceId) (result GetNetworkTracesSlotV2OperationResponse, err error) {
	req, err := c.preparerForGetNetworkTracesSlotV2(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetNetworkTracesSlotV2", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetNetworkTracesSlotV2", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetNetworkTracesSlotV2(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetNetworkTracesSlotV2", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetNetworkTracesSlotV2 prepares the GetNetworkTracesSlotV2 request.
func (c WebAppsClient) preparerForGetNetworkTracesSlotV2(ctx context.Context, id SiteSlotNetworkTraceId) (*http.Request, error) {
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

// responderForGetNetworkTracesSlotV2 handles the response to the GetNetworkTracesSlotV2 request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForGetNetworkTracesSlotV2(resp *http.Response) (result GetNetworkTracesSlotV2OperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

package webapps

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetNetworkTracesV2OperationResponse struct {
	HttpResponse *http.Response
	Model        *[]NetworkTrace
}

// GetNetworkTracesV2 ...
func (c WebAppsClient) GetNetworkTracesV2(ctx context.Context, id SiteNetworkTraceId) (result GetNetworkTracesV2OperationResponse, err error) {
	req, err := c.preparerForGetNetworkTracesV2(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetNetworkTracesV2", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetNetworkTracesV2", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetNetworkTracesV2(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetNetworkTracesV2", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetNetworkTracesV2 prepares the GetNetworkTracesV2 request.
func (c WebAppsClient) preparerForGetNetworkTracesV2(ctx context.Context, id SiteNetworkTraceId) (*http.Request, error) {
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

// responderForGetNetworkTracesV2 handles the response to the GetNetworkTracesV2 request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForGetNetworkTracesV2(resp *http.Response) (result GetNetworkTracesV2OperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

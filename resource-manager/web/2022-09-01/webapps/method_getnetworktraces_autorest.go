package webapps

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetNetworkTracesOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]NetworkTrace
}

// GetNetworkTraces ...
func (c WebAppsClient) GetNetworkTraces(ctx context.Context, id NetworkTraceId) (result GetNetworkTracesOperationResponse, err error) {
	req, err := c.preparerForGetNetworkTraces(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetNetworkTraces", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetNetworkTraces", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetNetworkTraces(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetNetworkTraces", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetNetworkTraces prepares the GetNetworkTraces request.
func (c WebAppsClient) preparerForGetNetworkTraces(ctx context.Context, id NetworkTraceId) (*http.Request, error) {
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

// responderForGetNetworkTraces handles the response to the GetNetworkTraces request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForGetNetworkTraces(resp *http.Response) (result GetNetworkTracesOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

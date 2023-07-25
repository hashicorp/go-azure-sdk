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

type StopNetworkTraceOperationResponse struct {
	HttpResponse *http.Response
}

// StopNetworkTrace ...
func (c WebAppsClient) StopNetworkTrace(ctx context.Context, id SiteId) (result StopNetworkTraceOperationResponse, err error) {
	req, err := c.preparerForStopNetworkTrace(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "StopNetworkTrace", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "StopNetworkTrace", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForStopNetworkTrace(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "StopNetworkTrace", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForStopNetworkTrace prepares the StopNetworkTrace request.
func (c WebAppsClient) preparerForStopNetworkTrace(ctx context.Context, id SiteId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/stopNetworkTrace", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForStopNetworkTrace handles the response to the StopNetworkTrace request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForStopNetworkTrace(resp *http.Response) (result StopNetworkTraceOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusNoContent, http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

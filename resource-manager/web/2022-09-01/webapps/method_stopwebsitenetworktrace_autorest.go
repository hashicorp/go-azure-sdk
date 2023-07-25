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

type StopWebSiteNetworkTraceOperationResponse struct {
	HttpResponse *http.Response
}

// StopWebSiteNetworkTrace ...
func (c WebAppsClient) StopWebSiteNetworkTrace(ctx context.Context, id SiteId) (result StopWebSiteNetworkTraceOperationResponse, err error) {
	req, err := c.preparerForStopWebSiteNetworkTrace(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "StopWebSiteNetworkTrace", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "StopWebSiteNetworkTrace", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForStopWebSiteNetworkTrace(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "StopWebSiteNetworkTrace", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForStopWebSiteNetworkTrace prepares the StopWebSiteNetworkTrace request.
func (c WebAppsClient) preparerForStopWebSiteNetworkTrace(ctx context.Context, id SiteId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/networkTrace/stop", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForStopWebSiteNetworkTrace handles the response to the StopWebSiteNetworkTrace request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForStopWebSiteNetworkTrace(resp *http.Response) (result StopWebSiteNetworkTraceOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusNoContent, http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

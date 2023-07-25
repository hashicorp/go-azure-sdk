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

type StopWebSiteNetworkTraceSlotOperationResponse struct {
	HttpResponse *http.Response
}

// StopWebSiteNetworkTraceSlot ...
func (c WebAppsClient) StopWebSiteNetworkTraceSlot(ctx context.Context, id SlotId) (result StopWebSiteNetworkTraceSlotOperationResponse, err error) {
	req, err := c.preparerForStopWebSiteNetworkTraceSlot(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "StopWebSiteNetworkTraceSlot", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "StopWebSiteNetworkTraceSlot", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForStopWebSiteNetworkTraceSlot(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "StopWebSiteNetworkTraceSlot", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForStopWebSiteNetworkTraceSlot prepares the StopWebSiteNetworkTraceSlot request.
func (c WebAppsClient) preparerForStopWebSiteNetworkTraceSlot(ctx context.Context, id SlotId) (*http.Request, error) {
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

// responderForStopWebSiteNetworkTraceSlot handles the response to the StopWebSiteNetworkTraceSlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForStopWebSiteNetworkTraceSlot(resp *http.Response) (result StopWebSiteNetworkTraceSlotOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusNoContent, http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

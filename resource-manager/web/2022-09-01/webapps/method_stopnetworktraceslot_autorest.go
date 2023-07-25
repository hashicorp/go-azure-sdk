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

type StopNetworkTraceSlotOperationResponse struct {
	HttpResponse *http.Response
}

// StopNetworkTraceSlot ...
func (c WebAppsClient) StopNetworkTraceSlot(ctx context.Context, id SlotId) (result StopNetworkTraceSlotOperationResponse, err error) {
	req, err := c.preparerForStopNetworkTraceSlot(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "StopNetworkTraceSlot", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "StopNetworkTraceSlot", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForStopNetworkTraceSlot(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "StopNetworkTraceSlot", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForStopNetworkTraceSlot prepares the StopNetworkTraceSlot request.
func (c WebAppsClient) preparerForStopNetworkTraceSlot(ctx context.Context, id SlotId) (*http.Request, error) {
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

// responderForStopNetworkTraceSlot handles the response to the StopNetworkTraceSlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForStopNetworkTraceSlot(resp *http.Response) (result StopNetworkTraceSlotOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusNoContent, http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

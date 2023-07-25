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

type ListHybridConnectionsSlotOperationResponse struct {
	HttpResponse *http.Response
	Model        *HybridConnection
}

// ListHybridConnectionsSlot ...
func (c WebAppsClient) ListHybridConnectionsSlot(ctx context.Context, id SlotId) (result ListHybridConnectionsSlotOperationResponse, err error) {
	req, err := c.preparerForListHybridConnectionsSlot(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListHybridConnectionsSlot", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListHybridConnectionsSlot", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForListHybridConnectionsSlot(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListHybridConnectionsSlot", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForListHybridConnectionsSlot prepares the ListHybridConnectionsSlot request.
func (c WebAppsClient) preparerForListHybridConnectionsSlot(ctx context.Context, id SlotId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/hybridConnectionRelays", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForListHybridConnectionsSlot handles the response to the ListHybridConnectionsSlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForListHybridConnectionsSlot(resp *http.Response) (result ListHybridConnectionsSlotOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

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

type ListRelayServiceConnectionsSlotOperationResponse struct {
	HttpResponse *http.Response
	Model        *RelayServiceConnectionEntity
}

// ListRelayServiceConnectionsSlot ...
func (c WebAppsClient) ListRelayServiceConnectionsSlot(ctx context.Context, id SlotId) (result ListRelayServiceConnectionsSlotOperationResponse, err error) {
	req, err := c.preparerForListRelayServiceConnectionsSlot(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListRelayServiceConnectionsSlot", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListRelayServiceConnectionsSlot", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForListRelayServiceConnectionsSlot(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListRelayServiceConnectionsSlot", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForListRelayServiceConnectionsSlot prepares the ListRelayServiceConnectionsSlot request.
func (c WebAppsClient) preparerForListRelayServiceConnectionsSlot(ctx context.Context, id SlotId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/hybridConnection", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForListRelayServiceConnectionsSlot handles the response to the ListRelayServiceConnectionsSlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForListRelayServiceConnectionsSlot(resp *http.Response) (result ListRelayServiceConnectionsSlotOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

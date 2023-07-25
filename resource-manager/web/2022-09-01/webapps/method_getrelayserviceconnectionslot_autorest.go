package webapps

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetRelayServiceConnectionSlotOperationResponse struct {
	HttpResponse *http.Response
	Model        *RelayServiceConnectionEntity
}

// GetRelayServiceConnectionSlot ...
func (c WebAppsClient) GetRelayServiceConnectionSlot(ctx context.Context, id SlotHybridConnectionId) (result GetRelayServiceConnectionSlotOperationResponse, err error) {
	req, err := c.preparerForGetRelayServiceConnectionSlot(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetRelayServiceConnectionSlot", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetRelayServiceConnectionSlot", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetRelayServiceConnectionSlot(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetRelayServiceConnectionSlot", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetRelayServiceConnectionSlot prepares the GetRelayServiceConnectionSlot request.
func (c WebAppsClient) preparerForGetRelayServiceConnectionSlot(ctx context.Context, id SlotHybridConnectionId) (*http.Request, error) {
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

// responderForGetRelayServiceConnectionSlot handles the response to the GetRelayServiceConnectionSlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForGetRelayServiceConnectionSlot(resp *http.Response) (result GetRelayServiceConnectionSlotOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

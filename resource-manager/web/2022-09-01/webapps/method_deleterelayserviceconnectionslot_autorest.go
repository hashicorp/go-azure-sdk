package webapps

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeleteRelayServiceConnectionSlotOperationResponse struct {
	HttpResponse *http.Response
}

// DeleteRelayServiceConnectionSlot ...
func (c WebAppsClient) DeleteRelayServiceConnectionSlot(ctx context.Context, id SlotHybridConnectionId) (result DeleteRelayServiceConnectionSlotOperationResponse, err error) {
	req, err := c.preparerForDeleteRelayServiceConnectionSlot(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "DeleteRelayServiceConnectionSlot", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "DeleteRelayServiceConnectionSlot", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForDeleteRelayServiceConnectionSlot(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "DeleteRelayServiceConnectionSlot", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForDeleteRelayServiceConnectionSlot prepares the DeleteRelayServiceConnectionSlot request.
func (c WebAppsClient) preparerForDeleteRelayServiceConnectionSlot(ctx context.Context, id SlotHybridConnectionId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsDelete(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForDeleteRelayServiceConnectionSlot handles the response to the DeleteRelayServiceConnectionSlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForDeleteRelayServiceConnectionSlot(resp *http.Response) (result DeleteRelayServiceConnectionSlotOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

package webapps

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeleteHybridConnectionSlotOperationResponse struct {
	HttpResponse *http.Response
}

// DeleteHybridConnectionSlot ...
func (c WebAppsClient) DeleteHybridConnectionSlot(ctx context.Context, id SlotHybridConnectionNamespaceRelayId) (result DeleteHybridConnectionSlotOperationResponse, err error) {
	req, err := c.preparerForDeleteHybridConnectionSlot(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "DeleteHybridConnectionSlot", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "DeleteHybridConnectionSlot", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForDeleteHybridConnectionSlot(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "DeleteHybridConnectionSlot", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForDeleteHybridConnectionSlot prepares the DeleteHybridConnectionSlot request.
func (c WebAppsClient) preparerForDeleteHybridConnectionSlot(ctx context.Context, id SlotHybridConnectionNamespaceRelayId) (*http.Request, error) {
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

// responderForDeleteHybridConnectionSlot handles the response to the DeleteHybridConnectionSlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForDeleteHybridConnectionSlot(resp *http.Response) (result DeleteHybridConnectionSlotOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

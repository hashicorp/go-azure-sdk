package webapps

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeleteVnetConnectionSlotOperationResponse struct {
	HttpResponse *http.Response
}

// DeleteVnetConnectionSlot ...
func (c WebAppsClient) DeleteVnetConnectionSlot(ctx context.Context, id SlotVirtualNetworkConnectionId) (result DeleteVnetConnectionSlotOperationResponse, err error) {
	req, err := c.preparerForDeleteVnetConnectionSlot(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "DeleteVnetConnectionSlot", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "DeleteVnetConnectionSlot", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForDeleteVnetConnectionSlot(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "DeleteVnetConnectionSlot", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForDeleteVnetConnectionSlot prepares the DeleteVnetConnectionSlot request.
func (c WebAppsClient) preparerForDeleteVnetConnectionSlot(ctx context.Context, id SlotVirtualNetworkConnectionId) (*http.Request, error) {
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

// responderForDeleteVnetConnectionSlot handles the response to the DeleteVnetConnectionSlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForDeleteVnetConnectionSlot(resp *http.Response) (result DeleteVnetConnectionSlotOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

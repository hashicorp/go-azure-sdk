package webapps

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetVnetConnectionSlotOperationResponse struct {
	HttpResponse *http.Response
	Model        *VnetInfoResource
}

// GetVnetConnectionSlot ...
func (c WebAppsClient) GetVnetConnectionSlot(ctx context.Context, id SlotVirtualNetworkConnectionId) (result GetVnetConnectionSlotOperationResponse, err error) {
	req, err := c.preparerForGetVnetConnectionSlot(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetVnetConnectionSlot", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetVnetConnectionSlot", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetVnetConnectionSlot(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetVnetConnectionSlot", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetVnetConnectionSlot prepares the GetVnetConnectionSlot request.
func (c WebAppsClient) preparerForGetVnetConnectionSlot(ctx context.Context, id SlotVirtualNetworkConnectionId) (*http.Request, error) {
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

// responderForGetVnetConnectionSlot handles the response to the GetVnetConnectionSlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForGetVnetConnectionSlot(resp *http.Response) (result GetVnetConnectionSlotOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

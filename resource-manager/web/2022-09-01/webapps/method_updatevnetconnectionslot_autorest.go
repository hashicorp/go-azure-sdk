package webapps

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateVnetConnectionSlotOperationResponse struct {
	HttpResponse *http.Response
	Model        *VnetInfoResource
}

// UpdateVnetConnectionSlot ...
func (c WebAppsClient) UpdateVnetConnectionSlot(ctx context.Context, id SlotVirtualNetworkConnectionId, input VnetInfoResource) (result UpdateVnetConnectionSlotOperationResponse, err error) {
	req, err := c.preparerForUpdateVnetConnectionSlot(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "UpdateVnetConnectionSlot", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "UpdateVnetConnectionSlot", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForUpdateVnetConnectionSlot(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "UpdateVnetConnectionSlot", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForUpdateVnetConnectionSlot prepares the UpdateVnetConnectionSlot request.
func (c WebAppsClient) preparerForUpdateVnetConnectionSlot(ctx context.Context, id SlotVirtualNetworkConnectionId, input VnetInfoResource) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForUpdateVnetConnectionSlot handles the response to the UpdateVnetConnectionSlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForUpdateVnetConnectionSlot(resp *http.Response) (result UpdateVnetConnectionSlotOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

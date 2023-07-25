package webapps

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateVnetConnectionGatewaySlotOperationResponse struct {
	HttpResponse *http.Response
	Model        *VnetGateway
}

// UpdateVnetConnectionGatewaySlot ...
func (c WebAppsClient) UpdateVnetConnectionGatewaySlot(ctx context.Context, id SlotVirtualNetworkConnectionGatewayId, input VnetGateway) (result UpdateVnetConnectionGatewaySlotOperationResponse, err error) {
	req, err := c.preparerForUpdateVnetConnectionGatewaySlot(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "UpdateVnetConnectionGatewaySlot", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "UpdateVnetConnectionGatewaySlot", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForUpdateVnetConnectionGatewaySlot(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "UpdateVnetConnectionGatewaySlot", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForUpdateVnetConnectionGatewaySlot prepares the UpdateVnetConnectionGatewaySlot request.
func (c WebAppsClient) preparerForUpdateVnetConnectionGatewaySlot(ctx context.Context, id SlotVirtualNetworkConnectionGatewayId, input VnetGateway) (*http.Request, error) {
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

// responderForUpdateVnetConnectionGatewaySlot handles the response to the UpdateVnetConnectionGatewaySlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForUpdateVnetConnectionGatewaySlot(resp *http.Response) (result UpdateVnetConnectionGatewaySlotOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

package webapps

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetVnetConnectionGatewaySlotOperationResponse struct {
	HttpResponse *http.Response
	Model        *VnetGateway
}

// GetVnetConnectionGatewaySlot ...
func (c WebAppsClient) GetVnetConnectionGatewaySlot(ctx context.Context, id SlotVirtualNetworkConnectionGatewayId) (result GetVnetConnectionGatewaySlotOperationResponse, err error) {
	req, err := c.preparerForGetVnetConnectionGatewaySlot(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetVnetConnectionGatewaySlot", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetVnetConnectionGatewaySlot", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetVnetConnectionGatewaySlot(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetVnetConnectionGatewaySlot", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetVnetConnectionGatewaySlot prepares the GetVnetConnectionGatewaySlot request.
func (c WebAppsClient) preparerForGetVnetConnectionGatewaySlot(ctx context.Context, id SlotVirtualNetworkConnectionGatewayId) (*http.Request, error) {
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

// responderForGetVnetConnectionGatewaySlot handles the response to the GetVnetConnectionGatewaySlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForGetVnetConnectionGatewaySlot(resp *http.Response) (result GetVnetConnectionGatewaySlotOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

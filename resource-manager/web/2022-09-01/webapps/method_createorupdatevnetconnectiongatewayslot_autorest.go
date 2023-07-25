package webapps

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateOrUpdateVnetConnectionGatewaySlotOperationResponse struct {
	HttpResponse *http.Response
	Model        *VnetGateway
}

// CreateOrUpdateVnetConnectionGatewaySlot ...
func (c WebAppsClient) CreateOrUpdateVnetConnectionGatewaySlot(ctx context.Context, id SlotVirtualNetworkConnectionGatewayId, input VnetGateway) (result CreateOrUpdateVnetConnectionGatewaySlotOperationResponse, err error) {
	req, err := c.preparerForCreateOrUpdateVnetConnectionGatewaySlot(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "CreateOrUpdateVnetConnectionGatewaySlot", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "CreateOrUpdateVnetConnectionGatewaySlot", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForCreateOrUpdateVnetConnectionGatewaySlot(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "CreateOrUpdateVnetConnectionGatewaySlot", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForCreateOrUpdateVnetConnectionGatewaySlot prepares the CreateOrUpdateVnetConnectionGatewaySlot request.
func (c WebAppsClient) preparerForCreateOrUpdateVnetConnectionGatewaySlot(ctx context.Context, id SlotVirtualNetworkConnectionGatewayId, input VnetGateway) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForCreateOrUpdateVnetConnectionGatewaySlot handles the response to the CreateOrUpdateVnetConnectionGatewaySlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForCreateOrUpdateVnetConnectionGatewaySlot(resp *http.Response) (result CreateOrUpdateVnetConnectionGatewaySlotOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

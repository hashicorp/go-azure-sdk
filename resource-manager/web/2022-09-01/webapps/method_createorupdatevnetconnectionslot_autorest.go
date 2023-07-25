package webapps

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateOrUpdateVnetConnectionSlotOperationResponse struct {
	HttpResponse *http.Response
	Model        *VnetInfoResource
}

// CreateOrUpdateVnetConnectionSlot ...
func (c WebAppsClient) CreateOrUpdateVnetConnectionSlot(ctx context.Context, id SlotVirtualNetworkConnectionId, input VnetInfoResource) (result CreateOrUpdateVnetConnectionSlotOperationResponse, err error) {
	req, err := c.preparerForCreateOrUpdateVnetConnectionSlot(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "CreateOrUpdateVnetConnectionSlot", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "CreateOrUpdateVnetConnectionSlot", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForCreateOrUpdateVnetConnectionSlot(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "CreateOrUpdateVnetConnectionSlot", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForCreateOrUpdateVnetConnectionSlot prepares the CreateOrUpdateVnetConnectionSlot request.
func (c WebAppsClient) preparerForCreateOrUpdateVnetConnectionSlot(ctx context.Context, id SlotVirtualNetworkConnectionId, input VnetInfoResource) (*http.Request, error) {
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

// responderForCreateOrUpdateVnetConnectionSlot handles the response to the CreateOrUpdateVnetConnectionSlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForCreateOrUpdateVnetConnectionSlot(resp *http.Response) (result CreateOrUpdateVnetConnectionSlotOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

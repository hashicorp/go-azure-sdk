package webapps

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AddPremierAddOnSlotOperationResponse struct {
	HttpResponse *http.Response
	Model        *PremierAddOn
}

// AddPremierAddOnSlot ...
func (c WebAppsClient) AddPremierAddOnSlot(ctx context.Context, id SlotPremierAddonId, input PremierAddOn) (result AddPremierAddOnSlotOperationResponse, err error) {
	req, err := c.preparerForAddPremierAddOnSlot(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "AddPremierAddOnSlot", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "AddPremierAddOnSlot", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForAddPremierAddOnSlot(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "AddPremierAddOnSlot", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForAddPremierAddOnSlot prepares the AddPremierAddOnSlot request.
func (c WebAppsClient) preparerForAddPremierAddOnSlot(ctx context.Context, id SlotPremierAddonId, input PremierAddOn) (*http.Request, error) {
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

// responderForAddPremierAddOnSlot handles the response to the AddPremierAddOnSlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForAddPremierAddOnSlot(resp *http.Response) (result AddPremierAddOnSlotOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

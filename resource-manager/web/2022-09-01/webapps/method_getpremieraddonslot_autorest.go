package webapps

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetPremierAddOnSlotOperationResponse struct {
	HttpResponse *http.Response
	Model        *PremierAddOn
}

// GetPremierAddOnSlot ...
func (c WebAppsClient) GetPremierAddOnSlot(ctx context.Context, id SlotPremierAddonId) (result GetPremierAddOnSlotOperationResponse, err error) {
	req, err := c.preparerForGetPremierAddOnSlot(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetPremierAddOnSlot", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetPremierAddOnSlot", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetPremierAddOnSlot(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetPremierAddOnSlot", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetPremierAddOnSlot prepares the GetPremierAddOnSlot request.
func (c WebAppsClient) preparerForGetPremierAddOnSlot(ctx context.Context, id SlotPremierAddonId) (*http.Request, error) {
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

// responderForGetPremierAddOnSlot handles the response to the GetPremierAddOnSlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForGetPremierAddOnSlot(resp *http.Response) (result GetPremierAddOnSlotOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

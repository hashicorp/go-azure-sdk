package webapps

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeletePremierAddOnSlotOperationResponse struct {
	HttpResponse *http.Response
}

// DeletePremierAddOnSlot ...
func (c WebAppsClient) DeletePremierAddOnSlot(ctx context.Context, id SlotPremierAddonId) (result DeletePremierAddOnSlotOperationResponse, err error) {
	req, err := c.preparerForDeletePremierAddOnSlot(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "DeletePremierAddOnSlot", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "DeletePremierAddOnSlot", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForDeletePremierAddOnSlot(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "DeletePremierAddOnSlot", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForDeletePremierAddOnSlot prepares the DeletePremierAddOnSlot request.
func (c WebAppsClient) preparerForDeletePremierAddOnSlot(ctx context.Context, id SlotPremierAddonId) (*http.Request, error) {
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

// responderForDeletePremierAddOnSlot handles the response to the DeletePremierAddOnSlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForDeletePremierAddOnSlot(resp *http.Response) (result DeletePremierAddOnSlotOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

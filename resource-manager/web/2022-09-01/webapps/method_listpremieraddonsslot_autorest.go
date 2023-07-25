package webapps

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListPremierAddOnsSlotOperationResponse struct {
	HttpResponse *http.Response
	Model        *PremierAddOn
}

// ListPremierAddOnsSlot ...
func (c WebAppsClient) ListPremierAddOnsSlot(ctx context.Context, id SlotId) (result ListPremierAddOnsSlotOperationResponse, err error) {
	req, err := c.preparerForListPremierAddOnsSlot(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListPremierAddOnsSlot", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListPremierAddOnsSlot", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForListPremierAddOnsSlot(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListPremierAddOnsSlot", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForListPremierAddOnsSlot prepares the ListPremierAddOnsSlot request.
func (c WebAppsClient) preparerForListPremierAddOnsSlot(ctx context.Context, id SlotId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/premierAddons", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForListPremierAddOnsSlot handles the response to the ListPremierAddOnsSlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForListPremierAddOnsSlot(resp *http.Response) (result ListPremierAddOnsSlotOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

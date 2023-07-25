package webapps

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetSiteExtensionSlotOperationResponse struct {
	HttpResponse *http.Response
	Model        *SiteExtensionInfo
}

// GetSiteExtensionSlot ...
func (c WebAppsClient) GetSiteExtensionSlot(ctx context.Context, id SlotSiteExtensionId) (result GetSiteExtensionSlotOperationResponse, err error) {
	req, err := c.preparerForGetSiteExtensionSlot(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetSiteExtensionSlot", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetSiteExtensionSlot", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetSiteExtensionSlot(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetSiteExtensionSlot", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetSiteExtensionSlot prepares the GetSiteExtensionSlot request.
func (c WebAppsClient) preparerForGetSiteExtensionSlot(ctx context.Context, id SlotSiteExtensionId) (*http.Request, error) {
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

// responderForGetSiteExtensionSlot handles the response to the GetSiteExtensionSlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForGetSiteExtensionSlot(resp *http.Response) (result GetSiteExtensionSlotOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

package webapps

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeleteSiteExtensionSlotOperationResponse struct {
	HttpResponse *http.Response
}

// DeleteSiteExtensionSlot ...
func (c WebAppsClient) DeleteSiteExtensionSlot(ctx context.Context, id SlotSiteExtensionId) (result DeleteSiteExtensionSlotOperationResponse, err error) {
	req, err := c.preparerForDeleteSiteExtensionSlot(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "DeleteSiteExtensionSlot", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "DeleteSiteExtensionSlot", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForDeleteSiteExtensionSlot(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "DeleteSiteExtensionSlot", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForDeleteSiteExtensionSlot prepares the DeleteSiteExtensionSlot request.
func (c WebAppsClient) preparerForDeleteSiteExtensionSlot(ctx context.Context, id SlotSiteExtensionId) (*http.Request, error) {
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

// responderForDeleteSiteExtensionSlot handles the response to the DeleteSiteExtensionSlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForDeleteSiteExtensionSlot(resp *http.Response) (result DeleteSiteExtensionSlotOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusNoContent),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

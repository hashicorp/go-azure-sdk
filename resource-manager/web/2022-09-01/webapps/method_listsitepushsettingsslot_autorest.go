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

type ListSitePushSettingsSlotOperationResponse struct {
	HttpResponse *http.Response
	Model        *PushSettings
}

// ListSitePushSettingsSlot ...
func (c WebAppsClient) ListSitePushSettingsSlot(ctx context.Context, id SlotId) (result ListSitePushSettingsSlotOperationResponse, err error) {
	req, err := c.preparerForListSitePushSettingsSlot(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListSitePushSettingsSlot", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListSitePushSettingsSlot", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForListSitePushSettingsSlot(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListSitePushSettingsSlot", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForListSitePushSettingsSlot prepares the ListSitePushSettingsSlot request.
func (c WebAppsClient) preparerForListSitePushSettingsSlot(ctx context.Context, id SlotId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/config/pushsettings/list", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForListSitePushSettingsSlot handles the response to the ListSitePushSettingsSlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForListSitePushSettingsSlot(resp *http.Response) (result ListSitePushSettingsSlotOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

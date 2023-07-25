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

type UpdateSitePushSettingsSlotOperationResponse struct {
	HttpResponse *http.Response
	Model        *PushSettings
}

// UpdateSitePushSettingsSlot ...
func (c WebAppsClient) UpdateSitePushSettingsSlot(ctx context.Context, id SlotId, input PushSettings) (result UpdateSitePushSettingsSlotOperationResponse, err error) {
	req, err := c.preparerForUpdateSitePushSettingsSlot(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "UpdateSitePushSettingsSlot", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "UpdateSitePushSettingsSlot", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForUpdateSitePushSettingsSlot(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "UpdateSitePushSettingsSlot", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForUpdateSitePushSettingsSlot prepares the UpdateSitePushSettingsSlot request.
func (c WebAppsClient) preparerForUpdateSitePushSettingsSlot(ctx context.Context, id SlotId, input PushSettings) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/config/pushsettings", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForUpdateSitePushSettingsSlot handles the response to the UpdateSitePushSettingsSlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForUpdateSitePushSettingsSlot(resp *http.Response) (result UpdateSitePushSettingsSlotOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

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

type UpdateAuthSettingsSlotOperationResponse struct {
	HttpResponse *http.Response
	Model        *SiteAuthSettings
}

// UpdateAuthSettingsSlot ...
func (c WebAppsClient) UpdateAuthSettingsSlot(ctx context.Context, id SlotId, input SiteAuthSettings) (result UpdateAuthSettingsSlotOperationResponse, err error) {
	req, err := c.preparerForUpdateAuthSettingsSlot(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "UpdateAuthSettingsSlot", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "UpdateAuthSettingsSlot", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForUpdateAuthSettingsSlot(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "UpdateAuthSettingsSlot", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForUpdateAuthSettingsSlot prepares the UpdateAuthSettingsSlot request.
func (c WebAppsClient) preparerForUpdateAuthSettingsSlot(ctx context.Context, id SlotId, input SiteAuthSettings) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/config/authsettings", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForUpdateAuthSettingsSlot handles the response to the UpdateAuthSettingsSlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForUpdateAuthSettingsSlot(resp *http.Response) (result UpdateAuthSettingsSlotOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

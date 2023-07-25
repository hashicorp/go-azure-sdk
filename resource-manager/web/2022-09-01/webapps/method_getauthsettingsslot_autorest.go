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

type GetAuthSettingsSlotOperationResponse struct {
	HttpResponse *http.Response
	Model        *SiteAuthSettings
}

// GetAuthSettingsSlot ...
func (c WebAppsClient) GetAuthSettingsSlot(ctx context.Context, id SlotId) (result GetAuthSettingsSlotOperationResponse, err error) {
	req, err := c.preparerForGetAuthSettingsSlot(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetAuthSettingsSlot", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetAuthSettingsSlot", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetAuthSettingsSlot(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetAuthSettingsSlot", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetAuthSettingsSlot prepares the GetAuthSettingsSlot request.
func (c WebAppsClient) preparerForGetAuthSettingsSlot(ctx context.Context, id SlotId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/config/authsettings/list", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForGetAuthSettingsSlot handles the response to the GetAuthSettingsSlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForGetAuthSettingsSlot(resp *http.Response) (result GetAuthSettingsSlotOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

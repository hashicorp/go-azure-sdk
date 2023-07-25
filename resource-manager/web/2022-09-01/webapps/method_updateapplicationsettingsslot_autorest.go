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

type UpdateApplicationSettingsSlotOperationResponse struct {
	HttpResponse *http.Response
	Model        *StringDictionary
}

// UpdateApplicationSettingsSlot ...
func (c WebAppsClient) UpdateApplicationSettingsSlot(ctx context.Context, id SlotId, input StringDictionary) (result UpdateApplicationSettingsSlotOperationResponse, err error) {
	req, err := c.preparerForUpdateApplicationSettingsSlot(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "UpdateApplicationSettingsSlot", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "UpdateApplicationSettingsSlot", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForUpdateApplicationSettingsSlot(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "UpdateApplicationSettingsSlot", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForUpdateApplicationSettingsSlot prepares the UpdateApplicationSettingsSlot request.
func (c WebAppsClient) preparerForUpdateApplicationSettingsSlot(ctx context.Context, id SlotId, input StringDictionary) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/config/appSettings", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForUpdateApplicationSettingsSlot handles the response to the UpdateApplicationSettingsSlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForUpdateApplicationSettingsSlot(resp *http.Response) (result UpdateApplicationSettingsSlotOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

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

type UpdateAuthSettingsV2SlotOperationResponse struct {
	HttpResponse *http.Response
	Model        *SiteAuthSettingsV2
}

// UpdateAuthSettingsV2Slot ...
func (c WebAppsClient) UpdateAuthSettingsV2Slot(ctx context.Context, id SlotId, input SiteAuthSettingsV2) (result UpdateAuthSettingsV2SlotOperationResponse, err error) {
	req, err := c.preparerForUpdateAuthSettingsV2Slot(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "UpdateAuthSettingsV2Slot", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "UpdateAuthSettingsV2Slot", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForUpdateAuthSettingsV2Slot(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "UpdateAuthSettingsV2Slot", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForUpdateAuthSettingsV2Slot prepares the UpdateAuthSettingsV2Slot request.
func (c WebAppsClient) preparerForUpdateAuthSettingsV2Slot(ctx context.Context, id SlotId, input SiteAuthSettingsV2) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/config/authsettingsV2", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForUpdateAuthSettingsV2Slot handles the response to the UpdateAuthSettingsV2Slot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForUpdateAuthSettingsV2Slot(resp *http.Response) (result UpdateAuthSettingsV2SlotOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

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

type GetAuthSettingsV2SlotOperationResponse struct {
	HttpResponse *http.Response
	Model        *SiteAuthSettingsV2
}

// GetAuthSettingsV2Slot ...
func (c WebAppsClient) GetAuthSettingsV2Slot(ctx context.Context, id SlotId) (result GetAuthSettingsV2SlotOperationResponse, err error) {
	req, err := c.preparerForGetAuthSettingsV2Slot(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetAuthSettingsV2Slot", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetAuthSettingsV2Slot", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetAuthSettingsV2Slot(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetAuthSettingsV2Slot", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetAuthSettingsV2Slot prepares the GetAuthSettingsV2Slot request.
func (c WebAppsClient) preparerForGetAuthSettingsV2Slot(ctx context.Context, id SlotId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/config/authsettingsV2/list", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForGetAuthSettingsV2Slot handles the response to the GetAuthSettingsV2Slot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForGetAuthSettingsV2Slot(resp *http.Response) (result GetAuthSettingsV2SlotOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

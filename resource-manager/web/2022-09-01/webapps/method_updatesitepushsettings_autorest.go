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

type UpdateSitePushSettingsOperationResponse struct {
	HttpResponse *http.Response
	Model        *PushSettings
}

// UpdateSitePushSettings ...
func (c WebAppsClient) UpdateSitePushSettings(ctx context.Context, id SiteId, input PushSettings) (result UpdateSitePushSettingsOperationResponse, err error) {
	req, err := c.preparerForUpdateSitePushSettings(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "UpdateSitePushSettings", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "UpdateSitePushSettings", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForUpdateSitePushSettings(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "UpdateSitePushSettings", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForUpdateSitePushSettings prepares the UpdateSitePushSettings request.
func (c WebAppsClient) preparerForUpdateSitePushSettings(ctx context.Context, id SiteId, input PushSettings) (*http.Request, error) {
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

// responderForUpdateSitePushSettings handles the response to the UpdateSitePushSettings request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForUpdateSitePushSettings(resp *http.Response) (result UpdateSitePushSettingsOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

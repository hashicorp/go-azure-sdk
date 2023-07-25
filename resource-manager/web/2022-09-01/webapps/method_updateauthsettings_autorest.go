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

type UpdateAuthSettingsOperationResponse struct {
	HttpResponse *http.Response
	Model        *SiteAuthSettings
}

// UpdateAuthSettings ...
func (c WebAppsClient) UpdateAuthSettings(ctx context.Context, id SiteId, input SiteAuthSettings) (result UpdateAuthSettingsOperationResponse, err error) {
	req, err := c.preparerForUpdateAuthSettings(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "UpdateAuthSettings", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "UpdateAuthSettings", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForUpdateAuthSettings(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "UpdateAuthSettings", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForUpdateAuthSettings prepares the UpdateAuthSettings request.
func (c WebAppsClient) preparerForUpdateAuthSettings(ctx context.Context, id SiteId, input SiteAuthSettings) (*http.Request, error) {
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

// responderForUpdateAuthSettings handles the response to the UpdateAuthSettings request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForUpdateAuthSettings(resp *http.Response) (result UpdateAuthSettingsOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

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

type GetAuthSettingsOperationResponse struct {
	HttpResponse *http.Response
	Model        *SiteAuthSettings
}

// GetAuthSettings ...
func (c WebAppsClient) GetAuthSettings(ctx context.Context, id SiteId) (result GetAuthSettingsOperationResponse, err error) {
	req, err := c.preparerForGetAuthSettings(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetAuthSettings", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetAuthSettings", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetAuthSettings(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetAuthSettings", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetAuthSettings prepares the GetAuthSettings request.
func (c WebAppsClient) preparerForGetAuthSettings(ctx context.Context, id SiteId) (*http.Request, error) {
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

// responderForGetAuthSettings handles the response to the GetAuthSettings request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForGetAuthSettings(resp *http.Response) (result GetAuthSettingsOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

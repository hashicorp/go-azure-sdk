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

type ListSitePushSettingsOperationResponse struct {
	HttpResponse *http.Response
	Model        *PushSettings
}

// ListSitePushSettings ...
func (c WebAppsClient) ListSitePushSettings(ctx context.Context, id SiteId) (result ListSitePushSettingsOperationResponse, err error) {
	req, err := c.preparerForListSitePushSettings(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListSitePushSettings", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListSitePushSettings", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForListSitePushSettings(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListSitePushSettings", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForListSitePushSettings prepares the ListSitePushSettings request.
func (c WebAppsClient) preparerForListSitePushSettings(ctx context.Context, id SiteId) (*http.Request, error) {
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

// responderForListSitePushSettings handles the response to the ListSitePushSettings request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForListSitePushSettings(resp *http.Response) (result ListSitePushSettingsOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

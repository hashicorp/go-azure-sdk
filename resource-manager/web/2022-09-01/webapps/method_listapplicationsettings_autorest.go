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

type ListApplicationSettingsOperationResponse struct {
	HttpResponse *http.Response
	Model        *StringDictionary
}

// ListApplicationSettings ...
func (c WebAppsClient) ListApplicationSettings(ctx context.Context, id SiteId) (result ListApplicationSettingsOperationResponse, err error) {
	req, err := c.preparerForListApplicationSettings(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListApplicationSettings", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListApplicationSettings", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForListApplicationSettings(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListApplicationSettings", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForListApplicationSettings prepares the ListApplicationSettings request.
func (c WebAppsClient) preparerForListApplicationSettings(ctx context.Context, id SiteId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/config/appSettings/list", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForListApplicationSettings handles the response to the ListApplicationSettings request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForListApplicationSettings(resp *http.Response) (result ListApplicationSettingsOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

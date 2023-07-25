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

type UpdateApplicationSettingsOperationResponse struct {
	HttpResponse *http.Response
	Model        *StringDictionary
}

// UpdateApplicationSettings ...
func (c WebAppsClient) UpdateApplicationSettings(ctx context.Context, id SiteId, input StringDictionary) (result UpdateApplicationSettingsOperationResponse, err error) {
	req, err := c.preparerForUpdateApplicationSettings(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "UpdateApplicationSettings", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "UpdateApplicationSettings", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForUpdateApplicationSettings(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "UpdateApplicationSettings", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForUpdateApplicationSettings prepares the UpdateApplicationSettings request.
func (c WebAppsClient) preparerForUpdateApplicationSettings(ctx context.Context, id SiteId, input StringDictionary) (*http.Request, error) {
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

// responderForUpdateApplicationSettings handles the response to the UpdateApplicationSettings request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForUpdateApplicationSettings(resp *http.Response) (result UpdateApplicationSettingsOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

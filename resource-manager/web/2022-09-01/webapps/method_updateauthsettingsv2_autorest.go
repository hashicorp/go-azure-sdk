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

type UpdateAuthSettingsV2OperationResponse struct {
	HttpResponse *http.Response
	Model        *SiteAuthSettingsV2
}

// UpdateAuthSettingsV2 ...
func (c WebAppsClient) UpdateAuthSettingsV2(ctx context.Context, id SiteId, input SiteAuthSettingsV2) (result UpdateAuthSettingsV2OperationResponse, err error) {
	req, err := c.preparerForUpdateAuthSettingsV2(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "UpdateAuthSettingsV2", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "UpdateAuthSettingsV2", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForUpdateAuthSettingsV2(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "UpdateAuthSettingsV2", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForUpdateAuthSettingsV2 prepares the UpdateAuthSettingsV2 request.
func (c WebAppsClient) preparerForUpdateAuthSettingsV2(ctx context.Context, id SiteId, input SiteAuthSettingsV2) (*http.Request, error) {
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

// responderForUpdateAuthSettingsV2 handles the response to the UpdateAuthSettingsV2 request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForUpdateAuthSettingsV2(resp *http.Response) (result UpdateAuthSettingsV2OperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

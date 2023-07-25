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

type GetAuthSettingsV2OperationResponse struct {
	HttpResponse *http.Response
	Model        *SiteAuthSettingsV2
}

// GetAuthSettingsV2 ...
func (c WebAppsClient) GetAuthSettingsV2(ctx context.Context, id SiteId) (result GetAuthSettingsV2OperationResponse, err error) {
	req, err := c.preparerForGetAuthSettingsV2(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetAuthSettingsV2", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetAuthSettingsV2", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetAuthSettingsV2(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetAuthSettingsV2", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetAuthSettingsV2 prepares the GetAuthSettingsV2 request.
func (c WebAppsClient) preparerForGetAuthSettingsV2(ctx context.Context, id SiteId) (*http.Request, error) {
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

// responderForGetAuthSettingsV2 handles the response to the GetAuthSettingsV2 request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForGetAuthSettingsV2(resp *http.Response) (result GetAuthSettingsV2OperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

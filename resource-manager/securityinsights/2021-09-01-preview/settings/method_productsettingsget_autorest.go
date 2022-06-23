package settings

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProductSettingsGetOperationResponse struct {
	HttpResponse *http.Response
	Model        *Settings
}

// ProductSettingsGet ...
func (c SettingsClient) ProductSettingsGet(ctx context.Context, id SettingId) (result ProductSettingsGetOperationResponse, err error) {
	req, err := c.preparerForProductSettingsGet(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "settings.SettingsClient", "ProductSettingsGet", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "settings.SettingsClient", "ProductSettingsGet", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForProductSettingsGet(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "settings.SettingsClient", "ProductSettingsGet", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForProductSettingsGet prepares the ProductSettingsGet request.
func (c SettingsClient) preparerForProductSettingsGet(ctx context.Context, id SettingId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForProductSettingsGet handles the response to the ProductSettingsGet request. The method always
// closes the http.Response Body.
func (c SettingsClient) responderForProductSettingsGet(resp *http.Response) (result ProductSettingsGetOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp
	return
}

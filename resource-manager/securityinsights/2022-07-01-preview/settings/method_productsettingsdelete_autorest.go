package settings

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProductSettingsDeleteOperationResponse struct {
	HttpResponse *http.Response
}

// ProductSettingsDelete ...
func (c SettingsClient) ProductSettingsDelete(ctx context.Context, id SettingId) (result ProductSettingsDeleteOperationResponse, err error) {
	req, err := c.preparerForProductSettingsDelete(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "settings.SettingsClient", "ProductSettingsDelete", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "settings.SettingsClient", "ProductSettingsDelete", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForProductSettingsDelete(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "settings.SettingsClient", "ProductSettingsDelete", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForProductSettingsDelete prepares the ProductSettingsDelete request.
func (c SettingsClient) preparerForProductSettingsDelete(ctx context.Context, id SettingId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsDelete(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForProductSettingsDelete handles the response to the ProductSettingsDelete request. The method always
// closes the http.Response Body.
func (c SettingsClient) responderForProductSettingsDelete(resp *http.Response) (result ProductSettingsDeleteOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusNoContent, http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

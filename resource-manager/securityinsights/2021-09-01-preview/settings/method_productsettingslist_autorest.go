package settings

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProductSettingsListOperationResponse struct {
	HttpResponse *http.Response
	Model        *SettingList
}

// ProductSettingsList ...
func (c SettingsClient) ProductSettingsList(ctx context.Context, id WorkspaceId) (result ProductSettingsListOperationResponse, err error) {
	req, err := c.preparerForProductSettingsList(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "settings.SettingsClient", "ProductSettingsList", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "settings.SettingsClient", "ProductSettingsList", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForProductSettingsList(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "settings.SettingsClient", "ProductSettingsList", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForProductSettingsList prepares the ProductSettingsList request.
func (c SettingsClient) preparerForProductSettingsList(ctx context.Context, id WorkspaceId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/providers/Microsoft.SecurityInsights/settings", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForProductSettingsList handles the response to the ProductSettingsList request. The method always
// closes the http.Response Body.
func (c SettingsClient) responderForProductSettingsList(resp *http.Response) (result ProductSettingsListOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp
	return
}

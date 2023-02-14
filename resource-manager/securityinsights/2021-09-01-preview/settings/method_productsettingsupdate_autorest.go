package settings

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProductSettingsUpdateOperationResponse struct {
	HttpResponse *http.Response
	Model        *Settings
}

// ProductSettingsUpdate ...
func (c SettingsClient) ProductSettingsUpdate(ctx context.Context, id SettingId, input Settings) (result ProductSettingsUpdateOperationResponse, err error) {
	req, err := c.preparerForProductSettingsUpdate(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "settings.SettingsClient", "ProductSettingsUpdate", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "settings.SettingsClient", "ProductSettingsUpdate", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForProductSettingsUpdate(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "settings.SettingsClient", "ProductSettingsUpdate", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForProductSettingsUpdate prepares the ProductSettingsUpdate request.
func (c SettingsClient) preparerForProductSettingsUpdate(ctx context.Context, id SettingId, input Settings) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForProductSettingsUpdate handles the response to the ProductSettingsUpdate request. The method always
// closes the http.Response Body.
func (c SettingsClient) responderForProductSettingsUpdate(resp *http.Response) (result ProductSettingsUpdateOperationResponse, err error) {
	var respObj json.RawMessage
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&respObj),
		autorest.ByClosing())
	result.HttpResponse = resp
	model, err := unmarshalSettingsImplementation(respObj)
	if err != nil {
		return
	}
	result.Model = &model
	return
}

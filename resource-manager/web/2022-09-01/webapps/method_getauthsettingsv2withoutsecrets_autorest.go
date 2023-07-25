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

type GetAuthSettingsV2WithoutSecretsOperationResponse struct {
	HttpResponse *http.Response
	Model        *SiteAuthSettingsV2
}

// GetAuthSettingsV2WithoutSecrets ...
func (c WebAppsClient) GetAuthSettingsV2WithoutSecrets(ctx context.Context, id SiteId) (result GetAuthSettingsV2WithoutSecretsOperationResponse, err error) {
	req, err := c.preparerForGetAuthSettingsV2WithoutSecrets(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetAuthSettingsV2WithoutSecrets", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetAuthSettingsV2WithoutSecrets", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetAuthSettingsV2WithoutSecrets(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetAuthSettingsV2WithoutSecrets", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetAuthSettingsV2WithoutSecrets prepares the GetAuthSettingsV2WithoutSecrets request.
func (c WebAppsClient) preparerForGetAuthSettingsV2WithoutSecrets(ctx context.Context, id SiteId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/config/authsettingsV2", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForGetAuthSettingsV2WithoutSecrets handles the response to the GetAuthSettingsV2WithoutSecrets request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForGetAuthSettingsV2WithoutSecrets(resp *http.Response) (result GetAuthSettingsV2WithoutSecretsOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

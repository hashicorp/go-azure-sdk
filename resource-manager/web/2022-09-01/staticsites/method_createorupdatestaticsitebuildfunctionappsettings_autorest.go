package staticsites

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateOrUpdateStaticSiteBuildFunctionAppSettingsOperationResponse struct {
	HttpResponse *http.Response
	Model        *StringDictionary
}

// CreateOrUpdateStaticSiteBuildFunctionAppSettings ...
func (c StaticSitesClient) CreateOrUpdateStaticSiteBuildFunctionAppSettings(ctx context.Context, id BuildId, input StringDictionary) (result CreateOrUpdateStaticSiteBuildFunctionAppSettingsOperationResponse, err error) {
	req, err := c.preparerForCreateOrUpdateStaticSiteBuildFunctionAppSettings(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "CreateOrUpdateStaticSiteBuildFunctionAppSettings", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "CreateOrUpdateStaticSiteBuildFunctionAppSettings", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForCreateOrUpdateStaticSiteBuildFunctionAppSettings(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "CreateOrUpdateStaticSiteBuildFunctionAppSettings", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForCreateOrUpdateStaticSiteBuildFunctionAppSettings prepares the CreateOrUpdateStaticSiteBuildFunctionAppSettings request.
func (c StaticSitesClient) preparerForCreateOrUpdateStaticSiteBuildFunctionAppSettings(ctx context.Context, id BuildId, input StringDictionary) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/config/functionAppSettings", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForCreateOrUpdateStaticSiteBuildFunctionAppSettings handles the response to the CreateOrUpdateStaticSiteBuildFunctionAppSettings request. The method always
// closes the http.Response Body.
func (c StaticSitesClient) responderForCreateOrUpdateStaticSiteBuildFunctionAppSettings(resp *http.Response) (result CreateOrUpdateStaticSiteBuildFunctionAppSettingsOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

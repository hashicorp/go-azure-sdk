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

type CreateOrUpdateStaticSiteFunctionAppSettingsOperationResponse struct {
	HttpResponse *http.Response
	Model        *StringDictionary
}

// CreateOrUpdateStaticSiteFunctionAppSettings ...
func (c StaticSitesClient) CreateOrUpdateStaticSiteFunctionAppSettings(ctx context.Context, id StaticSiteId, input StringDictionary) (result CreateOrUpdateStaticSiteFunctionAppSettingsOperationResponse, err error) {
	req, err := c.preparerForCreateOrUpdateStaticSiteFunctionAppSettings(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "CreateOrUpdateStaticSiteFunctionAppSettings", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "CreateOrUpdateStaticSiteFunctionAppSettings", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForCreateOrUpdateStaticSiteFunctionAppSettings(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "CreateOrUpdateStaticSiteFunctionAppSettings", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForCreateOrUpdateStaticSiteFunctionAppSettings prepares the CreateOrUpdateStaticSiteFunctionAppSettings request.
func (c StaticSitesClient) preparerForCreateOrUpdateStaticSiteFunctionAppSettings(ctx context.Context, id StaticSiteId, input StringDictionary) (*http.Request, error) {
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

// responderForCreateOrUpdateStaticSiteFunctionAppSettings handles the response to the CreateOrUpdateStaticSiteFunctionAppSettings request. The method always
// closes the http.Response Body.
func (c StaticSitesClient) responderForCreateOrUpdateStaticSiteFunctionAppSettings(resp *http.Response) (result CreateOrUpdateStaticSiteFunctionAppSettingsOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

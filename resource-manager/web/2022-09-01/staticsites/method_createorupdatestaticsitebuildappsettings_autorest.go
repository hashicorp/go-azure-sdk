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

type CreateOrUpdateStaticSiteBuildAppSettingsOperationResponse struct {
	HttpResponse *http.Response
	Model        *StringDictionary
}

// CreateOrUpdateStaticSiteBuildAppSettings ...
func (c StaticSitesClient) CreateOrUpdateStaticSiteBuildAppSettings(ctx context.Context, id BuildId, input StringDictionary) (result CreateOrUpdateStaticSiteBuildAppSettingsOperationResponse, err error) {
	req, err := c.preparerForCreateOrUpdateStaticSiteBuildAppSettings(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "CreateOrUpdateStaticSiteBuildAppSettings", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "CreateOrUpdateStaticSiteBuildAppSettings", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForCreateOrUpdateStaticSiteBuildAppSettings(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "CreateOrUpdateStaticSiteBuildAppSettings", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForCreateOrUpdateStaticSiteBuildAppSettings prepares the CreateOrUpdateStaticSiteBuildAppSettings request.
func (c StaticSitesClient) preparerForCreateOrUpdateStaticSiteBuildAppSettings(ctx context.Context, id BuildId, input StringDictionary) (*http.Request, error) {
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

// responderForCreateOrUpdateStaticSiteBuildAppSettings handles the response to the CreateOrUpdateStaticSiteBuildAppSettings request. The method always
// closes the http.Response Body.
func (c StaticSitesClient) responderForCreateOrUpdateStaticSiteBuildAppSettings(resp *http.Response) (result CreateOrUpdateStaticSiteBuildAppSettingsOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

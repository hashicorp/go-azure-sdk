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

type CreateOrUpdateStaticSiteAppSettingsOperationResponse struct {
	HttpResponse *http.Response
	Model        *StringDictionary
}

// CreateOrUpdateStaticSiteAppSettings ...
func (c StaticSitesClient) CreateOrUpdateStaticSiteAppSettings(ctx context.Context, id StaticSiteId, input StringDictionary) (result CreateOrUpdateStaticSiteAppSettingsOperationResponse, err error) {
	req, err := c.preparerForCreateOrUpdateStaticSiteAppSettings(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "CreateOrUpdateStaticSiteAppSettings", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "CreateOrUpdateStaticSiteAppSettings", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForCreateOrUpdateStaticSiteAppSettings(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "CreateOrUpdateStaticSiteAppSettings", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForCreateOrUpdateStaticSiteAppSettings prepares the CreateOrUpdateStaticSiteAppSettings request.
func (c StaticSitesClient) preparerForCreateOrUpdateStaticSiteAppSettings(ctx context.Context, id StaticSiteId, input StringDictionary) (*http.Request, error) {
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

// responderForCreateOrUpdateStaticSiteAppSettings handles the response to the CreateOrUpdateStaticSiteAppSettings request. The method always
// closes the http.Response Body.
func (c StaticSitesClient) responderForCreateOrUpdateStaticSiteAppSettings(resp *http.Response) (result CreateOrUpdateStaticSiteAppSettingsOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

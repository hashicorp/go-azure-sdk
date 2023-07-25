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

type ListStaticSiteFunctionAppSettingsOperationResponse struct {
	HttpResponse *http.Response
	Model        *StringDictionary
}

// ListStaticSiteFunctionAppSettings ...
func (c StaticSitesClient) ListStaticSiteFunctionAppSettings(ctx context.Context, id StaticSiteId) (result ListStaticSiteFunctionAppSettingsOperationResponse, err error) {
	req, err := c.preparerForListStaticSiteFunctionAppSettings(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "ListStaticSiteFunctionAppSettings", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "ListStaticSiteFunctionAppSettings", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForListStaticSiteFunctionAppSettings(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "ListStaticSiteFunctionAppSettings", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForListStaticSiteFunctionAppSettings prepares the ListStaticSiteFunctionAppSettings request.
func (c StaticSitesClient) preparerForListStaticSiteFunctionAppSettings(ctx context.Context, id StaticSiteId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/listFunctionAppSettings", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForListStaticSiteFunctionAppSettings handles the response to the ListStaticSiteFunctionAppSettings request. The method always
// closes the http.Response Body.
func (c StaticSitesClient) responderForListStaticSiteFunctionAppSettings(resp *http.Response) (result ListStaticSiteFunctionAppSettingsOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

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

type ListStaticSiteBuildFunctionAppSettingsOperationResponse struct {
	HttpResponse *http.Response
	Model        *StringDictionary
}

// ListStaticSiteBuildFunctionAppSettings ...
func (c StaticSitesClient) ListStaticSiteBuildFunctionAppSettings(ctx context.Context, id BuildId) (result ListStaticSiteBuildFunctionAppSettingsOperationResponse, err error) {
	req, err := c.preparerForListStaticSiteBuildFunctionAppSettings(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "ListStaticSiteBuildFunctionAppSettings", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "ListStaticSiteBuildFunctionAppSettings", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForListStaticSiteBuildFunctionAppSettings(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "ListStaticSiteBuildFunctionAppSettings", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForListStaticSiteBuildFunctionAppSettings prepares the ListStaticSiteBuildFunctionAppSettings request.
func (c StaticSitesClient) preparerForListStaticSiteBuildFunctionAppSettings(ctx context.Context, id BuildId) (*http.Request, error) {
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

// responderForListStaticSiteBuildFunctionAppSettings handles the response to the ListStaticSiteBuildFunctionAppSettings request. The method always
// closes the http.Response Body.
func (c StaticSitesClient) responderForListStaticSiteBuildFunctionAppSettings(resp *http.Response) (result ListStaticSiteBuildFunctionAppSettingsOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

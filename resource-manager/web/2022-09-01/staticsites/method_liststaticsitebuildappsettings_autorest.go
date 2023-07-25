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

type ListStaticSiteBuildAppSettingsOperationResponse struct {
	HttpResponse *http.Response
	Model        *StringDictionary
}

// ListStaticSiteBuildAppSettings ...
func (c StaticSitesClient) ListStaticSiteBuildAppSettings(ctx context.Context, id BuildId) (result ListStaticSiteBuildAppSettingsOperationResponse, err error) {
	req, err := c.preparerForListStaticSiteBuildAppSettings(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "ListStaticSiteBuildAppSettings", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "ListStaticSiteBuildAppSettings", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForListStaticSiteBuildAppSettings(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "ListStaticSiteBuildAppSettings", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForListStaticSiteBuildAppSettings prepares the ListStaticSiteBuildAppSettings request.
func (c StaticSitesClient) preparerForListStaticSiteBuildAppSettings(ctx context.Context, id BuildId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/listAppSettings", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForListStaticSiteBuildAppSettings handles the response to the ListStaticSiteBuildAppSettings request. The method always
// closes the http.Response Body.
func (c StaticSitesClient) responderForListStaticSiteBuildAppSettings(resp *http.Response) (result ListStaticSiteBuildAppSettingsOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

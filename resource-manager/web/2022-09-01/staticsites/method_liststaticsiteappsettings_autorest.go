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

type ListStaticSiteAppSettingsOperationResponse struct {
	HttpResponse *http.Response
	Model        *StringDictionary
}

// ListStaticSiteAppSettings ...
func (c StaticSitesClient) ListStaticSiteAppSettings(ctx context.Context, id StaticSiteId) (result ListStaticSiteAppSettingsOperationResponse, err error) {
	req, err := c.preparerForListStaticSiteAppSettings(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "ListStaticSiteAppSettings", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "ListStaticSiteAppSettings", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForListStaticSiteAppSettings(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "ListStaticSiteAppSettings", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForListStaticSiteAppSettings prepares the ListStaticSiteAppSettings request.
func (c StaticSitesClient) preparerForListStaticSiteAppSettings(ctx context.Context, id StaticSiteId) (*http.Request, error) {
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

// responderForListStaticSiteAppSettings handles the response to the ListStaticSiteAppSettings request. The method always
// closes the http.Response Body.
func (c StaticSitesClient) responderForListStaticSiteAppSettings(resp *http.Response) (result ListStaticSiteAppSettingsOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

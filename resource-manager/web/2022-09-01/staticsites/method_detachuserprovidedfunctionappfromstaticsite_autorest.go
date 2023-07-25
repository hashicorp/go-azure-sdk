package staticsites

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DetachUserProvidedFunctionAppFromStaticSiteOperationResponse struct {
	HttpResponse *http.Response
}

// DetachUserProvidedFunctionAppFromStaticSite ...
func (c StaticSitesClient) DetachUserProvidedFunctionAppFromStaticSite(ctx context.Context, id UserProvidedFunctionAppId) (result DetachUserProvidedFunctionAppFromStaticSiteOperationResponse, err error) {
	req, err := c.preparerForDetachUserProvidedFunctionAppFromStaticSite(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "DetachUserProvidedFunctionAppFromStaticSite", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "DetachUserProvidedFunctionAppFromStaticSite", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForDetachUserProvidedFunctionAppFromStaticSite(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "DetachUserProvidedFunctionAppFromStaticSite", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForDetachUserProvidedFunctionAppFromStaticSite prepares the DetachUserProvidedFunctionAppFromStaticSite request.
func (c StaticSitesClient) preparerForDetachUserProvidedFunctionAppFromStaticSite(ctx context.Context, id UserProvidedFunctionAppId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsDelete(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForDetachUserProvidedFunctionAppFromStaticSite handles the response to the DetachUserProvidedFunctionAppFromStaticSite request. The method always
// closes the http.Response Body.
func (c StaticSitesClient) responderForDetachUserProvidedFunctionAppFromStaticSite(resp *http.Response) (result DetachUserProvidedFunctionAppFromStaticSiteOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusNoContent, http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

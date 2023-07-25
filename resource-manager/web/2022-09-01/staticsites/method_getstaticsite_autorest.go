package staticsites

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetStaticSiteOperationResponse struct {
	HttpResponse *http.Response
	Model        *StaticSiteARMResource
}

// GetStaticSite ...
func (c StaticSitesClient) GetStaticSite(ctx context.Context, id StaticSiteId) (result GetStaticSiteOperationResponse, err error) {
	req, err := c.preparerForGetStaticSite(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "GetStaticSite", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "GetStaticSite", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetStaticSite(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "GetStaticSite", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetStaticSite prepares the GetStaticSite request.
func (c StaticSitesClient) preparerForGetStaticSite(ctx context.Context, id StaticSiteId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForGetStaticSite handles the response to the GetStaticSite request. The method always
// closes the http.Response Body.
func (c StaticSitesClient) responderForGetStaticSite(resp *http.Response) (result GetStaticSiteOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

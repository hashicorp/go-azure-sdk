package staticsites

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetStaticSiteCustomDomainOperationResponse struct {
	HttpResponse *http.Response
	Model        *StaticSiteCustomDomainOverviewARMResource
}

// GetStaticSiteCustomDomain ...
func (c StaticSitesClient) GetStaticSiteCustomDomain(ctx context.Context, id CustomDomainId) (result GetStaticSiteCustomDomainOperationResponse, err error) {
	req, err := c.preparerForGetStaticSiteCustomDomain(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "GetStaticSiteCustomDomain", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "GetStaticSiteCustomDomain", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetStaticSiteCustomDomain(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "GetStaticSiteCustomDomain", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetStaticSiteCustomDomain prepares the GetStaticSiteCustomDomain request.
func (c StaticSitesClient) preparerForGetStaticSiteCustomDomain(ctx context.Context, id CustomDomainId) (*http.Request, error) {
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

// responderForGetStaticSiteCustomDomain handles the response to the GetStaticSiteCustomDomain request. The method always
// closes the http.Response Body.
func (c StaticSitesClient) responderForGetStaticSiteCustomDomain(resp *http.Response) (result GetStaticSiteCustomDomainOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

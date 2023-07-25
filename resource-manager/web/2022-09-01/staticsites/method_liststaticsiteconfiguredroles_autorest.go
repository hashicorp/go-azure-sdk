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

type ListStaticSiteConfiguredRolesOperationResponse struct {
	HttpResponse *http.Response
	Model        *StringList
}

// ListStaticSiteConfiguredRoles ...
func (c StaticSitesClient) ListStaticSiteConfiguredRoles(ctx context.Context, id StaticSiteId) (result ListStaticSiteConfiguredRolesOperationResponse, err error) {
	req, err := c.preparerForListStaticSiteConfiguredRoles(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "ListStaticSiteConfiguredRoles", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "ListStaticSiteConfiguredRoles", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForListStaticSiteConfiguredRoles(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "ListStaticSiteConfiguredRoles", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForListStaticSiteConfiguredRoles prepares the ListStaticSiteConfiguredRoles request.
func (c StaticSitesClient) preparerForListStaticSiteConfiguredRoles(ctx context.Context, id StaticSiteId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/listConfiguredRoles", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForListStaticSiteConfiguredRoles handles the response to the ListStaticSiteConfiguredRoles request. The method always
// closes the http.Response Body.
func (c StaticSitesClient) responderForListStaticSiteConfiguredRoles(resp *http.Response) (result ListStaticSiteConfiguredRolesOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

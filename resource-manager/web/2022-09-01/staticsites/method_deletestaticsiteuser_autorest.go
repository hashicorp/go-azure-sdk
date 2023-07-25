package staticsites

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeleteStaticSiteUserOperationResponse struct {
	HttpResponse *http.Response
}

// DeleteStaticSiteUser ...
func (c StaticSitesClient) DeleteStaticSiteUser(ctx context.Context, id UserId) (result DeleteStaticSiteUserOperationResponse, err error) {
	req, err := c.preparerForDeleteStaticSiteUser(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "DeleteStaticSiteUser", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "DeleteStaticSiteUser", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForDeleteStaticSiteUser(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "DeleteStaticSiteUser", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForDeleteStaticSiteUser prepares the DeleteStaticSiteUser request.
func (c StaticSitesClient) preparerForDeleteStaticSiteUser(ctx context.Context, id UserId) (*http.Request, error) {
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

// responderForDeleteStaticSiteUser handles the response to the DeleteStaticSiteUser request. The method always
// closes the http.Response Body.
func (c StaticSitesClient) responderForDeleteStaticSiteUser(resp *http.Response) (result DeleteStaticSiteUserOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

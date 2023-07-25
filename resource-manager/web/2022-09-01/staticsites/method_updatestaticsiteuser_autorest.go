package staticsites

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateStaticSiteUserOperationResponse struct {
	HttpResponse *http.Response
	Model        *StaticSiteUserARMResource
}

// UpdateStaticSiteUser ...
func (c StaticSitesClient) UpdateStaticSiteUser(ctx context.Context, id UserId, input StaticSiteUserARMResource) (result UpdateStaticSiteUserOperationResponse, err error) {
	req, err := c.preparerForUpdateStaticSiteUser(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "UpdateStaticSiteUser", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "UpdateStaticSiteUser", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForUpdateStaticSiteUser(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "UpdateStaticSiteUser", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForUpdateStaticSiteUser prepares the UpdateStaticSiteUser request.
func (c StaticSitesClient) preparerForUpdateStaticSiteUser(ctx context.Context, id UserId, input StaticSiteUserARMResource) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForUpdateStaticSiteUser handles the response to the UpdateStaticSiteUser request. The method always
// closes the http.Response Body.
func (c StaticSitesClient) responderForUpdateStaticSiteUser(resp *http.Response) (result UpdateStaticSiteUserOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

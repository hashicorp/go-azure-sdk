package staticsites

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateStaticSiteOperationResponse struct {
	HttpResponse *http.Response
	Model        *StaticSiteARMResource
}

// UpdateStaticSite ...
func (c StaticSitesClient) UpdateStaticSite(ctx context.Context, id StaticSiteId, input StaticSitePatchResource) (result UpdateStaticSiteOperationResponse, err error) {
	req, err := c.preparerForUpdateStaticSite(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "UpdateStaticSite", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "UpdateStaticSite", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForUpdateStaticSite(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "UpdateStaticSite", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForUpdateStaticSite prepares the UpdateStaticSite request.
func (c StaticSitesClient) preparerForUpdateStaticSite(ctx context.Context, id StaticSiteId, input StaticSitePatchResource) (*http.Request, error) {
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

// responderForUpdateStaticSite handles the response to the UpdateStaticSite request. The method always
// closes the http.Response Body.
func (c StaticSitesClient) responderForUpdateStaticSite(resp *http.Response) (result UpdateStaticSiteOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusAccepted, http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

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

type ResetStaticSiteApiKeyOperationResponse struct {
	HttpResponse *http.Response
}

// ResetStaticSiteApiKey ...
func (c StaticSitesClient) ResetStaticSiteApiKey(ctx context.Context, id StaticSiteId, input StaticSiteResetPropertiesARMResource) (result ResetStaticSiteApiKeyOperationResponse, err error) {
	req, err := c.preparerForResetStaticSiteApiKey(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "ResetStaticSiteApiKey", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "ResetStaticSiteApiKey", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForResetStaticSiteApiKey(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "ResetStaticSiteApiKey", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForResetStaticSiteApiKey prepares the ResetStaticSiteApiKey request.
func (c StaticSitesClient) preparerForResetStaticSiteApiKey(ctx context.Context, id StaticSiteId, input StaticSiteResetPropertiesARMResource) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/resetapikey", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForResetStaticSiteApiKey handles the response to the ResetStaticSiteApiKey request. The method always
// closes the http.Response Body.
func (c StaticSitesClient) responderForResetStaticSiteApiKey(resp *http.Response) (result ResetStaticSiteApiKeyOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

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

type ListStaticSiteSecretsOperationResponse struct {
	HttpResponse *http.Response
	Model        *StringDictionary
}

// ListStaticSiteSecrets ...
func (c StaticSitesClient) ListStaticSiteSecrets(ctx context.Context, id StaticSiteId) (result ListStaticSiteSecretsOperationResponse, err error) {
	req, err := c.preparerForListStaticSiteSecrets(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "ListStaticSiteSecrets", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "ListStaticSiteSecrets", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForListStaticSiteSecrets(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "ListStaticSiteSecrets", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForListStaticSiteSecrets prepares the ListStaticSiteSecrets request.
func (c StaticSitesClient) preparerForListStaticSiteSecrets(ctx context.Context, id StaticSiteId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/listSecrets", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForListStaticSiteSecrets handles the response to the ListStaticSiteSecrets request. The method always
// closes the http.Response Body.
func (c StaticSitesClient) responderForListStaticSiteSecrets(resp *http.Response) (result ListStaticSiteSecretsOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

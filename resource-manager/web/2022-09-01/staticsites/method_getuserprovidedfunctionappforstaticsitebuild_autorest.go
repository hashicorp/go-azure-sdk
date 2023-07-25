package staticsites

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetUserProvidedFunctionAppForStaticSiteBuildOperationResponse struct {
	HttpResponse *http.Response
	Model        *StaticSiteUserProvidedFunctionAppARMResource
}

// GetUserProvidedFunctionAppForStaticSiteBuild ...
func (c StaticSitesClient) GetUserProvidedFunctionAppForStaticSiteBuild(ctx context.Context, id BuildUserProvidedFunctionAppId) (result GetUserProvidedFunctionAppForStaticSiteBuildOperationResponse, err error) {
	req, err := c.preparerForGetUserProvidedFunctionAppForStaticSiteBuild(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "GetUserProvidedFunctionAppForStaticSiteBuild", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "GetUserProvidedFunctionAppForStaticSiteBuild", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetUserProvidedFunctionAppForStaticSiteBuild(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "GetUserProvidedFunctionAppForStaticSiteBuild", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetUserProvidedFunctionAppForStaticSiteBuild prepares the GetUserProvidedFunctionAppForStaticSiteBuild request.
func (c StaticSitesClient) preparerForGetUserProvidedFunctionAppForStaticSiteBuild(ctx context.Context, id BuildUserProvidedFunctionAppId) (*http.Request, error) {
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

// responderForGetUserProvidedFunctionAppForStaticSiteBuild handles the response to the GetUserProvidedFunctionAppForStaticSiteBuild request. The method always
// closes the http.Response Body.
func (c StaticSitesClient) responderForGetUserProvidedFunctionAppForStaticSiteBuild(resp *http.Response) (result GetUserProvidedFunctionAppForStaticSiteBuildOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

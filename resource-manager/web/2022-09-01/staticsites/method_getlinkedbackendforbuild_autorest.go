package staticsites

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetLinkedBackendForBuildOperationResponse struct {
	HttpResponse *http.Response
	Model        *StaticSiteLinkedBackendARMResource
}

// GetLinkedBackendForBuild ...
func (c StaticSitesClient) GetLinkedBackendForBuild(ctx context.Context, id BuildLinkedBackendId) (result GetLinkedBackendForBuildOperationResponse, err error) {
	req, err := c.preparerForGetLinkedBackendForBuild(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "GetLinkedBackendForBuild", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "GetLinkedBackendForBuild", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetLinkedBackendForBuild(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "GetLinkedBackendForBuild", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetLinkedBackendForBuild prepares the GetLinkedBackendForBuild request.
func (c StaticSitesClient) preparerForGetLinkedBackendForBuild(ctx context.Context, id BuildLinkedBackendId) (*http.Request, error) {
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

// responderForGetLinkedBackendForBuild handles the response to the GetLinkedBackendForBuild request. The method always
// closes the http.Response Body.
func (c StaticSitesClient) responderForGetLinkedBackendForBuild(resp *http.Response) (result GetLinkedBackendForBuildOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

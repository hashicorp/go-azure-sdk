package staticsites

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetStaticSiteBuildOperationResponse struct {
	HttpResponse *http.Response
	Model        *StaticSiteBuildARMResource
}

// GetStaticSiteBuild ...
func (c StaticSitesClient) GetStaticSiteBuild(ctx context.Context, id BuildId) (result GetStaticSiteBuildOperationResponse, err error) {
	req, err := c.preparerForGetStaticSiteBuild(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "GetStaticSiteBuild", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "GetStaticSiteBuild", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetStaticSiteBuild(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "GetStaticSiteBuild", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetStaticSiteBuild prepares the GetStaticSiteBuild request.
func (c StaticSitesClient) preparerForGetStaticSiteBuild(ctx context.Context, id BuildId) (*http.Request, error) {
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

// responderForGetStaticSiteBuild handles the response to the GetStaticSiteBuild request. The method always
// closes the http.Response Body.
func (c StaticSitesClient) responderForGetStaticSiteBuild(resp *http.Response) (result GetStaticSiteBuildOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

package staticsites

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetLinkedBackendOperationResponse struct {
	HttpResponse *http.Response
	Model        *StaticSiteLinkedBackendARMResource
}

// GetLinkedBackend ...
func (c StaticSitesClient) GetLinkedBackend(ctx context.Context, id LinkedBackendId) (result GetLinkedBackendOperationResponse, err error) {
	req, err := c.preparerForGetLinkedBackend(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "GetLinkedBackend", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "GetLinkedBackend", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetLinkedBackend(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "GetLinkedBackend", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetLinkedBackend prepares the GetLinkedBackend request.
func (c StaticSitesClient) preparerForGetLinkedBackend(ctx context.Context, id LinkedBackendId) (*http.Request, error) {
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

// responderForGetLinkedBackend handles the response to the GetLinkedBackend request. The method always
// closes the http.Response Body.
func (c StaticSitesClient) responderForGetLinkedBackend(resp *http.Response) (result GetLinkedBackendOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

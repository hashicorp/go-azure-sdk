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

type GetBasicAuthOperationResponse struct {
	HttpResponse *http.Response
	Model        *StaticSiteBasicAuthPropertiesARMResource
}

// GetBasicAuth ...
func (c StaticSitesClient) GetBasicAuth(ctx context.Context, id StaticSiteId) (result GetBasicAuthOperationResponse, err error) {
	req, err := c.preparerForGetBasicAuth(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "GetBasicAuth", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "GetBasicAuth", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetBasicAuth(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "GetBasicAuth", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetBasicAuth prepares the GetBasicAuth request.
func (c StaticSitesClient) preparerForGetBasicAuth(ctx context.Context, id StaticSiteId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/basicAuth/default", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForGetBasicAuth handles the response to the GetBasicAuth request. The method always
// closes the http.Response Body.
func (c StaticSitesClient) responderForGetBasicAuth(resp *http.Response) (result GetBasicAuthOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

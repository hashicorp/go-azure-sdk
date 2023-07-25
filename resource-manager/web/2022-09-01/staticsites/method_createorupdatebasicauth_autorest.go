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

type CreateOrUpdateBasicAuthOperationResponse struct {
	HttpResponse *http.Response
	Model        *StaticSiteBasicAuthPropertiesARMResource
}

// CreateOrUpdateBasicAuth ...
func (c StaticSitesClient) CreateOrUpdateBasicAuth(ctx context.Context, id StaticSiteId, input StaticSiteBasicAuthPropertiesARMResource) (result CreateOrUpdateBasicAuthOperationResponse, err error) {
	req, err := c.preparerForCreateOrUpdateBasicAuth(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "CreateOrUpdateBasicAuth", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "CreateOrUpdateBasicAuth", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForCreateOrUpdateBasicAuth(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "CreateOrUpdateBasicAuth", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForCreateOrUpdateBasicAuth prepares the CreateOrUpdateBasicAuth request.
func (c StaticSitesClient) preparerForCreateOrUpdateBasicAuth(ctx context.Context, id StaticSiteId, input StaticSiteBasicAuthPropertiesARMResource) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/basicAuth/default", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForCreateOrUpdateBasicAuth handles the response to the CreateOrUpdateBasicAuth request. The method always
// closes the http.Response Body.
func (c StaticSitesClient) responderForCreateOrUpdateBasicAuth(resp *http.Response) (result CreateOrUpdateBasicAuthOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

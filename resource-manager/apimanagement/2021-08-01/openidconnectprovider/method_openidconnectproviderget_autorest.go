package openidconnectprovider

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OpenIdConnectProviderGetOperationResponse struct {
	HttpResponse *http.Response
	Model        *OpenidConnectProviderContract
}

// OpenIdConnectProviderGet ...
func (c OpenidConnectProviderClient) OpenIdConnectProviderGet(ctx context.Context, id OpenidConnectProviderId) (result OpenIdConnectProviderGetOperationResponse, err error) {
	req, err := c.preparerForOpenIdConnectProviderGet(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "openidconnectprovider.OpenidConnectProviderClient", "OpenIdConnectProviderGet", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "openidconnectprovider.OpenidConnectProviderClient", "OpenIdConnectProviderGet", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForOpenIdConnectProviderGet(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "openidconnectprovider.OpenidConnectProviderClient", "OpenIdConnectProviderGet", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForOpenIdConnectProviderGet prepares the OpenIdConnectProviderGet request.
func (c OpenidConnectProviderClient) preparerForOpenIdConnectProviderGet(ctx context.Context, id OpenidConnectProviderId) (*http.Request, error) {
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

// responderForOpenIdConnectProviderGet handles the response to the OpenIdConnectProviderGet request. The method always
// closes the http.Response Body.
func (c OpenidConnectProviderClient) responderForOpenIdConnectProviderGet(resp *http.Response) (result OpenIdConnectProviderGetOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

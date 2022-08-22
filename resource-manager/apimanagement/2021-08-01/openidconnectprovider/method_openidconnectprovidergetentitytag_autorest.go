package openidconnectprovider

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OpenIdConnectProviderGetEntityTagOperationResponse struct {
	HttpResponse *http.Response
}

// OpenIdConnectProviderGetEntityTag ...
func (c OpenidConnectProviderClient) OpenIdConnectProviderGetEntityTag(ctx context.Context, id OpenidConnectProviderId) (result OpenIdConnectProviderGetEntityTagOperationResponse, err error) {
	req, err := c.preparerForOpenIdConnectProviderGetEntityTag(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "openidconnectprovider.OpenidConnectProviderClient", "OpenIdConnectProviderGetEntityTag", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "openidconnectprovider.OpenidConnectProviderClient", "OpenIdConnectProviderGetEntityTag", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForOpenIdConnectProviderGetEntityTag(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "openidconnectprovider.OpenidConnectProviderClient", "OpenIdConnectProviderGetEntityTag", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForOpenIdConnectProviderGetEntityTag prepares the OpenIdConnectProviderGetEntityTag request.
func (c OpenidConnectProviderClient) preparerForOpenIdConnectProviderGetEntityTag(ctx context.Context, id OpenidConnectProviderId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsHead(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForOpenIdConnectProviderGetEntityTag handles the response to the OpenIdConnectProviderGetEntityTag request. The method always
// closes the http.Response Body.
func (c OpenidConnectProviderClient) responderForOpenIdConnectProviderGetEntityTag(resp *http.Response) (result OpenIdConnectProviderGetEntityTagOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

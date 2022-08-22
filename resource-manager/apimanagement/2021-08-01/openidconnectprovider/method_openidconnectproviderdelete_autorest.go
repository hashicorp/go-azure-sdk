package openidconnectprovider

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OpenIdConnectProviderDeleteOperationResponse struct {
	HttpResponse *http.Response
}

type OpenIdConnectProviderDeleteOperationOptions struct {
	IfMatch *string
}

func DefaultOpenIdConnectProviderDeleteOperationOptions() OpenIdConnectProviderDeleteOperationOptions {
	return OpenIdConnectProviderDeleteOperationOptions{}
}

func (o OpenIdConnectProviderDeleteOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	if o.IfMatch != nil {
		out["If-Match"] = *o.IfMatch
	}

	return out
}

func (o OpenIdConnectProviderDeleteOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

// OpenIdConnectProviderDelete ...
func (c OpenidConnectProviderClient) OpenIdConnectProviderDelete(ctx context.Context, id OpenidConnectProviderId, options OpenIdConnectProviderDeleteOperationOptions) (result OpenIdConnectProviderDeleteOperationResponse, err error) {
	req, err := c.preparerForOpenIdConnectProviderDelete(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "openidconnectprovider.OpenidConnectProviderClient", "OpenIdConnectProviderDelete", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "openidconnectprovider.OpenidConnectProviderClient", "OpenIdConnectProviderDelete", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForOpenIdConnectProviderDelete(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "openidconnectprovider.OpenidConnectProviderClient", "OpenIdConnectProviderDelete", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForOpenIdConnectProviderDelete prepares the OpenIdConnectProviderDelete request.
func (c OpenidConnectProviderClient) preparerForOpenIdConnectProviderDelete(ctx context.Context, id OpenidConnectProviderId, options OpenIdConnectProviderDeleteOperationOptions) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	for k, v := range options.toQueryString() {
		queryParameters[k] = autorest.Encode("query", v)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsDelete(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithHeaders(options.toHeaders()),
		autorest.WithPath(id.ID()),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForOpenIdConnectProviderDelete handles the response to the OpenIdConnectProviderDelete request. The method always
// closes the http.Response Body.
func (c OpenidConnectProviderClient) responderForOpenIdConnectProviderDelete(resp *http.Response) (result OpenIdConnectProviderDeleteOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusNoContent, http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

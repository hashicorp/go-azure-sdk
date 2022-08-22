package openidconnectprovider

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OpenIdConnectProviderCreateOrUpdateOperationResponse struct {
	HttpResponse *http.Response
	Model        *OpenidConnectProviderContract
}

type OpenIdConnectProviderCreateOrUpdateOperationOptions struct {
	IfMatch *string
}

func DefaultOpenIdConnectProviderCreateOrUpdateOperationOptions() OpenIdConnectProviderCreateOrUpdateOperationOptions {
	return OpenIdConnectProviderCreateOrUpdateOperationOptions{}
}

func (o OpenIdConnectProviderCreateOrUpdateOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	if o.IfMatch != nil {
		out["If-Match"] = *o.IfMatch
	}

	return out
}

func (o OpenIdConnectProviderCreateOrUpdateOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

// OpenIdConnectProviderCreateOrUpdate ...
func (c OpenidConnectProviderClient) OpenIdConnectProviderCreateOrUpdate(ctx context.Context, id OpenidConnectProviderId, input OpenidConnectProviderContract, options OpenIdConnectProviderCreateOrUpdateOperationOptions) (result OpenIdConnectProviderCreateOrUpdateOperationResponse, err error) {
	req, err := c.preparerForOpenIdConnectProviderCreateOrUpdate(ctx, id, input, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "openidconnectprovider.OpenidConnectProviderClient", "OpenIdConnectProviderCreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "openidconnectprovider.OpenidConnectProviderClient", "OpenIdConnectProviderCreateOrUpdate", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForOpenIdConnectProviderCreateOrUpdate(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "openidconnectprovider.OpenidConnectProviderClient", "OpenIdConnectProviderCreateOrUpdate", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForOpenIdConnectProviderCreateOrUpdate prepares the OpenIdConnectProviderCreateOrUpdate request.
func (c OpenidConnectProviderClient) preparerForOpenIdConnectProviderCreateOrUpdate(ctx context.Context, id OpenidConnectProviderId, input OpenidConnectProviderContract, options OpenIdConnectProviderCreateOrUpdateOperationOptions) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	for k, v := range options.toQueryString() {
		queryParameters[k] = autorest.Encode("query", v)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithHeaders(options.toHeaders()),
		autorest.WithPath(id.ID()),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForOpenIdConnectProviderCreateOrUpdate handles the response to the OpenIdConnectProviderCreateOrUpdate request. The method always
// closes the http.Response Body.
func (c OpenidConnectProviderClient) responderForOpenIdConnectProviderCreateOrUpdate(resp *http.Response) (result OpenIdConnectProviderCreateOrUpdateOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusCreated, http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

package openidconnectprovider

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OpenIdConnectProviderUpdateOperationResponse struct {
	HttpResponse *http.Response
	Model        *OpenidConnectProviderContract
}

type OpenIdConnectProviderUpdateOperationOptions struct {
	IfMatch *string
}

func DefaultOpenIdConnectProviderUpdateOperationOptions() OpenIdConnectProviderUpdateOperationOptions {
	return OpenIdConnectProviderUpdateOperationOptions{}
}

func (o OpenIdConnectProviderUpdateOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	if o.IfMatch != nil {
		out["If-Match"] = *o.IfMatch
	}

	return out
}

func (o OpenIdConnectProviderUpdateOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

// OpenIdConnectProviderUpdate ...
func (c OpenidConnectProviderClient) OpenIdConnectProviderUpdate(ctx context.Context, id OpenidConnectProviderId, input OpenidConnectProviderUpdateContract, options OpenIdConnectProviderUpdateOperationOptions) (result OpenIdConnectProviderUpdateOperationResponse, err error) {
	req, err := c.preparerForOpenIdConnectProviderUpdate(ctx, id, input, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "openidconnectprovider.OpenidConnectProviderClient", "OpenIdConnectProviderUpdate", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "openidconnectprovider.OpenidConnectProviderClient", "OpenIdConnectProviderUpdate", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForOpenIdConnectProviderUpdate(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "openidconnectprovider.OpenidConnectProviderClient", "OpenIdConnectProviderUpdate", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForOpenIdConnectProviderUpdate prepares the OpenIdConnectProviderUpdate request.
func (c OpenidConnectProviderClient) preparerForOpenIdConnectProviderUpdate(ctx context.Context, id OpenidConnectProviderId, input OpenidConnectProviderUpdateContract, options OpenIdConnectProviderUpdateOperationOptions) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	for k, v := range options.toQueryString() {
		queryParameters[k] = autorest.Encode("query", v)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithHeaders(options.toHeaders()),
		autorest.WithPath(id.ID()),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForOpenIdConnectProviderUpdate handles the response to the OpenIdConnectProviderUpdate request. The method always
// closes the http.Response Body.
func (c OpenidConnectProviderClient) responderForOpenIdConnectProviderUpdate(resp *http.Response) (result OpenIdConnectProviderUpdateOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

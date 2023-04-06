package dppresourceguardproxies

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DppResourceGuardProxyCreateOrUpdateOperationResponse struct {
	HttpResponse *http.Response
	Model        *ResourceGuardProxyBaseResource
}

// DppResourceGuardProxyCreateOrUpdate ...
func (c DppResourceGuardProxiesClient) DppResourceGuardProxyCreateOrUpdate(ctx context.Context, id BackupResourceGuardProxyId, input ResourceGuardProxyBaseResource) (result DppResourceGuardProxyCreateOrUpdateOperationResponse, err error) {
	req, err := c.preparerForDppResourceGuardProxyCreateOrUpdate(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dppresourceguardproxies.DppResourceGuardProxiesClient", "DppResourceGuardProxyCreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "dppresourceguardproxies.DppResourceGuardProxiesClient", "DppResourceGuardProxyCreateOrUpdate", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForDppResourceGuardProxyCreateOrUpdate(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dppresourceguardproxies.DppResourceGuardProxiesClient", "DppResourceGuardProxyCreateOrUpdate", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForDppResourceGuardProxyCreateOrUpdate prepares the DppResourceGuardProxyCreateOrUpdate request.
func (c DppResourceGuardProxiesClient) preparerForDppResourceGuardProxyCreateOrUpdate(ctx context.Context, id BackupResourceGuardProxyId, input ResourceGuardProxyBaseResource) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForDppResourceGuardProxyCreateOrUpdate handles the response to the DppResourceGuardProxyCreateOrUpdate request. The method always
// closes the http.Response Body.
func (c DppResourceGuardProxiesClient) responderForDppResourceGuardProxyCreateOrUpdate(resp *http.Response) (result DppResourceGuardProxyCreateOrUpdateOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

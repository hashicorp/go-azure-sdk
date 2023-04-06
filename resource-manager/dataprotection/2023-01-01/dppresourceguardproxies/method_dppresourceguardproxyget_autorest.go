package dppresourceguardproxies

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DppResourceGuardProxyGetOperationResponse struct {
	HttpResponse *http.Response
	Model        *ResourceGuardProxyBaseResource
}

// DppResourceGuardProxyGet ...
func (c DppResourceGuardProxiesClient) DppResourceGuardProxyGet(ctx context.Context, id BackupResourceGuardProxyId) (result DppResourceGuardProxyGetOperationResponse, err error) {
	req, err := c.preparerForDppResourceGuardProxyGet(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dppresourceguardproxies.DppResourceGuardProxiesClient", "DppResourceGuardProxyGet", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "dppresourceguardproxies.DppResourceGuardProxiesClient", "DppResourceGuardProxyGet", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForDppResourceGuardProxyGet(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dppresourceguardproxies.DppResourceGuardProxiesClient", "DppResourceGuardProxyGet", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForDppResourceGuardProxyGet prepares the DppResourceGuardProxyGet request.
func (c DppResourceGuardProxiesClient) preparerForDppResourceGuardProxyGet(ctx context.Context, id BackupResourceGuardProxyId) (*http.Request, error) {
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

// responderForDppResourceGuardProxyGet handles the response to the DppResourceGuardProxyGet request. The method always
// closes the http.Response Body.
func (c DppResourceGuardProxiesClient) responderForDppResourceGuardProxyGet(resp *http.Response) (result DppResourceGuardProxyGetOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

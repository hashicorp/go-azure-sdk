package dppresourceguardproxies

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DppResourceGuardProxyDeleteOperationResponse struct {
	HttpResponse *http.Response
}

// DppResourceGuardProxyDelete ...
func (c DppResourceGuardProxiesClient) DppResourceGuardProxyDelete(ctx context.Context, id BackupResourceGuardProxyId) (result DppResourceGuardProxyDeleteOperationResponse, err error) {
	req, err := c.preparerForDppResourceGuardProxyDelete(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dppresourceguardproxies.DppResourceGuardProxiesClient", "DppResourceGuardProxyDelete", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "dppresourceguardproxies.DppResourceGuardProxiesClient", "DppResourceGuardProxyDelete", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForDppResourceGuardProxyDelete(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dppresourceguardproxies.DppResourceGuardProxiesClient", "DppResourceGuardProxyDelete", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForDppResourceGuardProxyDelete prepares the DppResourceGuardProxyDelete request.
func (c DppResourceGuardProxiesClient) preparerForDppResourceGuardProxyDelete(ctx context.Context, id BackupResourceGuardProxyId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsDelete(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForDppResourceGuardProxyDelete handles the response to the DppResourceGuardProxyDelete request. The method always
// closes the http.Response Body.
func (c DppResourceGuardProxiesClient) responderForDppResourceGuardProxyDelete(resp *http.Response) (result DppResourceGuardProxyDeleteOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusNoContent, http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

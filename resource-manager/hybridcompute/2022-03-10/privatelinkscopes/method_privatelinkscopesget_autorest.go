package privatelinkscopes

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivateLinkScopesGetOperationResponse struct {
	HttpResponse *http.Response
	Model        *HybridComputePrivateLinkScope
}

// PrivateLinkScopesGet ...
func (c PrivateLinkScopesClient) PrivateLinkScopesGet(ctx context.Context, id ProviderPrivateLinkScopeId) (result PrivateLinkScopesGetOperationResponse, err error) {
	req, err := c.preparerForPrivateLinkScopesGet(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "privatelinkscopes.PrivateLinkScopesClient", "PrivateLinkScopesGet", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "privatelinkscopes.PrivateLinkScopesClient", "PrivateLinkScopesGet", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForPrivateLinkScopesGet(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "privatelinkscopes.PrivateLinkScopesClient", "PrivateLinkScopesGet", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForPrivateLinkScopesGet prepares the PrivateLinkScopesGet request.
func (c PrivateLinkScopesClient) preparerForPrivateLinkScopesGet(ctx context.Context, id ProviderPrivateLinkScopeId) (*http.Request, error) {
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

// responderForPrivateLinkScopesGet handles the response to the PrivateLinkScopesGet request. The method always
// closes the http.Response Body.
func (c PrivateLinkScopesClient) responderForPrivateLinkScopesGet(resp *http.Response) (result PrivateLinkScopesGetOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

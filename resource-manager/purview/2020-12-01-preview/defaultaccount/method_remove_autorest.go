package defaultaccount

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RemoveOperationResponse struct {
	HttpResponse *http.Response
}

type RemoveOperationOptions struct {
	Scope         *string
	ScopeTenantId *string
	ScopeType     *ScopeType
}

func DefaultRemoveOperationOptions() RemoveOperationOptions {
	return RemoveOperationOptions{}
}

func (o RemoveOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o RemoveOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.Scope != nil {
		out["scope"] = *o.Scope
	}

	if o.ScopeTenantId != nil {
		out["scopeTenantId"] = *o.ScopeTenantId
	}

	if o.ScopeType != nil {
		out["scopeType"] = *o.ScopeType
	}

	return out
}

// Remove ...
func (c DefaultAccountClient) Remove(ctx context.Context, options RemoveOperationOptions) (result RemoveOperationResponse, err error) {
	req, err := c.preparerForRemove(ctx, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "defaultaccount.DefaultAccountClient", "Remove", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "defaultaccount.DefaultAccountClient", "Remove", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForRemove(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "defaultaccount.DefaultAccountClient", "Remove", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForRemove prepares the Remove request.
func (c DefaultAccountClient) preparerForRemove(ctx context.Context, options RemoveOperationOptions) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	for k, v := range options.toQueryString() {
		queryParameters[k] = autorest.Encode("query", v)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithHeaders(options.toHeaders()),
		autorest.WithPath("/providers/Microsoft.Purview/removeDefaultAccount"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForRemove handles the response to the Remove request. The method always
// closes the http.Response Body.
func (c DefaultAccountClient) responderForRemove(resp *http.Response) (result RemoveOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusNoContent, http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp
	return
}

package bms

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityPINsGetOperationResponse struct {
	HttpResponse *http.Response
	Model        *TokenInformation
}

type SecurityPINsGetOperationOptions struct {
	XMsAuthorizationAuxiliary *string
}

func DefaultSecurityPINsGetOperationOptions() SecurityPINsGetOperationOptions {
	return SecurityPINsGetOperationOptions{}
}

func (o SecurityPINsGetOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	if o.XMsAuthorizationAuxiliary != nil {
		out["x-ms-authorization-auxiliary"] = *o.XMsAuthorizationAuxiliary
	}

	return out
}

func (o SecurityPINsGetOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

// SecurityPINsGet ...
func (c BmsClient) SecurityPINsGet(ctx context.Context, id VaultId, input SecurityPinBase, options SecurityPINsGetOperationOptions) (result SecurityPINsGetOperationResponse, err error) {
	req, err := c.preparerForSecurityPINsGet(ctx, id, input, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "bms.BmsClient", "SecurityPINsGet", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "bms.BmsClient", "SecurityPINsGet", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForSecurityPINsGet(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "bms.BmsClient", "SecurityPINsGet", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForSecurityPINsGet prepares the SecurityPINsGet request.
func (c BmsClient) preparerForSecurityPINsGet(ctx context.Context, id VaultId, input SecurityPinBase, options SecurityPINsGetOperationOptions) (*http.Request, error) {
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
		autorest.WithPath(fmt.Sprintf("%s/backupSecurityPIN", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForSecurityPINsGet handles the response to the SecurityPINsGet request. The method always
// closes the http.Response Body.
func (c BmsClient) responderForSecurityPINsGet(resp *http.Response) (result SecurityPINsGetOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

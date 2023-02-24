package credentials

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CredentialOperationsGetOperationResponse struct {
	HttpResponse *http.Response
	Model        *ManagedIdentityCredentialResource
}

type CredentialOperationsGetOperationOptions struct {
	IfNoneMatch *string
}

func DefaultCredentialOperationsGetOperationOptions() CredentialOperationsGetOperationOptions {
	return CredentialOperationsGetOperationOptions{}
}

func (o CredentialOperationsGetOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	if o.IfNoneMatch != nil {
		out["If-None-Match"] = *o.IfNoneMatch
	}

	return out
}

func (o CredentialOperationsGetOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

// CredentialOperationsGet ...
func (c CredentialsClient) CredentialOperationsGet(ctx context.Context, id CredentialId, options CredentialOperationsGetOperationOptions) (result CredentialOperationsGetOperationResponse, err error) {
	req, err := c.preparerForCredentialOperationsGet(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "credentials.CredentialsClient", "CredentialOperationsGet", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "credentials.CredentialsClient", "CredentialOperationsGet", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForCredentialOperationsGet(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "credentials.CredentialsClient", "CredentialOperationsGet", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForCredentialOperationsGet prepares the CredentialOperationsGet request.
func (c CredentialsClient) preparerForCredentialOperationsGet(ctx context.Context, id CredentialId, options CredentialOperationsGetOperationOptions) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	for k, v := range options.toQueryString() {
		queryParameters[k] = autorest.Encode("query", v)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithHeaders(options.toHeaders()),
		autorest.WithPath(id.ID()),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForCredentialOperationsGet handles the response to the CredentialOperationsGet request. The method always
// closes the http.Response Body.
func (c CredentialsClient) responderForCredentialOperationsGet(resp *http.Response) (result CredentialOperationsGetOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

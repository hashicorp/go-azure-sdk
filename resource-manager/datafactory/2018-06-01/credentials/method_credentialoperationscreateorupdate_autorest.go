package credentials

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CredentialOperationsCreateOrUpdateOperationResponse struct {
	HttpResponse *http.Response
	Model        *ManagedIdentityCredentialResource
}

type CredentialOperationsCreateOrUpdateOperationOptions struct {
	IfMatch *string
}

func DefaultCredentialOperationsCreateOrUpdateOperationOptions() CredentialOperationsCreateOrUpdateOperationOptions {
	return CredentialOperationsCreateOrUpdateOperationOptions{}
}

func (o CredentialOperationsCreateOrUpdateOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	if o.IfMatch != nil {
		out["If-Match"] = *o.IfMatch
	}

	return out
}

func (o CredentialOperationsCreateOrUpdateOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

// CredentialOperationsCreateOrUpdate ...
func (c CredentialsClient) CredentialOperationsCreateOrUpdate(ctx context.Context, id CredentialId, input ManagedIdentityCredentialResource, options CredentialOperationsCreateOrUpdateOperationOptions) (result CredentialOperationsCreateOrUpdateOperationResponse, err error) {
	req, err := c.preparerForCredentialOperationsCreateOrUpdate(ctx, id, input, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "credentials.CredentialsClient", "CredentialOperationsCreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "credentials.CredentialsClient", "CredentialOperationsCreateOrUpdate", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForCredentialOperationsCreateOrUpdate(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "credentials.CredentialsClient", "CredentialOperationsCreateOrUpdate", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForCredentialOperationsCreateOrUpdate prepares the CredentialOperationsCreateOrUpdate request.
func (c CredentialsClient) preparerForCredentialOperationsCreateOrUpdate(ctx context.Context, id CredentialId, input ManagedIdentityCredentialResource, options CredentialOperationsCreateOrUpdateOperationOptions) (*http.Request, error) {
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

// responderForCredentialOperationsCreateOrUpdate handles the response to the CredentialOperationsCreateOrUpdate request. The method always
// closes the http.Response Body.
func (c CredentialsClient) responderForCredentialOperationsCreateOrUpdate(resp *http.Response) (result CredentialOperationsCreateOrUpdateOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

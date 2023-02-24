package credentials

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CredentialOperationsDeleteOperationResponse struct {
	HttpResponse *http.Response
}

// CredentialOperationsDelete ...
func (c CredentialsClient) CredentialOperationsDelete(ctx context.Context, id CredentialId) (result CredentialOperationsDeleteOperationResponse, err error) {
	req, err := c.preparerForCredentialOperationsDelete(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "credentials.CredentialsClient", "CredentialOperationsDelete", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "credentials.CredentialsClient", "CredentialOperationsDelete", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForCredentialOperationsDelete(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "credentials.CredentialsClient", "CredentialOperationsDelete", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForCredentialOperationsDelete prepares the CredentialOperationsDelete request.
func (c CredentialsClient) preparerForCredentialOperationsDelete(ctx context.Context, id CredentialId) (*http.Request, error) {
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

// responderForCredentialOperationsDelete handles the response to the CredentialOperationsDelete request. The method always
// closes the http.Response Body.
func (c CredentialsClient) responderForCredentialOperationsDelete(resp *http.Response) (result CredentialOperationsDeleteOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusNoContent, http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

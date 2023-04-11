package securitycontacts

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityContactsDeleteOperationResponse struct {
	HttpResponse *http.Response
}

// SecurityContactsDelete ...
func (c SecurityContactsClient) SecurityContactsDelete(ctx context.Context, id SecurityContactId) (result SecurityContactsDeleteOperationResponse, err error) {
	req, err := c.preparerForSecurityContactsDelete(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securitycontacts.SecurityContactsClient", "SecurityContactsDelete", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "securitycontacts.SecurityContactsClient", "SecurityContactsDelete", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForSecurityContactsDelete(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securitycontacts.SecurityContactsClient", "SecurityContactsDelete", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForSecurityContactsDelete prepares the SecurityContactsDelete request.
func (c SecurityContactsClient) preparerForSecurityContactsDelete(ctx context.Context, id SecurityContactId) (*http.Request, error) {
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

// responderForSecurityContactsDelete handles the response to the SecurityContactsDelete request. The method always
// closes the http.Response Body.
func (c SecurityContactsClient) responderForSecurityContactsDelete(resp *http.Response) (result SecurityContactsDeleteOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusNoContent),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

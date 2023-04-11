package securitycontacts

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityContactsCreateOperationResponse struct {
	HttpResponse *http.Response
	Model        *SecurityContact
}

// SecurityContactsCreate ...
func (c SecurityContactsClient) SecurityContactsCreate(ctx context.Context, id SecurityContactId, input SecurityContact) (result SecurityContactsCreateOperationResponse, err error) {
	req, err := c.preparerForSecurityContactsCreate(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securitycontacts.SecurityContactsClient", "SecurityContactsCreate", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "securitycontacts.SecurityContactsClient", "SecurityContactsCreate", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForSecurityContactsCreate(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securitycontacts.SecurityContactsClient", "SecurityContactsCreate", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForSecurityContactsCreate prepares the SecurityContactsCreate request.
func (c SecurityContactsClient) preparerForSecurityContactsCreate(ctx context.Context, id SecurityContactId, input SecurityContact) (*http.Request, error) {
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

// responderForSecurityContactsCreate handles the response to the SecurityContactsCreate request. The method always
// closes the http.Response Body.
func (c SecurityContactsClient) responderForSecurityContactsCreate(resp *http.Response) (result SecurityContactsCreateOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

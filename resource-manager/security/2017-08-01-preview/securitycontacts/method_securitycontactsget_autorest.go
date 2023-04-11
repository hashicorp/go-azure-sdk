package securitycontacts

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityContactsGetOperationResponse struct {
	HttpResponse *http.Response
	Model        *SecurityContact
}

// SecurityContactsGet ...
func (c SecurityContactsClient) SecurityContactsGet(ctx context.Context, id SecurityContactId) (result SecurityContactsGetOperationResponse, err error) {
	req, err := c.preparerForSecurityContactsGet(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securitycontacts.SecurityContactsClient", "SecurityContactsGet", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "securitycontacts.SecurityContactsClient", "SecurityContactsGet", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForSecurityContactsGet(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securitycontacts.SecurityContactsClient", "SecurityContactsGet", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForSecurityContactsGet prepares the SecurityContactsGet request.
func (c SecurityContactsClient) preparerForSecurityContactsGet(ctx context.Context, id SecurityContactId) (*http.Request, error) {
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

// responderForSecurityContactsGet handles the response to the SecurityContactsGet request. The method always
// closes the http.Response Body.
func (c SecurityContactsClient) responderForSecurityContactsGet(resp *http.Response) (result SecurityContactsGetOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

package securitycontacts

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityContactsUpdateOperationResponse struct {
	HttpResponse *http.Response
	Model        *SecurityContact
}

// SecurityContactsUpdate ...
func (c SecurityContactsClient) SecurityContactsUpdate(ctx context.Context, id SecurityContactId, input SecurityContact) (result SecurityContactsUpdateOperationResponse, err error) {
	req, err := c.preparerForSecurityContactsUpdate(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securitycontacts.SecurityContactsClient", "SecurityContactsUpdate", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "securitycontacts.SecurityContactsClient", "SecurityContactsUpdate", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForSecurityContactsUpdate(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securitycontacts.SecurityContactsClient", "SecurityContactsUpdate", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForSecurityContactsUpdate prepares the SecurityContactsUpdate request.
func (c SecurityContactsClient) preparerForSecurityContactsUpdate(ctx context.Context, id SecurityContactId, input SecurityContact) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForSecurityContactsUpdate handles the response to the SecurityContactsUpdate request. The method always
// closes the http.Response Body.
func (c SecurityContactsClient) responderForSecurityContactsUpdate(resp *http.Response) (result SecurityContactsUpdateOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

package emailregistration

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ActivateEmailOperationResponse struct {
	HttpResponse *http.Response
	Model        *EmailRegistration
}

// ActivateEmail ...
func (c EmailRegistrationClient) ActivateEmail(ctx context.Context, id LocationId, input EmailRegistration) (result ActivateEmailOperationResponse, err error) {
	req, err := c.preparerForActivateEmail(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "emailregistration.EmailRegistrationClient", "ActivateEmail", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "emailregistration.EmailRegistrationClient", "ActivateEmail", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForActivateEmail(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "emailregistration.EmailRegistrationClient", "ActivateEmail", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForActivateEmail prepares the ActivateEmail request.
func (c EmailRegistrationClient) preparerForActivateEmail(ctx context.Context, id LocationId, input EmailRegistration) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/activateEmail", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForActivateEmail handles the response to the ActivateEmail request. The method always
// closes the http.Response Body.
func (c EmailRegistrationClient) responderForActivateEmail(resp *http.Response) (result ActivateEmailOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

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

type RegisterEmailOperationResponse struct {
	HttpResponse *http.Response
	Model        *EmailRegistration
}

// RegisterEmail ...
func (c EmailRegistrationClient) RegisterEmail(ctx context.Context, id LocationId) (result RegisterEmailOperationResponse, err error) {
	req, err := c.preparerForRegisterEmail(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "emailregistration.EmailRegistrationClient", "RegisterEmail", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "emailregistration.EmailRegistrationClient", "RegisterEmail", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForRegisterEmail(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "emailregistration.EmailRegistrationClient", "RegisterEmail", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForRegisterEmail prepares the RegisterEmail request.
func (c EmailRegistrationClient) preparerForRegisterEmail(ctx context.Context, id LocationId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/registerEmail", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForRegisterEmail handles the response to the RegisterEmail request. The method always
// closes the http.Response Body.
func (c EmailRegistrationClient) responderForRegisterEmail(resp *http.Response) (result RegisterEmailOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

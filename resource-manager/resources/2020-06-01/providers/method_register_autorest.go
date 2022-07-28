package providers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RegisterOperationResponse struct {
	HttpResponse *http.Response
	Model        *Provider
}

// Register ...
func (c ProvidersClient) Register(ctx context.Context, id SubscriptionProviderId) (result RegisterOperationResponse, err error) {
	req, err := c.preparerForRegister(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "providers.ProvidersClient", "Register", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "providers.ProvidersClient", "Register", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForRegister(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "providers.ProvidersClient", "Register", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForRegister prepares the Register request.
func (c ProvidersClient) preparerForRegister(ctx context.Context, id SubscriptionProviderId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/register", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForRegister handles the response to the Register request. The method always
// closes the http.Response Body.
func (c ProvidersClient) responderForRegister(resp *http.Response) (result RegisterOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

package attestations

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AttestationsDeleteAtSubscriptionOperationResponse struct {
	HttpResponse *http.Response
}

// AttestationsDeleteAtSubscription ...
func (c AttestationsClient) AttestationsDeleteAtSubscription(ctx context.Context, id AttestationId) (result AttestationsDeleteAtSubscriptionOperationResponse, err error) {
	req, err := c.preparerForAttestationsDeleteAtSubscription(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "attestations.AttestationsClient", "AttestationsDeleteAtSubscription", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "attestations.AttestationsClient", "AttestationsDeleteAtSubscription", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForAttestationsDeleteAtSubscription(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "attestations.AttestationsClient", "AttestationsDeleteAtSubscription", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForAttestationsDeleteAtSubscription prepares the AttestationsDeleteAtSubscription request.
func (c AttestationsClient) preparerForAttestationsDeleteAtSubscription(ctx context.Context, id AttestationId) (*http.Request, error) {
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

// responderForAttestationsDeleteAtSubscription handles the response to the AttestationsDeleteAtSubscription request. The method always
// closes the http.Response Body.
func (c AttestationsClient) responderForAttestationsDeleteAtSubscription(resp *http.Response) (result AttestationsDeleteAtSubscriptionOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusNoContent, http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

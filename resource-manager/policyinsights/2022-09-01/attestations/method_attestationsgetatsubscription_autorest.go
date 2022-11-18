package attestations

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AttestationsGetAtSubscriptionOperationResponse struct {
	HttpResponse *http.Response
	Model        *Attestation
}

// AttestationsGetAtSubscription ...
func (c AttestationsClient) AttestationsGetAtSubscription(ctx context.Context, id AttestationId) (result AttestationsGetAtSubscriptionOperationResponse, err error) {
	req, err := c.preparerForAttestationsGetAtSubscription(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "attestations.AttestationsClient", "AttestationsGetAtSubscription", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "attestations.AttestationsClient", "AttestationsGetAtSubscription", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForAttestationsGetAtSubscription(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "attestations.AttestationsClient", "AttestationsGetAtSubscription", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForAttestationsGetAtSubscription prepares the AttestationsGetAtSubscription request.
func (c AttestationsClient) preparerForAttestationsGetAtSubscription(ctx context.Context, id AttestationId) (*http.Request, error) {
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

// responderForAttestationsGetAtSubscription handles the response to the AttestationsGetAtSubscription request. The method always
// closes the http.Response Body.
func (c AttestationsClient) responderForAttestationsGetAtSubscription(resp *http.Response) (result AttestationsGetAtSubscriptionOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

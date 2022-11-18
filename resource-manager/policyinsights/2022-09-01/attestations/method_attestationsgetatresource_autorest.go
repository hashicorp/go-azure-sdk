package attestations

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AttestationsGetAtResourceOperationResponse struct {
	HttpResponse *http.Response
	Model        *Attestation
}

// AttestationsGetAtResource ...
func (c AttestationsClient) AttestationsGetAtResource(ctx context.Context, id ScopedAttestationId) (result AttestationsGetAtResourceOperationResponse, err error) {
	req, err := c.preparerForAttestationsGetAtResource(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "attestations.AttestationsClient", "AttestationsGetAtResource", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "attestations.AttestationsClient", "AttestationsGetAtResource", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForAttestationsGetAtResource(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "attestations.AttestationsClient", "AttestationsGetAtResource", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForAttestationsGetAtResource prepares the AttestationsGetAtResource request.
func (c AttestationsClient) preparerForAttestationsGetAtResource(ctx context.Context, id ScopedAttestationId) (*http.Request, error) {
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

// responderForAttestationsGetAtResource handles the response to the AttestationsGetAtResource request. The method always
// closes the http.Response Body.
func (c AttestationsClient) responderForAttestationsGetAtResource(resp *http.Response) (result AttestationsGetAtResourceOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

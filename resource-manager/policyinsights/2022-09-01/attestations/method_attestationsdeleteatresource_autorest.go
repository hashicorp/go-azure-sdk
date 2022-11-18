package attestations

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AttestationsDeleteAtResourceOperationResponse struct {
	HttpResponse *http.Response
}

// AttestationsDeleteAtResource ...
func (c AttestationsClient) AttestationsDeleteAtResource(ctx context.Context, id ScopedAttestationId) (result AttestationsDeleteAtResourceOperationResponse, err error) {
	req, err := c.preparerForAttestationsDeleteAtResource(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "attestations.AttestationsClient", "AttestationsDeleteAtResource", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "attestations.AttestationsClient", "AttestationsDeleteAtResource", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForAttestationsDeleteAtResource(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "attestations.AttestationsClient", "AttestationsDeleteAtResource", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForAttestationsDeleteAtResource prepares the AttestationsDeleteAtResource request.
func (c AttestationsClient) preparerForAttestationsDeleteAtResource(ctx context.Context, id ScopedAttestationId) (*http.Request, error) {
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

// responderForAttestationsDeleteAtResource handles the response to the AttestationsDeleteAtResource request. The method always
// closes the http.Response Body.
func (c AttestationsClient) responderForAttestationsDeleteAtResource(resp *http.Response) (result AttestationsDeleteAtResourceOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusNoContent, http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

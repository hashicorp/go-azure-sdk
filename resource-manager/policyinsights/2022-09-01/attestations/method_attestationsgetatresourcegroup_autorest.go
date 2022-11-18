package attestations

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AttestationsGetAtResourceGroupOperationResponse struct {
	HttpResponse *http.Response
	Model        *Attestation
}

// AttestationsGetAtResourceGroup ...
func (c AttestationsClient) AttestationsGetAtResourceGroup(ctx context.Context, id ProviderAttestationId) (result AttestationsGetAtResourceGroupOperationResponse, err error) {
	req, err := c.preparerForAttestationsGetAtResourceGroup(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "attestations.AttestationsClient", "AttestationsGetAtResourceGroup", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "attestations.AttestationsClient", "AttestationsGetAtResourceGroup", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForAttestationsGetAtResourceGroup(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "attestations.AttestationsClient", "AttestationsGetAtResourceGroup", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForAttestationsGetAtResourceGroup prepares the AttestationsGetAtResourceGroup request.
func (c AttestationsClient) preparerForAttestationsGetAtResourceGroup(ctx context.Context, id ProviderAttestationId) (*http.Request, error) {
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

// responderForAttestationsGetAtResourceGroup handles the response to the AttestationsGetAtResourceGroup request. The method always
// closes the http.Response Body.
func (c AttestationsClient) responderForAttestationsGetAtResourceGroup(resp *http.Response) (result AttestationsGetAtResourceGroupOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

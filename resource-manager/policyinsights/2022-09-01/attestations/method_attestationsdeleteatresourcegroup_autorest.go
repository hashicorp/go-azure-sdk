package attestations

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AttestationsDeleteAtResourceGroupOperationResponse struct {
	HttpResponse *http.Response
}

// AttestationsDeleteAtResourceGroup ...
func (c AttestationsClient) AttestationsDeleteAtResourceGroup(ctx context.Context, id ProviderAttestationId) (result AttestationsDeleteAtResourceGroupOperationResponse, err error) {
	req, err := c.preparerForAttestationsDeleteAtResourceGroup(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "attestations.AttestationsClient", "AttestationsDeleteAtResourceGroup", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "attestations.AttestationsClient", "AttestationsDeleteAtResourceGroup", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForAttestationsDeleteAtResourceGroup(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "attestations.AttestationsClient", "AttestationsDeleteAtResourceGroup", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForAttestationsDeleteAtResourceGroup prepares the AttestationsDeleteAtResourceGroup request.
func (c AttestationsClient) preparerForAttestationsDeleteAtResourceGroup(ctx context.Context, id ProviderAttestationId) (*http.Request, error) {
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

// responderForAttestationsDeleteAtResourceGroup handles the response to the AttestationsDeleteAtResourceGroup request. The method always
// closes the http.Response Body.
func (c AttestationsClient) responderForAttestationsDeleteAtResourceGroup(resp *http.Response) (result AttestationsDeleteAtResourceGroupOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusNoContent, http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

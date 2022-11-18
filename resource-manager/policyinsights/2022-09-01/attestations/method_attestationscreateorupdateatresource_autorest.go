package attestations

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/hashicorp/go-azure-helpers/polling"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AttestationsCreateOrUpdateAtResourceOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// AttestationsCreateOrUpdateAtResource ...
func (c AttestationsClient) AttestationsCreateOrUpdateAtResource(ctx context.Context, id ScopedAttestationId, input Attestation) (result AttestationsCreateOrUpdateAtResourceOperationResponse, err error) {
	req, err := c.preparerForAttestationsCreateOrUpdateAtResource(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "attestations.AttestationsClient", "AttestationsCreateOrUpdateAtResource", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForAttestationsCreateOrUpdateAtResource(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "attestations.AttestationsClient", "AttestationsCreateOrUpdateAtResource", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// AttestationsCreateOrUpdateAtResourceThenPoll performs AttestationsCreateOrUpdateAtResource then polls until it's completed
func (c AttestationsClient) AttestationsCreateOrUpdateAtResourceThenPoll(ctx context.Context, id ScopedAttestationId, input Attestation) error {
	result, err := c.AttestationsCreateOrUpdateAtResource(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing AttestationsCreateOrUpdateAtResource: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after AttestationsCreateOrUpdateAtResource: %+v", err)
	}

	return nil
}

// preparerForAttestationsCreateOrUpdateAtResource prepares the AttestationsCreateOrUpdateAtResource request.
func (c AttestationsClient) preparerForAttestationsCreateOrUpdateAtResource(ctx context.Context, id ScopedAttestationId, input Attestation) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForAttestationsCreateOrUpdateAtResource sends the AttestationsCreateOrUpdateAtResource request. The method will close the
// http.Response Body if it receives an error.
func (c AttestationsClient) senderForAttestationsCreateOrUpdateAtResource(ctx context.Context, req *http.Request) (future AttestationsCreateOrUpdateAtResourceOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}

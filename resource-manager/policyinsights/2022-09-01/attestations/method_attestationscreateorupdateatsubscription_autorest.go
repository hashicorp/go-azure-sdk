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

type AttestationsCreateOrUpdateAtSubscriptionOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// AttestationsCreateOrUpdateAtSubscription ...
func (c AttestationsClient) AttestationsCreateOrUpdateAtSubscription(ctx context.Context, id AttestationId, input Attestation) (result AttestationsCreateOrUpdateAtSubscriptionOperationResponse, err error) {
	req, err := c.preparerForAttestationsCreateOrUpdateAtSubscription(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "attestations.AttestationsClient", "AttestationsCreateOrUpdateAtSubscription", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForAttestationsCreateOrUpdateAtSubscription(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "attestations.AttestationsClient", "AttestationsCreateOrUpdateAtSubscription", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// AttestationsCreateOrUpdateAtSubscriptionThenPoll performs AttestationsCreateOrUpdateAtSubscription then polls until it's completed
func (c AttestationsClient) AttestationsCreateOrUpdateAtSubscriptionThenPoll(ctx context.Context, id AttestationId, input Attestation) error {
	result, err := c.AttestationsCreateOrUpdateAtSubscription(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing AttestationsCreateOrUpdateAtSubscription: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after AttestationsCreateOrUpdateAtSubscription: %+v", err)
	}

	return nil
}

// preparerForAttestationsCreateOrUpdateAtSubscription prepares the AttestationsCreateOrUpdateAtSubscription request.
func (c AttestationsClient) preparerForAttestationsCreateOrUpdateAtSubscription(ctx context.Context, id AttestationId, input Attestation) (*http.Request, error) {
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

// senderForAttestationsCreateOrUpdateAtSubscription sends the AttestationsCreateOrUpdateAtSubscription request. The method will close the
// http.Response Body if it receives an error.
func (c AttestationsClient) senderForAttestationsCreateOrUpdateAtSubscription(ctx context.Context, req *http.Request) (future AttestationsCreateOrUpdateAtSubscriptionOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}

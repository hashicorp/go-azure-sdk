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

type AttestationsCreateOrUpdateAtResourceGroupOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// AttestationsCreateOrUpdateAtResourceGroup ...
func (c AttestationsClient) AttestationsCreateOrUpdateAtResourceGroup(ctx context.Context, id ProviderAttestationId, input Attestation) (result AttestationsCreateOrUpdateAtResourceGroupOperationResponse, err error) {
	req, err := c.preparerForAttestationsCreateOrUpdateAtResourceGroup(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "attestations.AttestationsClient", "AttestationsCreateOrUpdateAtResourceGroup", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForAttestationsCreateOrUpdateAtResourceGroup(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "attestations.AttestationsClient", "AttestationsCreateOrUpdateAtResourceGroup", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// AttestationsCreateOrUpdateAtResourceGroupThenPoll performs AttestationsCreateOrUpdateAtResourceGroup then polls until it's completed
func (c AttestationsClient) AttestationsCreateOrUpdateAtResourceGroupThenPoll(ctx context.Context, id ProviderAttestationId, input Attestation) error {
	result, err := c.AttestationsCreateOrUpdateAtResourceGroup(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing AttestationsCreateOrUpdateAtResourceGroup: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after AttestationsCreateOrUpdateAtResourceGroup: %+v", err)
	}

	return nil
}

// preparerForAttestationsCreateOrUpdateAtResourceGroup prepares the AttestationsCreateOrUpdateAtResourceGroup request.
func (c AttestationsClient) preparerForAttestationsCreateOrUpdateAtResourceGroup(ctx context.Context, id ProviderAttestationId, input Attestation) (*http.Request, error) {
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

// senderForAttestationsCreateOrUpdateAtResourceGroup sends the AttestationsCreateOrUpdateAtResourceGroup request. The method will close the
// http.Response Body if it receives an error.
func (c AttestationsClient) senderForAttestationsCreateOrUpdateAtResourceGroup(ctx context.Context, req *http.Request) (future AttestationsCreateOrUpdateAtResourceGroupOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}

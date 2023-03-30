package domains

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

type InitiateVerificationOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// InitiateVerification ...
func (c DomainsClient) InitiateVerification(ctx context.Context, id DomainId, input VerificationParameter) (result InitiateVerificationOperationResponse, err error) {
	req, err := c.preparerForInitiateVerification(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "domains.DomainsClient", "InitiateVerification", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForInitiateVerification(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "domains.DomainsClient", "InitiateVerification", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// InitiateVerificationThenPoll performs InitiateVerification then polls until it's completed
func (c DomainsClient) InitiateVerificationThenPoll(ctx context.Context, id DomainId, input VerificationParameter) error {
	result, err := c.InitiateVerification(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing InitiateVerification: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after InitiateVerification: %+v", err)
	}

	return nil
}

// preparerForInitiateVerification prepares the InitiateVerification request.
func (c DomainsClient) preparerForInitiateVerification(ctx context.Context, id DomainId, input VerificationParameter) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/initiateVerification", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForInitiateVerification sends the InitiateVerification request. The method will close the
// http.Response Body if it receives an error.
func (c DomainsClient) senderForInitiateVerification(ctx context.Context, req *http.Request) (future InitiateVerificationOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}

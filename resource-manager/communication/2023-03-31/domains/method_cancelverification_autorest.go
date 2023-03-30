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

type CancelVerificationOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// CancelVerification ...
func (c DomainsClient) CancelVerification(ctx context.Context, id DomainId, input VerificationParameter) (result CancelVerificationOperationResponse, err error) {
	req, err := c.preparerForCancelVerification(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "domains.DomainsClient", "CancelVerification", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForCancelVerification(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "domains.DomainsClient", "CancelVerification", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// CancelVerificationThenPoll performs CancelVerification then polls until it's completed
func (c DomainsClient) CancelVerificationThenPoll(ctx context.Context, id DomainId, input VerificationParameter) error {
	result, err := c.CancelVerification(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing CancelVerification: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after CancelVerification: %+v", err)
	}

	return nil
}

// preparerForCancelVerification prepares the CancelVerification request.
func (c DomainsClient) preparerForCancelVerification(ctx context.Context, id DomainId, input VerificationParameter) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/cancelVerification", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForCancelVerification sends the CancelVerification request. The method will close the
// http.Response Body if it receives an error.
func (c DomainsClient) senderForCancelVerification(ctx context.Context, req *http.Request) (future CancelVerificationOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}

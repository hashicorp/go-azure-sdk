package sharesubscription

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

type CancelSynchronizationOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// CancelSynchronization ...
func (c ShareSubscriptionClient) CancelSynchronization(ctx context.Context, id ShareSubscriptionId, input ShareSubscriptionSynchronization) (result CancelSynchronizationOperationResponse, err error) {
	req, err := c.preparerForCancelSynchronization(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sharesubscription.ShareSubscriptionClient", "CancelSynchronization", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForCancelSynchronization(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sharesubscription.ShareSubscriptionClient", "CancelSynchronization", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// CancelSynchronizationThenPoll performs CancelSynchronization then polls until it's completed
func (c ShareSubscriptionClient) CancelSynchronizationThenPoll(ctx context.Context, id ShareSubscriptionId, input ShareSubscriptionSynchronization) error {
	result, err := c.CancelSynchronization(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing CancelSynchronization: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after CancelSynchronization: %+v", err)
	}

	return nil
}

// preparerForCancelSynchronization prepares the CancelSynchronization request.
func (c ShareSubscriptionClient) preparerForCancelSynchronization(ctx context.Context, id ShareSubscriptionId, input ShareSubscriptionSynchronization) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/cancelSynchronization", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForCancelSynchronization sends the CancelSynchronization request. The method will close the
// http.Response Body if it receives an error.
func (c ShareSubscriptionClient) senderForCancelSynchronization(ctx context.Context, req *http.Request) (future CancelSynchronizationOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}

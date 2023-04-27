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

type SynchronizeOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// Synchronize ...
func (c ShareSubscriptionClient) Synchronize(ctx context.Context, id ShareSubscriptionId, input Synchronize) (result SynchronizeOperationResponse, err error) {
	req, err := c.preparerForSynchronize(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sharesubscription.ShareSubscriptionClient", "Synchronize", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForSynchronize(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sharesubscription.ShareSubscriptionClient", "Synchronize", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// SynchronizeThenPoll performs Synchronize then polls until it's completed
func (c ShareSubscriptionClient) SynchronizeThenPoll(ctx context.Context, id ShareSubscriptionId, input Synchronize) error {
	result, err := c.Synchronize(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing Synchronize: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after Synchronize: %+v", err)
	}

	return nil
}

// preparerForSynchronize prepares the Synchronize request.
func (c ShareSubscriptionClient) preparerForSynchronize(ctx context.Context, id ShareSubscriptionId, input Synchronize) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/synchronize", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForSynchronize sends the Synchronize request. The method will close the
// http.Response Body if it receives an error.
func (c ShareSubscriptionClient) senderForSynchronize(ctx context.Context, req *http.Request) (future SynchronizeOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}

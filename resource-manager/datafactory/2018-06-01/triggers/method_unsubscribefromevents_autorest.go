package triggers

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

type UnsubscribeFromEventsOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// UnsubscribeFromEvents ...
func (c TriggersClient) UnsubscribeFromEvents(ctx context.Context, id TriggerId) (result UnsubscribeFromEventsOperationResponse, err error) {
	req, err := c.preparerForUnsubscribeFromEvents(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "triggers.TriggersClient", "UnsubscribeFromEvents", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForUnsubscribeFromEvents(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "triggers.TriggersClient", "UnsubscribeFromEvents", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// UnsubscribeFromEventsThenPoll performs UnsubscribeFromEvents then polls until it's completed
func (c TriggersClient) UnsubscribeFromEventsThenPoll(ctx context.Context, id TriggerId) error {
	result, err := c.UnsubscribeFromEvents(ctx, id)
	if err != nil {
		return fmt.Errorf("performing UnsubscribeFromEvents: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after UnsubscribeFromEvents: %+v", err)
	}

	return nil
}

// preparerForUnsubscribeFromEvents prepares the UnsubscribeFromEvents request.
func (c TriggersClient) preparerForUnsubscribeFromEvents(ctx context.Context, id TriggerId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/unsubscribeFromEvents", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForUnsubscribeFromEvents sends the UnsubscribeFromEvents request. The method will close the
// http.Response Body if it receives an error.
func (c TriggersClient) senderForUnsubscribeFromEvents(ctx context.Context, req *http.Request) (future UnsubscribeFromEventsOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}

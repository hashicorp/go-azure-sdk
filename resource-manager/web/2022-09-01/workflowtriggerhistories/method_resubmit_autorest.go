package workflowtriggerhistories

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

type ResubmitOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// Resubmit ...
func (c WorkflowTriggerHistoriesClient) Resubmit(ctx context.Context, id TriggerHistoryId) (result ResubmitOperationResponse, err error) {
	req, err := c.preparerForResubmit(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workflowtriggerhistories.WorkflowTriggerHistoriesClient", "Resubmit", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForResubmit(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workflowtriggerhistories.WorkflowTriggerHistoriesClient", "Resubmit", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// ResubmitThenPoll performs Resubmit then polls until it's completed
func (c WorkflowTriggerHistoriesClient) ResubmitThenPoll(ctx context.Context, id TriggerHistoryId) error {
	result, err := c.Resubmit(ctx, id)
	if err != nil {
		return fmt.Errorf("performing Resubmit: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after Resubmit: %+v", err)
	}

	return nil
}

// preparerForResubmit prepares the Resubmit request.
func (c WorkflowTriggerHistoriesClient) preparerForResubmit(ctx context.Context, id TriggerHistoryId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/resubmit", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForResubmit sends the Resubmit request. The method will close the
// http.Response Body if it receives an error.
func (c WorkflowTriggerHistoriesClient) senderForResubmit(ctx context.Context, req *http.Request) (future ResubmitOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}

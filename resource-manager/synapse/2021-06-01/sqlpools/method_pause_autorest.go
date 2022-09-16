package sqlpools

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

type PauseOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// Pause ...
func (c SqlPoolsClient) Pause(ctx context.Context, id SqlPoolId) (result PauseOperationResponse, err error) {
	req, err := c.preparerForPause(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlpools.SqlPoolsClient", "Pause", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForPause(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlpools.SqlPoolsClient", "Pause", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// PauseThenPoll performs Pause then polls until it's completed
func (c SqlPoolsClient) PauseThenPoll(ctx context.Context, id SqlPoolId) error {
	result, err := c.Pause(ctx, id)
	if err != nil {
		return fmt.Errorf("performing Pause: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after Pause: %+v", err)
	}

	return nil
}

// preparerForPause prepares the Pause request.
func (c SqlPoolsClient) preparerForPause(ctx context.Context, id SqlPoolId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/pause", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForPause sends the Pause request. The method will close the
// http.Response Body if it receives an error.
func (c SqlPoolsClient) senderForPause(ctx context.Context, req *http.Request) (future PauseOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}

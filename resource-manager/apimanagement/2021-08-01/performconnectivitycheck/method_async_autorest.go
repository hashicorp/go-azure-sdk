package performconnectivitycheck

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

type AsyncOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// Async ...
func (c PerformConnectivityCheckClient) Async(ctx context.Context, id ServiceId, input ConnectivityCheckRequest) (result AsyncOperationResponse, err error) {
	req, err := c.preparerForAsync(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "performconnectivitycheck.PerformConnectivityCheckClient", "Async", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForAsync(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "performconnectivitycheck.PerformConnectivityCheckClient", "Async", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// AsyncThenPoll performs Async then polls until it's completed
func (c PerformConnectivityCheckClient) AsyncThenPoll(ctx context.Context, id ServiceId, input ConnectivityCheckRequest) error {
	result, err := c.Async(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing Async: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after Async: %+v", err)
	}

	return nil
}

// preparerForAsync prepares the Async request.
func (c PerformConnectivityCheckClient) preparerForAsync(ctx context.Context, id ServiceId, input ConnectivityCheckRequest) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/connectivityCheck", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForAsync sends the Async request. The method will close the
// http.Response Body if it receives an error.
func (c PerformConnectivityCheckClient) senderForAsync(ctx context.Context, req *http.Request) (future AsyncOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}

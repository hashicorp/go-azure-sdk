package servergroupoperations

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

type ServerGroupsStartOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// ServerGroupsStart ...
func (c ServerGroupOperationsClient) ServerGroupsStart(ctx context.Context, id ServerGroupsv2Id) (result ServerGroupsStartOperationResponse, err error) {
	req, err := c.preparerForServerGroupsStart(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servergroupoperations.ServerGroupOperationsClient", "ServerGroupsStart", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForServerGroupsStart(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servergroupoperations.ServerGroupOperationsClient", "ServerGroupsStart", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// ServerGroupsStartThenPoll performs ServerGroupsStart then polls until it's completed
func (c ServerGroupOperationsClient) ServerGroupsStartThenPoll(ctx context.Context, id ServerGroupsv2Id) error {
	result, err := c.ServerGroupsStart(ctx, id)
	if err != nil {
		return fmt.Errorf("performing ServerGroupsStart: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after ServerGroupsStart: %+v", err)
	}

	return nil
}

// preparerForServerGroupsStart prepares the ServerGroupsStart request.
func (c ServerGroupOperationsClient) preparerForServerGroupsStart(ctx context.Context, id ServerGroupsv2Id) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/start", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForServerGroupsStart sends the ServerGroupsStart request. The method will close the
// http.Response Body if it receives an error.
func (c ServerGroupOperationsClient) senderForServerGroupsStart(ctx context.Context, req *http.Request) (future ServerGroupsStartOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}

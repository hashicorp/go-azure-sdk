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

type ServerGroupsStopOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// ServerGroupsStop ...
func (c ServerGroupOperationsClient) ServerGroupsStop(ctx context.Context, id ServerGroupsv2Id) (result ServerGroupsStopOperationResponse, err error) {
	req, err := c.preparerForServerGroupsStop(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servergroupoperations.ServerGroupOperationsClient", "ServerGroupsStop", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForServerGroupsStop(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servergroupoperations.ServerGroupOperationsClient", "ServerGroupsStop", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// ServerGroupsStopThenPoll performs ServerGroupsStop then polls until it's completed
func (c ServerGroupOperationsClient) ServerGroupsStopThenPoll(ctx context.Context, id ServerGroupsv2Id) error {
	result, err := c.ServerGroupsStop(ctx, id)
	if err != nil {
		return fmt.Errorf("performing ServerGroupsStop: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after ServerGroupsStop: %+v", err)
	}

	return nil
}

// preparerForServerGroupsStop prepares the ServerGroupsStop request.
func (c ServerGroupOperationsClient) preparerForServerGroupsStop(ctx context.Context, id ServerGroupsv2Id) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/stop", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForServerGroupsStop sends the ServerGroupsStop request. The method will close the
// http.Response Body if it receives an error.
func (c ServerGroupOperationsClient) senderForServerGroupsStop(ctx context.Context, req *http.Request) (future ServerGroupsStopOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}

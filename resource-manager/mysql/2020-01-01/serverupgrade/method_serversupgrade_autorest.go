package serverupgrade

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

type ServersUpgradeOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// ServersUpgrade ...
func (c ServerUpgradeClient) ServersUpgrade(ctx context.Context, id ServerId, input ServerUpgradeParameters) (result ServersUpgradeOperationResponse, err error) {
	req, err := c.preparerForServersUpgrade(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "serverupgrade.ServerUpgradeClient", "ServersUpgrade", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForServersUpgrade(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "serverupgrade.ServerUpgradeClient", "ServersUpgrade", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// ServersUpgradeThenPoll performs ServersUpgrade then polls until it's completed
func (c ServerUpgradeClient) ServersUpgradeThenPoll(ctx context.Context, id ServerId, input ServerUpgradeParameters) error {
	result, err := c.ServersUpgrade(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing ServersUpgrade: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after ServersUpgrade: %+v", err)
	}

	return nil
}

// preparerForServersUpgrade prepares the ServersUpgrade request.
func (c ServerUpgradeClient) preparerForServersUpgrade(ctx context.Context, id ServerId, input ServerUpgradeParameters) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/upgrade", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForServersUpgrade sends the ServersUpgrade request. The method will close the
// http.Response Body if it receives an error.
func (c ServerUpgradeClient) senderForServersUpgrade(ctx context.Context, req *http.Request) (future ServersUpgradeOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}

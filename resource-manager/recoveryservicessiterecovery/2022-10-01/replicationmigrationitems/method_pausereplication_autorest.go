package replicationmigrationitems

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

type PauseReplicationOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// PauseReplication ...
func (c ReplicationMigrationItemsClient) PauseReplication(ctx context.Context, id ReplicationMigrationItemId, input PauseReplicationInput) (result PauseReplicationOperationResponse, err error) {
	req, err := c.preparerForPauseReplication(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "replicationmigrationitems.ReplicationMigrationItemsClient", "PauseReplication", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForPauseReplication(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "replicationmigrationitems.ReplicationMigrationItemsClient", "PauseReplication", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// PauseReplicationThenPoll performs PauseReplication then polls until it's completed
func (c ReplicationMigrationItemsClient) PauseReplicationThenPoll(ctx context.Context, id ReplicationMigrationItemId, input PauseReplicationInput) error {
	result, err := c.PauseReplication(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing PauseReplication: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after PauseReplication: %+v", err)
	}

	return nil
}

// preparerForPauseReplication prepares the PauseReplication request.
func (c ReplicationMigrationItemsClient) preparerForPauseReplication(ctx context.Context, id ReplicationMigrationItemId, input PauseReplicationInput) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/pauseReplication", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForPauseReplication sends the PauseReplication request. The method will close the
// http.Response Body if it receives an error.
func (c ReplicationMigrationItemsClient) senderForPauseReplication(ctx context.Context, req *http.Request) (future PauseReplicationOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}

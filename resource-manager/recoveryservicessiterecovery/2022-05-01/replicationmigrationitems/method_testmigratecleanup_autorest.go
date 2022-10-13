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

type TestMigrateCleanupOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// TestMigrateCleanup ...
func (c ReplicationMigrationItemsClient) TestMigrateCleanup(ctx context.Context, id ReplicationMigrationItemId, input TestMigrateCleanupInput) (result TestMigrateCleanupOperationResponse, err error) {
	req, err := c.preparerForTestMigrateCleanup(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "replicationmigrationitems.ReplicationMigrationItemsClient", "TestMigrateCleanup", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForTestMigrateCleanup(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "replicationmigrationitems.ReplicationMigrationItemsClient", "TestMigrateCleanup", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// TestMigrateCleanupThenPoll performs TestMigrateCleanup then polls until it's completed
func (c ReplicationMigrationItemsClient) TestMigrateCleanupThenPoll(ctx context.Context, id ReplicationMigrationItemId, input TestMigrateCleanupInput) error {
	result, err := c.TestMigrateCleanup(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing TestMigrateCleanup: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after TestMigrateCleanup: %+v", err)
	}

	return nil
}

// preparerForTestMigrateCleanup prepares the TestMigrateCleanup request.
func (c ReplicationMigrationItemsClient) preparerForTestMigrateCleanup(ctx context.Context, id ReplicationMigrationItemId, input TestMigrateCleanupInput) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/testMigrateCleanup", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForTestMigrateCleanup sends the TestMigrateCleanup request. The method will close the
// http.Response Body if it receives an error.
func (c ReplicationMigrationItemsClient) senderForTestMigrateCleanup(ctx context.Context, req *http.Request) (future TestMigrateCleanupOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}

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

type TestMigrateOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// TestMigrate ...
func (c ReplicationMigrationItemsClient) TestMigrate(ctx context.Context, id ReplicationMigrationItemId, input TestMigrateInput) (result TestMigrateOperationResponse, err error) {
	req, err := c.preparerForTestMigrate(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "replicationmigrationitems.ReplicationMigrationItemsClient", "TestMigrate", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForTestMigrate(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "replicationmigrationitems.ReplicationMigrationItemsClient", "TestMigrate", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// TestMigrateThenPoll performs TestMigrate then polls until it's completed
func (c ReplicationMigrationItemsClient) TestMigrateThenPoll(ctx context.Context, id ReplicationMigrationItemId, input TestMigrateInput) error {
	result, err := c.TestMigrate(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing TestMigrate: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after TestMigrate: %+v", err)
	}

	return nil
}

// preparerForTestMigrate prepares the TestMigrate request.
func (c ReplicationMigrationItemsClient) preparerForTestMigrate(ctx context.Context, id ReplicationMigrationItemId, input TestMigrateInput) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/testMigrate", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForTestMigrate sends the TestMigrate request. The method will close the
// http.Response Body if it receives an error.
func (c ReplicationMigrationItemsClient) senderForTestMigrate(ctx context.Context, req *http.Request) (future TestMigrateOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}

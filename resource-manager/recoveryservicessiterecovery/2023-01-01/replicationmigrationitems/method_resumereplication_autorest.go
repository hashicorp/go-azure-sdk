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

type ResumeReplicationOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// ResumeReplication ...
func (c ReplicationMigrationItemsClient) ResumeReplication(ctx context.Context, id ReplicationMigrationItemId, input ResumeReplicationInput) (result ResumeReplicationOperationResponse, err error) {
	req, err := c.preparerForResumeReplication(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "replicationmigrationitems.ReplicationMigrationItemsClient", "ResumeReplication", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForResumeReplication(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "replicationmigrationitems.ReplicationMigrationItemsClient", "ResumeReplication", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// ResumeReplicationThenPoll performs ResumeReplication then polls until it's completed
func (c ReplicationMigrationItemsClient) ResumeReplicationThenPoll(ctx context.Context, id ReplicationMigrationItemId, input ResumeReplicationInput) error {
	result, err := c.ResumeReplication(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing ResumeReplication: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after ResumeReplication: %+v", err)
	}

	return nil
}

// preparerForResumeReplication prepares the ResumeReplication request.
func (c ReplicationMigrationItemsClient) preparerForResumeReplication(ctx context.Context, id ReplicationMigrationItemId, input ResumeReplicationInput) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/resumeReplication", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForResumeReplication sends the ResumeReplication request. The method will close the
// http.Response Body if it receives an error.
func (c ReplicationMigrationItemsClient) senderForResumeReplication(ctx context.Context, req *http.Request) (future ResumeReplicationOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}

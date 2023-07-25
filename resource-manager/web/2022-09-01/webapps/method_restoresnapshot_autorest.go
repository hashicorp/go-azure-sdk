package webapps

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

type RestoreSnapshotOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// RestoreSnapshot ...
func (c WebAppsClient) RestoreSnapshot(ctx context.Context, id SiteId, input SnapshotRestoreRequest) (result RestoreSnapshotOperationResponse, err error) {
	req, err := c.preparerForRestoreSnapshot(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "RestoreSnapshot", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForRestoreSnapshot(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "RestoreSnapshot", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// RestoreSnapshotThenPoll performs RestoreSnapshot then polls until it's completed
func (c WebAppsClient) RestoreSnapshotThenPoll(ctx context.Context, id SiteId, input SnapshotRestoreRequest) error {
	result, err := c.RestoreSnapshot(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing RestoreSnapshot: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after RestoreSnapshot: %+v", err)
	}

	return nil
}

// preparerForRestoreSnapshot prepares the RestoreSnapshot request.
func (c WebAppsClient) preparerForRestoreSnapshot(ctx context.Context, id SiteId, input SnapshotRestoreRequest) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/restoreSnapshot", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForRestoreSnapshot sends the RestoreSnapshot request. The method will close the
// http.Response Body if it receives an error.
func (c WebAppsClient) senderForRestoreSnapshot(ctx context.Context, req *http.Request) (future RestoreSnapshotOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}

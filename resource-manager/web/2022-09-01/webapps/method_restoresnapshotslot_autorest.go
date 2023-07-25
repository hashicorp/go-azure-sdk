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

type RestoreSnapshotSlotOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// RestoreSnapshotSlot ...
func (c WebAppsClient) RestoreSnapshotSlot(ctx context.Context, id SlotId, input SnapshotRestoreRequest) (result RestoreSnapshotSlotOperationResponse, err error) {
	req, err := c.preparerForRestoreSnapshotSlot(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "RestoreSnapshotSlot", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForRestoreSnapshotSlot(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "RestoreSnapshotSlot", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// RestoreSnapshotSlotThenPoll performs RestoreSnapshotSlot then polls until it's completed
func (c WebAppsClient) RestoreSnapshotSlotThenPoll(ctx context.Context, id SlotId, input SnapshotRestoreRequest) error {
	result, err := c.RestoreSnapshotSlot(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing RestoreSnapshotSlot: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after RestoreSnapshotSlot: %+v", err)
	}

	return nil
}

// preparerForRestoreSnapshotSlot prepares the RestoreSnapshotSlot request.
func (c WebAppsClient) preparerForRestoreSnapshotSlot(ctx context.Context, id SlotId, input SnapshotRestoreRequest) (*http.Request, error) {
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

// senderForRestoreSnapshotSlot sends the RestoreSnapshotSlot request. The method will close the
// http.Response Body if it receives an error.
func (c WebAppsClient) senderForRestoreSnapshotSlot(ctx context.Context, req *http.Request) (future RestoreSnapshotSlotOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}

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

type RestoreFromBackupBlobSlotOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// RestoreFromBackupBlobSlot ...
func (c WebAppsClient) RestoreFromBackupBlobSlot(ctx context.Context, id SlotId, input RestoreRequest) (result RestoreFromBackupBlobSlotOperationResponse, err error) {
	req, err := c.preparerForRestoreFromBackupBlobSlot(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "RestoreFromBackupBlobSlot", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForRestoreFromBackupBlobSlot(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "RestoreFromBackupBlobSlot", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// RestoreFromBackupBlobSlotThenPoll performs RestoreFromBackupBlobSlot then polls until it's completed
func (c WebAppsClient) RestoreFromBackupBlobSlotThenPoll(ctx context.Context, id SlotId, input RestoreRequest) error {
	result, err := c.RestoreFromBackupBlobSlot(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing RestoreFromBackupBlobSlot: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after RestoreFromBackupBlobSlot: %+v", err)
	}

	return nil
}

// preparerForRestoreFromBackupBlobSlot prepares the RestoreFromBackupBlobSlot request.
func (c WebAppsClient) preparerForRestoreFromBackupBlobSlot(ctx context.Context, id SlotId, input RestoreRequest) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/restoreFromBackupBlob", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForRestoreFromBackupBlobSlot sends the RestoreFromBackupBlobSlot request. The method will close the
// http.Response Body if it receives an error.
func (c WebAppsClient) senderForRestoreFromBackupBlobSlot(ctx context.Context, req *http.Request) (future RestoreFromBackupBlobSlotOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}

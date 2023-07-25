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

type RestoreFromDeletedAppSlotOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// RestoreFromDeletedAppSlot ...
func (c WebAppsClient) RestoreFromDeletedAppSlot(ctx context.Context, id SlotId, input DeletedAppRestoreRequest) (result RestoreFromDeletedAppSlotOperationResponse, err error) {
	req, err := c.preparerForRestoreFromDeletedAppSlot(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "RestoreFromDeletedAppSlot", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForRestoreFromDeletedAppSlot(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "RestoreFromDeletedAppSlot", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// RestoreFromDeletedAppSlotThenPoll performs RestoreFromDeletedAppSlot then polls until it's completed
func (c WebAppsClient) RestoreFromDeletedAppSlotThenPoll(ctx context.Context, id SlotId, input DeletedAppRestoreRequest) error {
	result, err := c.RestoreFromDeletedAppSlot(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing RestoreFromDeletedAppSlot: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after RestoreFromDeletedAppSlot: %+v", err)
	}

	return nil
}

// preparerForRestoreFromDeletedAppSlot prepares the RestoreFromDeletedAppSlot request.
func (c WebAppsClient) preparerForRestoreFromDeletedAppSlot(ctx context.Context, id SlotId, input DeletedAppRestoreRequest) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/restoreFromDeletedApp", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForRestoreFromDeletedAppSlot sends the RestoreFromDeletedAppSlot request. The method will close the
// http.Response Body if it receives an error.
func (c WebAppsClient) senderForRestoreFromDeletedAppSlot(ctx context.Context, req *http.Request) (future RestoreFromDeletedAppSlotOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}

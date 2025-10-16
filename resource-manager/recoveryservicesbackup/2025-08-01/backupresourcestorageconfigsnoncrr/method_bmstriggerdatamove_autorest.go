package backupresourcestorageconfigsnoncrr

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

type BMSTriggerDataMoveOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// BMSTriggerDataMove ...
func (c BackupResourceStorageConfigsNonCRRClient) BMSTriggerDataMove(ctx context.Context, id VaultId, input TriggerDataMoveRequest) (result BMSTriggerDataMoveOperationResponse, err error) {
	req, err := c.preparerForBMSTriggerDataMove(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "backupresourcestorageconfigsnoncrr.BackupResourceStorageConfigsNonCRRClient", "BMSTriggerDataMove", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForBMSTriggerDataMove(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "backupresourcestorageconfigsnoncrr.BackupResourceStorageConfigsNonCRRClient", "BMSTriggerDataMove", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// BMSTriggerDataMoveThenPoll performs BMSTriggerDataMove then polls until it's completed
func (c BackupResourceStorageConfigsNonCRRClient) BMSTriggerDataMoveThenPoll(ctx context.Context, id VaultId, input TriggerDataMoveRequest) error {
	result, err := c.BMSTriggerDataMove(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing BMSTriggerDataMove: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after BMSTriggerDataMove: %+v", err)
	}

	return nil
}

// preparerForBMSTriggerDataMove prepares the BMSTriggerDataMove request.
func (c BackupResourceStorageConfigsNonCRRClient) preparerForBMSTriggerDataMove(ctx context.Context, id VaultId, input TriggerDataMoveRequest) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/backupStorageConfig/vaultStorageConfig/triggerDataMove", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForBMSTriggerDataMove sends the BMSTriggerDataMove request. The method will close the
// http.Response Body if it receives an error.
func (c BackupResourceStorageConfigsNonCRRClient) senderForBMSTriggerDataMove(ctx context.Context, req *http.Request) (future BMSTriggerDataMoveOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}

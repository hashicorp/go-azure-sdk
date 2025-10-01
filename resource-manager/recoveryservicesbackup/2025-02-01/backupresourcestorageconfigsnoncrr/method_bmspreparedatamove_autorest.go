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

type BMSPrepareDataMoveOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// BMSPrepareDataMove ...
func (c BackupResourceStorageConfigsNonCRRClient) BMSPrepareDataMove(ctx context.Context, id VaultId, input PrepareDataMoveRequest) (result BMSPrepareDataMoveOperationResponse, err error) {
	req, err := c.preparerForBMSPrepareDataMove(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "backupresourcestorageconfigsnoncrr.BackupResourceStorageConfigsNonCRRClient", "BMSPrepareDataMove", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForBMSPrepareDataMove(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "backupresourcestorageconfigsnoncrr.BackupResourceStorageConfigsNonCRRClient", "BMSPrepareDataMove", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// BMSPrepareDataMoveThenPoll performs BMSPrepareDataMove then polls until it's completed
func (c BackupResourceStorageConfigsNonCRRClient) BMSPrepareDataMoveThenPoll(ctx context.Context, id VaultId, input PrepareDataMoveRequest) error {
	result, err := c.BMSPrepareDataMove(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing BMSPrepareDataMove: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after BMSPrepareDataMove: %+v", err)
	}

	return nil
}

// preparerForBMSPrepareDataMove prepares the BMSPrepareDataMove request.
func (c BackupResourceStorageConfigsNonCRRClient) preparerForBMSPrepareDataMove(ctx context.Context, id VaultId, input PrepareDataMoveRequest) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/backupStorageConfig/vaultStorageConfig/prepareDataMove", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForBMSPrepareDataMove sends the BMSPrepareDataMove request. The method will close the
// http.Response Body if it receives an error.
func (c BackupResourceStorageConfigsNonCRRClient) senderForBMSPrepareDataMove(ctx context.Context, req *http.Request) (future BMSPrepareDataMoveOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}

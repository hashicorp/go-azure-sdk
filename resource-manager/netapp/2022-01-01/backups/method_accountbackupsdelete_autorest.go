package backups

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

type AccountBackupsDeleteOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// AccountBackupsDelete ...
func (c BackupsClient) AccountBackupsDelete(ctx context.Context, id AccountBackupId) (result AccountBackupsDeleteOperationResponse, err error) {
	req, err := c.preparerForAccountBackupsDelete(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "backups.BackupsClient", "AccountBackupsDelete", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForAccountBackupsDelete(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "backups.BackupsClient", "AccountBackupsDelete", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// AccountBackupsDeleteThenPoll performs AccountBackupsDelete then polls until it's completed
func (c BackupsClient) AccountBackupsDeleteThenPoll(ctx context.Context, id AccountBackupId) error {
	result, err := c.AccountBackupsDelete(ctx, id)
	if err != nil {
		return fmt.Errorf("performing AccountBackupsDelete: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after AccountBackupsDelete: %+v", err)
	}

	return nil
}

// preparerForAccountBackupsDelete prepares the AccountBackupsDelete request.
func (c BackupsClient) preparerForAccountBackupsDelete(ctx context.Context, id AccountBackupId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsDelete(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForAccountBackupsDelete sends the AccountBackupsDelete request. The method will close the
// http.Response Body if it receives an error.
func (c BackupsClient) senderForAccountBackupsDelete(ctx context.Context, req *http.Request) (future AccountBackupsDeleteOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}

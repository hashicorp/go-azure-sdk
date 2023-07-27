package apimanagementservice

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

type BackupOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// Backup ...
func (c ApiManagementServiceClient) Backup(ctx context.Context, id ServiceId, input ApiManagementServiceBackupRestoreParameters) (result BackupOperationResponse, err error) {
	req, err := c.preparerForBackup(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagementservice.ApiManagementServiceClient", "Backup", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForBackup(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagementservice.ApiManagementServiceClient", "Backup", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// BackupThenPoll performs Backup then polls until it's completed
func (c ApiManagementServiceClient) BackupThenPoll(ctx context.Context, id ServiceId, input ApiManagementServiceBackupRestoreParameters) error {
	result, err := c.Backup(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing Backup: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after Backup: %+v", err)
	}

	return nil
}

// preparerForBackup prepares the Backup request.
func (c ApiManagementServiceClient) preparerForBackup(ctx context.Context, id ServiceId, input ApiManagementServiceBackupRestoreParameters) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/backup", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForBackup sends the Backup request. The method will close the
// http.Response Body if it receives an error.
func (c ApiManagementServiceClient) senderForBackup(ctx context.Context, req *http.Request) (future BackupOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}

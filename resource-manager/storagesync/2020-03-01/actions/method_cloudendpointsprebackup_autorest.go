package actions

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

type CloudEndpointsPreBackupOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// CloudEndpointsPreBackup ...
func (c ActionsClient) CloudEndpointsPreBackup(ctx context.Context, id CloudEndpointId, input BackupRequest) (result CloudEndpointsPreBackupOperationResponse, err error) {
	req, err := c.preparerForCloudEndpointsPreBackup(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "actions.ActionsClient", "CloudEndpointsPreBackup", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForCloudEndpointsPreBackup(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "actions.ActionsClient", "CloudEndpointsPreBackup", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// CloudEndpointsPreBackupThenPoll performs CloudEndpointsPreBackup then polls until it's completed
func (c ActionsClient) CloudEndpointsPreBackupThenPoll(ctx context.Context, id CloudEndpointId, input BackupRequest) error {
	result, err := c.CloudEndpointsPreBackup(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing CloudEndpointsPreBackup: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after CloudEndpointsPreBackup: %+v", err)
	}

	return nil
}

// preparerForCloudEndpointsPreBackup prepares the CloudEndpointsPreBackup request.
func (c ActionsClient) preparerForCloudEndpointsPreBackup(ctx context.Context, id CloudEndpointId, input BackupRequest) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/prebackup", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForCloudEndpointsPreBackup sends the CloudEndpointsPreBackup request. The method will close the
// http.Response Body if it receives an error.
func (c ActionsClient) senderForCloudEndpointsPreBackup(ctx context.Context, req *http.Request) (future CloudEndpointsPreBackupOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}

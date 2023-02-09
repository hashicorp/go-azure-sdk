package cloudendpointresource

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

type CloudEndpointsPostBackupOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// CloudEndpointsPostBackup ...
func (c CloudEndpointResourceClient) CloudEndpointsPostBackup(ctx context.Context, id CloudEndpointId, input BackupRequest) (result CloudEndpointsPostBackupOperationResponse, err error) {
	req, err := c.preparerForCloudEndpointsPostBackup(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cloudendpointresource.CloudEndpointResourceClient", "CloudEndpointsPostBackup", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForCloudEndpointsPostBackup(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cloudendpointresource.CloudEndpointResourceClient", "CloudEndpointsPostBackup", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// CloudEndpointsPostBackupThenPoll performs CloudEndpointsPostBackup then polls until it's completed
func (c CloudEndpointResourceClient) CloudEndpointsPostBackupThenPoll(ctx context.Context, id CloudEndpointId, input BackupRequest) error {
	result, err := c.CloudEndpointsPostBackup(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing CloudEndpointsPostBackup: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after CloudEndpointsPostBackup: %+v", err)
	}

	return nil
}

// preparerForCloudEndpointsPostBackup prepares the CloudEndpointsPostBackup request.
func (c CloudEndpointResourceClient) preparerForCloudEndpointsPostBackup(ctx context.Context, id CloudEndpointId, input BackupRequest) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/postbackup", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForCloudEndpointsPostBackup sends the CloudEndpointsPostBackup request. The method will close the
// http.Response Body if it receives an error.
func (c CloudEndpointResourceClient) senderForCloudEndpointsPostBackup(ctx context.Context, req *http.Request) (future CloudEndpointsPostBackupOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}

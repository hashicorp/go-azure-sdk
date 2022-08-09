package incrementalrestorepoints

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

type DiskRestorePointRevokeAccessOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// DiskRestorePointRevokeAccess ...
func (c IncrementalRestorePointsClient) DiskRestorePointRevokeAccess(ctx context.Context, id DiskRestorePointId) (result DiskRestorePointRevokeAccessOperationResponse, err error) {
	req, err := c.preparerForDiskRestorePointRevokeAccess(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "incrementalrestorepoints.IncrementalRestorePointsClient", "DiskRestorePointRevokeAccess", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForDiskRestorePointRevokeAccess(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "incrementalrestorepoints.IncrementalRestorePointsClient", "DiskRestorePointRevokeAccess", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// DiskRestorePointRevokeAccessThenPoll performs DiskRestorePointRevokeAccess then polls until it's completed
func (c IncrementalRestorePointsClient) DiskRestorePointRevokeAccessThenPoll(ctx context.Context, id DiskRestorePointId) error {
	result, err := c.DiskRestorePointRevokeAccess(ctx, id)
	if err != nil {
		return fmt.Errorf("performing DiskRestorePointRevokeAccess: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after DiskRestorePointRevokeAccess: %+v", err)
	}

	return nil
}

// preparerForDiskRestorePointRevokeAccess prepares the DiskRestorePointRevokeAccess request.
func (c IncrementalRestorePointsClient) preparerForDiskRestorePointRevokeAccess(ctx context.Context, id DiskRestorePointId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/endGetAccess", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForDiskRestorePointRevokeAccess sends the DiskRestorePointRevokeAccess request. The method will close the
// http.Response Body if it receives an error.
func (c IncrementalRestorePointsClient) senderForDiskRestorePointRevokeAccess(ctx context.Context, req *http.Request) (future DiskRestorePointRevokeAccessOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}

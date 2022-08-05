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

type DiskRestorePointGrantAccessOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// DiskRestorePointGrantAccess ...
func (c IncrementalRestorePointsClient) DiskRestorePointGrantAccess(ctx context.Context, id DiskRestorePointId, input GrantAccessData) (result DiskRestorePointGrantAccessOperationResponse, err error) {
	req, err := c.preparerForDiskRestorePointGrantAccess(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "incrementalrestorepoints.IncrementalRestorePointsClient", "DiskRestorePointGrantAccess", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForDiskRestorePointGrantAccess(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "incrementalrestorepoints.IncrementalRestorePointsClient", "DiskRestorePointGrantAccess", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// DiskRestorePointGrantAccessThenPoll performs DiskRestorePointGrantAccess then polls until it's completed
func (c IncrementalRestorePointsClient) DiskRestorePointGrantAccessThenPoll(ctx context.Context, id DiskRestorePointId, input GrantAccessData) error {
	result, err := c.DiskRestorePointGrantAccess(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing DiskRestorePointGrantAccess: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after DiskRestorePointGrantAccess: %+v", err)
	}

	return nil
}

// preparerForDiskRestorePointGrantAccess prepares the DiskRestorePointGrantAccess request.
func (c IncrementalRestorePointsClient) preparerForDiskRestorePointGrantAccess(ctx context.Context, id DiskRestorePointId, input GrantAccessData) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/beginGetAccess", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForDiskRestorePointGrantAccess sends the DiskRestorePointGrantAccess request. The method will close the
// http.Response Body if it receives an error.
func (c IncrementalRestorePointsClient) senderForDiskRestorePointGrantAccess(ctx context.Context, req *http.Request) (future DiskRestorePointGrantAccessOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}

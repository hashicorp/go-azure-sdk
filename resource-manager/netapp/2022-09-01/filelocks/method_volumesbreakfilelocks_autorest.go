package filelocks

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

type VolumesBreakFileLocksOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// VolumesBreakFileLocks ...
func (c FileLocksClient) VolumesBreakFileLocks(ctx context.Context, id VolumeId, input BreakFileLocksRequest) (result VolumesBreakFileLocksOperationResponse, err error) {
	req, err := c.preparerForVolumesBreakFileLocks(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "filelocks.FileLocksClient", "VolumesBreakFileLocks", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForVolumesBreakFileLocks(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "filelocks.FileLocksClient", "VolumesBreakFileLocks", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// VolumesBreakFileLocksThenPoll performs VolumesBreakFileLocks then polls until it's completed
func (c FileLocksClient) VolumesBreakFileLocksThenPoll(ctx context.Context, id VolumeId, input BreakFileLocksRequest) error {
	result, err := c.VolumesBreakFileLocks(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing VolumesBreakFileLocks: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after VolumesBreakFileLocks: %+v", err)
	}

	return nil
}

// preparerForVolumesBreakFileLocks prepares the VolumesBreakFileLocks request.
func (c FileLocksClient) preparerForVolumesBreakFileLocks(ctx context.Context, id VolumeId, input BreakFileLocksRequest) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/breakFileLocks", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForVolumesBreakFileLocks sends the VolumesBreakFileLocks request. The method will close the
// http.Response Body if it receives an error.
func (c FileLocksClient) senderForVolumesBreakFileLocks(ctx context.Context, req *http.Request) (future VolumesBreakFileLocksOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}

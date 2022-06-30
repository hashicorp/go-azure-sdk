package volumesrevert

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

type VolumesRevertOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// VolumesRevert ...
func (c VolumesRevertClient) VolumesRevert(ctx context.Context, id VolumeId, input VolumeRevert) (result VolumesRevertOperationResponse, err error) {
	req, err := c.preparerForVolumesRevert(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "volumesrevert.VolumesRevertClient", "VolumesRevert", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForVolumesRevert(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "volumesrevert.VolumesRevertClient", "VolumesRevert", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// VolumesRevertThenPoll performs VolumesRevert then polls until it's completed
func (c VolumesRevertClient) VolumesRevertThenPoll(ctx context.Context, id VolumeId, input VolumeRevert) error {
	result, err := c.VolumesRevert(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing VolumesRevert: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after VolumesRevert: %+v", err)
	}

	return nil
}

// preparerForVolumesRevert prepares the VolumesRevert request.
func (c VolumesRevertClient) preparerForVolumesRevert(ctx context.Context, id VolumeId, input VolumeRevert) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/revert", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForVolumesRevert sends the VolumesRevert request. The method will close the
// http.Response Body if it receives an error.
func (c VolumesRevertClient) senderForVolumesRevert(ctx context.Context, req *http.Request) (future VolumesRevertOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}

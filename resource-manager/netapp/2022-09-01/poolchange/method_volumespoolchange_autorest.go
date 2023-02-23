package poolchange

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

type VolumesPoolChangeOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// VolumesPoolChange ...
func (c PoolChangeClient) VolumesPoolChange(ctx context.Context, id VolumeId, input PoolChangeRequest) (result VolumesPoolChangeOperationResponse, err error) {
	req, err := c.preparerForVolumesPoolChange(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "poolchange.PoolChangeClient", "VolumesPoolChange", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForVolumesPoolChange(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "poolchange.PoolChangeClient", "VolumesPoolChange", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// VolumesPoolChangeThenPoll performs VolumesPoolChange then polls until it's completed
func (c PoolChangeClient) VolumesPoolChangeThenPoll(ctx context.Context, id VolumeId, input PoolChangeRequest) error {
	result, err := c.VolumesPoolChange(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing VolumesPoolChange: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after VolumesPoolChange: %+v", err)
	}

	return nil
}

// preparerForVolumesPoolChange prepares the VolumesPoolChange request.
func (c PoolChangeClient) preparerForVolumesPoolChange(ctx context.Context, id VolumeId, input PoolChangeRequest) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/poolChange", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForVolumesPoolChange sends the VolumesPoolChange request. The method will close the
// http.Response Body if it receives an error.
func (c PoolChangeClient) senderForVolumesPoolChange(ctx context.Context, req *http.Request) (future VolumesPoolChangeOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}

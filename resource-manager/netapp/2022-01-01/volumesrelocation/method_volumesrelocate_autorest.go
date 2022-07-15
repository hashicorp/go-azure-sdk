package volumesrelocation

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

type VolumesRelocateOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// VolumesRelocate ...
func (c VolumesRelocationClient) VolumesRelocate(ctx context.Context, id VolumeId) (result VolumesRelocateOperationResponse, err error) {
	req, err := c.preparerForVolumesRelocate(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "volumesrelocation.VolumesRelocationClient", "VolumesRelocate", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForVolumesRelocate(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "volumesrelocation.VolumesRelocationClient", "VolumesRelocate", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// VolumesRelocateThenPoll performs VolumesRelocate then polls until it's completed
func (c VolumesRelocationClient) VolumesRelocateThenPoll(ctx context.Context, id VolumeId) error {
	result, err := c.VolumesRelocate(ctx, id)
	if err != nil {
		return fmt.Errorf("performing VolumesRelocate: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after VolumesRelocate: %+v", err)
	}

	return nil
}

// preparerForVolumesRelocate prepares the VolumesRelocate request.
func (c VolumesRelocationClient) preparerForVolumesRelocate(ctx context.Context, id VolumeId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/relocate", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForVolumesRelocate sends the VolumesRelocate request. The method will close the
// http.Response Body if it receives an error.
func (c VolumesRelocationClient) senderForVolumesRelocate(ctx context.Context, req *http.Request) (future VolumesRelocateOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}

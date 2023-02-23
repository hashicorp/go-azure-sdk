package resetcifspassword

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

type VolumesResetCifsPasswordOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// VolumesResetCifsPassword ...
func (c ResetCifsPasswordClient) VolumesResetCifsPassword(ctx context.Context, id VolumeId) (result VolumesResetCifsPasswordOperationResponse, err error) {
	req, err := c.preparerForVolumesResetCifsPassword(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resetcifspassword.ResetCifsPasswordClient", "VolumesResetCifsPassword", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForVolumesResetCifsPassword(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resetcifspassword.ResetCifsPasswordClient", "VolumesResetCifsPassword", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// VolumesResetCifsPasswordThenPoll performs VolumesResetCifsPassword then polls until it's completed
func (c ResetCifsPasswordClient) VolumesResetCifsPasswordThenPoll(ctx context.Context, id VolumeId) error {
	result, err := c.VolumesResetCifsPassword(ctx, id)
	if err != nil {
		return fmt.Errorf("performing VolumesResetCifsPassword: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after VolumesResetCifsPassword: %+v", err)
	}

	return nil
}

// preparerForVolumesResetCifsPassword prepares the VolumesResetCifsPassword request.
func (c ResetCifsPasswordClient) preparerForVolumesResetCifsPassword(ctx context.Context, id VolumeId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/resetCifsPassword", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForVolumesResetCifsPassword sends the VolumesResetCifsPassword request. The method will close the
// http.Response Body if it receives an error.
func (c ResetCifsPasswordClient) senderForVolumesResetCifsPassword(ctx context.Context, req *http.Request) (future VolumesResetCifsPasswordOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}

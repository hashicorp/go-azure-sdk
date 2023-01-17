package disks

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

type DetachOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// Detach ...
func (c DisksClient) Detach(ctx context.Context, id DiskId, input DetachDiskProperties) (result DetachOperationResponse, err error) {
	req, err := c.preparerForDetach(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "disks.DisksClient", "Detach", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForDetach(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "disks.DisksClient", "Detach", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// DetachThenPoll performs Detach then polls until it's completed
func (c DisksClient) DetachThenPoll(ctx context.Context, id DiskId, input DetachDiskProperties) error {
	result, err := c.Detach(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing Detach: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after Detach: %+v", err)
	}

	return nil
}

// preparerForDetach prepares the Detach request.
func (c DisksClient) preparerForDetach(ctx context.Context, id DiskId, input DetachDiskProperties) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/detach", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForDetach sends the Detach request. The method will close the
// http.Response Body if it receives an error.
func (c DisksClient) senderForDetach(ctx context.Context, req *http.Request) (future DetachOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}

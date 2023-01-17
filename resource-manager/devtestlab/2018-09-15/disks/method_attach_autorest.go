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

type AttachOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// Attach ...
func (c DisksClient) Attach(ctx context.Context, id DiskId, input AttachDiskProperties) (result AttachOperationResponse, err error) {
	req, err := c.preparerForAttach(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "disks.DisksClient", "Attach", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForAttach(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "disks.DisksClient", "Attach", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// AttachThenPoll performs Attach then polls until it's completed
func (c DisksClient) AttachThenPoll(ctx context.Context, id DiskId, input AttachDiskProperties) error {
	result, err := c.Attach(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing Attach: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after Attach: %+v", err)
	}

	return nil
}

// preparerForAttach prepares the Attach request.
func (c DisksClient) preparerForAttach(ctx context.Context, id DiskId, input AttachDiskProperties) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/attach", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForAttach sends the Attach request. The method will close the
// http.Response Body if it receives an error.
func (c DisksClient) senderForAttach(ctx context.Context, req *http.Request) (future AttachOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}

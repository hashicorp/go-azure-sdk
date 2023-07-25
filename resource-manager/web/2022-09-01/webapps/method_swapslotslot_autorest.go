package webapps

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

type SwapSlotSlotOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// SwapSlotSlot ...
func (c WebAppsClient) SwapSlotSlot(ctx context.Context, id SlotId, input CsmSlotEntity) (result SwapSlotSlotOperationResponse, err error) {
	req, err := c.preparerForSwapSlotSlot(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "SwapSlotSlot", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForSwapSlotSlot(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "SwapSlotSlot", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// SwapSlotSlotThenPoll performs SwapSlotSlot then polls until it's completed
func (c WebAppsClient) SwapSlotSlotThenPoll(ctx context.Context, id SlotId, input CsmSlotEntity) error {
	result, err := c.SwapSlotSlot(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing SwapSlotSlot: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after SwapSlotSlot: %+v", err)
	}

	return nil
}

// preparerForSwapSlotSlot prepares the SwapSlotSlot request.
func (c WebAppsClient) preparerForSwapSlotSlot(ctx context.Context, id SlotId, input CsmSlotEntity) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/slotsswap", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForSwapSlotSlot sends the SwapSlotSlot request. The method will close the
// http.Response Body if it receives an error.
func (c WebAppsClient) senderForSwapSlotSlot(ctx context.Context, req *http.Request) (future SwapSlotSlotOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}

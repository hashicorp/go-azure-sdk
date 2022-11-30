package virtualmachinescalesetrollingupgrades

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

type CancelOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// Cancel ...
func (c VirtualMachineScaleSetRollingUpgradesClient) Cancel(ctx context.Context, id VirtualMachineScaleSetId) (result CancelOperationResponse, err error) {
	req, err := c.preparerForCancel(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "virtualmachinescalesetrollingupgrades.VirtualMachineScaleSetRollingUpgradesClient", "Cancel", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForCancel(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "virtualmachinescalesetrollingupgrades.VirtualMachineScaleSetRollingUpgradesClient", "Cancel", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// CancelThenPoll performs Cancel then polls until it's completed
func (c VirtualMachineScaleSetRollingUpgradesClient) CancelThenPoll(ctx context.Context, id VirtualMachineScaleSetId) error {
	result, err := c.Cancel(ctx, id)
	if err != nil {
		return fmt.Errorf("performing Cancel: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after Cancel: %+v", err)
	}

	return nil
}

// preparerForCancel prepares the Cancel request.
func (c VirtualMachineScaleSetRollingUpgradesClient) preparerForCancel(ctx context.Context, id VirtualMachineScaleSetId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/rollingUpgrades/cancel", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForCancel sends the Cancel request. The method will close the
// http.Response Body if it receives an error.
func (c VirtualMachineScaleSetRollingUpgradesClient) senderForCancel(ctx context.Context, req *http.Request) (future CancelOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}

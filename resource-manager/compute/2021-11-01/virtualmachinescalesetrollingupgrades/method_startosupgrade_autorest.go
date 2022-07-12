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

type StartOSUpgradeOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// StartOSUpgrade ...
func (c VirtualMachineScaleSetRollingUpgradesClient) StartOSUpgrade(ctx context.Context, id VirtualMachineScaleSetId) (result StartOSUpgradeOperationResponse, err error) {
	req, err := c.preparerForStartOSUpgrade(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "virtualmachinescalesetrollingupgrades.VirtualMachineScaleSetRollingUpgradesClient", "StartOSUpgrade", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForStartOSUpgrade(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "virtualmachinescalesetrollingupgrades.VirtualMachineScaleSetRollingUpgradesClient", "StartOSUpgrade", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// StartOSUpgradeThenPoll performs StartOSUpgrade then polls until it's completed
func (c VirtualMachineScaleSetRollingUpgradesClient) StartOSUpgradeThenPoll(ctx context.Context, id VirtualMachineScaleSetId) error {
	result, err := c.StartOSUpgrade(ctx, id)
	if err != nil {
		return fmt.Errorf("performing StartOSUpgrade: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after StartOSUpgrade: %+v", err)
	}

	return nil
}

// preparerForStartOSUpgrade prepares the StartOSUpgrade request.
func (c VirtualMachineScaleSetRollingUpgradesClient) preparerForStartOSUpgrade(ctx context.Context, id VirtualMachineScaleSetId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/osRollingUpgrade", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForStartOSUpgrade sends the StartOSUpgrade request. The method will close the
// http.Response Body if it receives an error.
func (c VirtualMachineScaleSetRollingUpgradesClient) senderForStartOSUpgrade(ctx context.Context, req *http.Request) (future StartOSUpgradeOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}

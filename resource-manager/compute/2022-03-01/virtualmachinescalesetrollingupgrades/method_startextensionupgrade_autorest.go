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

type StartExtensionUpgradeOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// StartExtensionUpgrade ...
func (c VirtualMachineScaleSetRollingUpgradesClient) StartExtensionUpgrade(ctx context.Context, id VirtualMachineScaleSetId) (result StartExtensionUpgradeOperationResponse, err error) {
	req, err := c.preparerForStartExtensionUpgrade(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "virtualmachinescalesetrollingupgrades.VirtualMachineScaleSetRollingUpgradesClient", "StartExtensionUpgrade", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForStartExtensionUpgrade(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "virtualmachinescalesetrollingupgrades.VirtualMachineScaleSetRollingUpgradesClient", "StartExtensionUpgrade", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// StartExtensionUpgradeThenPoll performs StartExtensionUpgrade then polls until it's completed
func (c VirtualMachineScaleSetRollingUpgradesClient) StartExtensionUpgradeThenPoll(ctx context.Context, id VirtualMachineScaleSetId) error {
	result, err := c.StartExtensionUpgrade(ctx, id)
	if err != nil {
		return fmt.Errorf("performing StartExtensionUpgrade: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after StartExtensionUpgrade: %+v", err)
	}

	return nil
}

// preparerForStartExtensionUpgrade prepares the StartExtensionUpgrade request.
func (c VirtualMachineScaleSetRollingUpgradesClient) preparerForStartExtensionUpgrade(ctx context.Context, id VirtualMachineScaleSetId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/extensionRollingUpgrade", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForStartExtensionUpgrade sends the StartExtensionUpgrade request. The method will close the
// http.Response Body if it receives an error.
func (c VirtualMachineScaleSetRollingUpgradesClient) senderForStartExtensionUpgrade(ctx context.Context, req *http.Request) (future StartExtensionUpgradeOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}

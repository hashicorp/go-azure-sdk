package machineextensionsupgrade

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

type UpgradeExtensionsOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// UpgradeExtensions ...
func (c MachineExtensionsUpgradeClient) UpgradeExtensions(ctx context.Context, id MachineId, input MachineExtensionUpgrade) (result UpgradeExtensionsOperationResponse, err error) {
	req, err := c.preparerForUpgradeExtensions(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "machineextensionsupgrade.MachineExtensionsUpgradeClient", "UpgradeExtensions", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForUpgradeExtensions(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "machineextensionsupgrade.MachineExtensionsUpgradeClient", "UpgradeExtensions", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// UpgradeExtensionsThenPoll performs UpgradeExtensions then polls until it's completed
func (c MachineExtensionsUpgradeClient) UpgradeExtensionsThenPoll(ctx context.Context, id MachineId, input MachineExtensionUpgrade) error {
	result, err := c.UpgradeExtensions(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing UpgradeExtensions: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after UpgradeExtensions: %+v", err)
	}

	return nil
}

// preparerForUpgradeExtensions prepares the UpgradeExtensions request.
func (c MachineExtensionsUpgradeClient) preparerForUpgradeExtensions(ctx context.Context, id MachineId, input MachineExtensionUpgrade) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/upgradeExtensions", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForUpgradeExtensions sends the UpgradeExtensions request. The method will close the
// http.Response Body if it receives an error.
func (c MachineExtensionsUpgradeClient) senderForUpgradeExtensions(ctx context.Context, req *http.Request) (future UpgradeExtensionsOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}

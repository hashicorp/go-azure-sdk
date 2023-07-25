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

type InstallSiteExtensionSlotOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// InstallSiteExtensionSlot ...
func (c WebAppsClient) InstallSiteExtensionSlot(ctx context.Context, id SlotSiteExtensionId) (result InstallSiteExtensionSlotOperationResponse, err error) {
	req, err := c.preparerForInstallSiteExtensionSlot(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "InstallSiteExtensionSlot", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForInstallSiteExtensionSlot(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "InstallSiteExtensionSlot", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// InstallSiteExtensionSlotThenPoll performs InstallSiteExtensionSlot then polls until it's completed
func (c WebAppsClient) InstallSiteExtensionSlotThenPoll(ctx context.Context, id SlotSiteExtensionId) error {
	result, err := c.InstallSiteExtensionSlot(ctx, id)
	if err != nil {
		return fmt.Errorf("performing InstallSiteExtensionSlot: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after InstallSiteExtensionSlot: %+v", err)
	}

	return nil
}

// preparerForInstallSiteExtensionSlot prepares the InstallSiteExtensionSlot request.
func (c WebAppsClient) preparerForInstallSiteExtensionSlot(ctx context.Context, id SlotSiteExtensionId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForInstallSiteExtensionSlot sends the InstallSiteExtensionSlot request. The method will close the
// http.Response Body if it receives an error.
func (c WebAppsClient) senderForInstallSiteExtensionSlot(ctx context.Context, req *http.Request) (future InstallSiteExtensionSlotOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}

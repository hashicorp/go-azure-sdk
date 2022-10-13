package appplatform

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

type MonitoringSettingsUpdatePatchOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// MonitoringSettingsUpdatePatch ...
func (c AppPlatformClient) MonitoringSettingsUpdatePatch(ctx context.Context, id SpringId, input MonitoringSettingResource) (result MonitoringSettingsUpdatePatchOperationResponse, err error) {
	req, err := c.preparerForMonitoringSettingsUpdatePatch(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "MonitoringSettingsUpdatePatch", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForMonitoringSettingsUpdatePatch(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "MonitoringSettingsUpdatePatch", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// MonitoringSettingsUpdatePatchThenPoll performs MonitoringSettingsUpdatePatch then polls until it's completed
func (c AppPlatformClient) MonitoringSettingsUpdatePatchThenPoll(ctx context.Context, id SpringId, input MonitoringSettingResource) error {
	result, err := c.MonitoringSettingsUpdatePatch(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing MonitoringSettingsUpdatePatch: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after MonitoringSettingsUpdatePatch: %+v", err)
	}

	return nil
}

// preparerForMonitoringSettingsUpdatePatch prepares the MonitoringSettingsUpdatePatch request.
func (c AppPlatformClient) preparerForMonitoringSettingsUpdatePatch(ctx context.Context, id SpringId, input MonitoringSettingResource) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/monitoringSettings/default", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForMonitoringSettingsUpdatePatch sends the MonitoringSettingsUpdatePatch request. The method will close the
// http.Response Body if it receives an error.
func (c AppPlatformClient) senderForMonitoringSettingsUpdatePatch(ctx context.Context, req *http.Request) (future MonitoringSettingsUpdatePatchOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}

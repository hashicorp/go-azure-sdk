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

type MonitoringSettingsUpdatePutOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// MonitoringSettingsUpdatePut ...
func (c AppPlatformClient) MonitoringSettingsUpdatePut(ctx context.Context, id SpringId, input MonitoringSettingResource) (result MonitoringSettingsUpdatePutOperationResponse, err error) {
	req, err := c.preparerForMonitoringSettingsUpdatePut(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "MonitoringSettingsUpdatePut", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForMonitoringSettingsUpdatePut(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "MonitoringSettingsUpdatePut", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// MonitoringSettingsUpdatePutThenPoll performs MonitoringSettingsUpdatePut then polls until it's completed
func (c AppPlatformClient) MonitoringSettingsUpdatePutThenPoll(ctx context.Context, id SpringId, input MonitoringSettingResource) error {
	result, err := c.MonitoringSettingsUpdatePut(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing MonitoringSettingsUpdatePut: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after MonitoringSettingsUpdatePut: %+v", err)
	}

	return nil
}

// preparerForMonitoringSettingsUpdatePut prepares the MonitoringSettingsUpdatePut request.
func (c AppPlatformClient) preparerForMonitoringSettingsUpdatePut(ctx context.Context, id SpringId, input MonitoringSettingResource) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/monitoringSettings/default", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForMonitoringSettingsUpdatePut sends the MonitoringSettingsUpdatePut request. The method will close the
// http.Response Body if it receives an error.
func (c AppPlatformClient) senderForMonitoringSettingsUpdatePut(ctx context.Context, req *http.Request) (future MonitoringSettingsUpdatePutOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}

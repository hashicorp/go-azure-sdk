package diagnosticsettings

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

type UpdateDiagnosticRemoteSupportSettingsOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// UpdateDiagnosticRemoteSupportSettings ...
func (c DiagnosticSettingsClient) UpdateDiagnosticRemoteSupportSettings(ctx context.Context, id DataBoxEdgeDeviceId, input DiagnosticRemoteSupportSettings) (result UpdateDiagnosticRemoteSupportSettingsOperationResponse, err error) {
	req, err := c.preparerForUpdateDiagnosticRemoteSupportSettings(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "diagnosticsettings.DiagnosticSettingsClient", "UpdateDiagnosticRemoteSupportSettings", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForUpdateDiagnosticRemoteSupportSettings(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "diagnosticsettings.DiagnosticSettingsClient", "UpdateDiagnosticRemoteSupportSettings", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// UpdateDiagnosticRemoteSupportSettingsThenPoll performs UpdateDiagnosticRemoteSupportSettings then polls until it's completed
func (c DiagnosticSettingsClient) UpdateDiagnosticRemoteSupportSettingsThenPoll(ctx context.Context, id DataBoxEdgeDeviceId, input DiagnosticRemoteSupportSettings) error {
	result, err := c.UpdateDiagnosticRemoteSupportSettings(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing UpdateDiagnosticRemoteSupportSettings: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after UpdateDiagnosticRemoteSupportSettings: %+v", err)
	}

	return nil
}

// preparerForUpdateDiagnosticRemoteSupportSettings prepares the UpdateDiagnosticRemoteSupportSettings request.
func (c DiagnosticSettingsClient) preparerForUpdateDiagnosticRemoteSupportSettings(ctx context.Context, id DataBoxEdgeDeviceId, input DiagnosticRemoteSupportSettings) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/diagnosticRemoteSupportSettings/default", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForUpdateDiagnosticRemoteSupportSettings sends the UpdateDiagnosticRemoteSupportSettings request. The method will close the
// http.Response Body if it receives an error.
func (c DiagnosticSettingsClient) senderForUpdateDiagnosticRemoteSupportSettings(ctx context.Context, req *http.Request) (future UpdateDiagnosticRemoteSupportSettingsOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}

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

type UpdateDiagnosticProactiveLogCollectionSettingsOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// UpdateDiagnosticProactiveLogCollectionSettings ...
func (c DiagnosticSettingsClient) UpdateDiagnosticProactiveLogCollectionSettings(ctx context.Context, id DataBoxEdgeDeviceId, input DiagnosticProactiveLogCollectionSettings) (result UpdateDiagnosticProactiveLogCollectionSettingsOperationResponse, err error) {
	req, err := c.preparerForUpdateDiagnosticProactiveLogCollectionSettings(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "diagnosticsettings.DiagnosticSettingsClient", "UpdateDiagnosticProactiveLogCollectionSettings", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForUpdateDiagnosticProactiveLogCollectionSettings(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "diagnosticsettings.DiagnosticSettingsClient", "UpdateDiagnosticProactiveLogCollectionSettings", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// UpdateDiagnosticProactiveLogCollectionSettingsThenPoll performs UpdateDiagnosticProactiveLogCollectionSettings then polls until it's completed
func (c DiagnosticSettingsClient) UpdateDiagnosticProactiveLogCollectionSettingsThenPoll(ctx context.Context, id DataBoxEdgeDeviceId, input DiagnosticProactiveLogCollectionSettings) error {
	result, err := c.UpdateDiagnosticProactiveLogCollectionSettings(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing UpdateDiagnosticProactiveLogCollectionSettings: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after UpdateDiagnosticProactiveLogCollectionSettings: %+v", err)
	}

	return nil
}

// preparerForUpdateDiagnosticProactiveLogCollectionSettings prepares the UpdateDiagnosticProactiveLogCollectionSettings request.
func (c DiagnosticSettingsClient) preparerForUpdateDiagnosticProactiveLogCollectionSettings(ctx context.Context, id DataBoxEdgeDeviceId, input DiagnosticProactiveLogCollectionSettings) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/diagnosticProactiveLogCollectionSettings/default", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForUpdateDiagnosticProactiveLogCollectionSettings sends the UpdateDiagnosticProactiveLogCollectionSettings request. The method will close the
// http.Response Body if it receives an error.
func (c DiagnosticSettingsClient) senderForUpdateDiagnosticProactiveLogCollectionSettings(ctx context.Context, req *http.Request) (future UpdateDiagnosticProactiveLogCollectionSettingsOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}

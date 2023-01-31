package diagnosticsettings

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetDiagnosticRemoteSupportSettingsOperationResponse struct {
	HttpResponse *http.Response
	Model        *DiagnosticRemoteSupportSettings
}

// GetDiagnosticRemoteSupportSettings ...
func (c DiagnosticSettingsClient) GetDiagnosticRemoteSupportSettings(ctx context.Context, id DataBoxEdgeDeviceId) (result GetDiagnosticRemoteSupportSettingsOperationResponse, err error) {
	req, err := c.preparerForGetDiagnosticRemoteSupportSettings(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "diagnosticsettings.DiagnosticSettingsClient", "GetDiagnosticRemoteSupportSettings", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "diagnosticsettings.DiagnosticSettingsClient", "GetDiagnosticRemoteSupportSettings", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetDiagnosticRemoteSupportSettings(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "diagnosticsettings.DiagnosticSettingsClient", "GetDiagnosticRemoteSupportSettings", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetDiagnosticRemoteSupportSettings prepares the GetDiagnosticRemoteSupportSettings request.
func (c DiagnosticSettingsClient) preparerForGetDiagnosticRemoteSupportSettings(ctx context.Context, id DataBoxEdgeDeviceId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/diagnosticRemoteSupportSettings/default", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForGetDiagnosticRemoteSupportSettings handles the response to the GetDiagnosticRemoteSupportSettings request. The method always
// closes the http.Response Body.
func (c DiagnosticSettingsClient) responderForGetDiagnosticRemoteSupportSettings(resp *http.Response) (result GetDiagnosticRemoteSupportSettingsOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

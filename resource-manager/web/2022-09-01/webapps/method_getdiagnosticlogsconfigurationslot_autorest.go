package webapps

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetDiagnosticLogsConfigurationSlotOperationResponse struct {
	HttpResponse *http.Response
	Model        *SiteLogsConfig
}

// GetDiagnosticLogsConfigurationSlot ...
func (c WebAppsClient) GetDiagnosticLogsConfigurationSlot(ctx context.Context, id SlotId) (result GetDiagnosticLogsConfigurationSlotOperationResponse, err error) {
	req, err := c.preparerForGetDiagnosticLogsConfigurationSlot(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetDiagnosticLogsConfigurationSlot", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetDiagnosticLogsConfigurationSlot", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetDiagnosticLogsConfigurationSlot(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetDiagnosticLogsConfigurationSlot", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetDiagnosticLogsConfigurationSlot prepares the GetDiagnosticLogsConfigurationSlot request.
func (c WebAppsClient) preparerForGetDiagnosticLogsConfigurationSlot(ctx context.Context, id SlotId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/config/logs", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForGetDiagnosticLogsConfigurationSlot handles the response to the GetDiagnosticLogsConfigurationSlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForGetDiagnosticLogsConfigurationSlot(resp *http.Response) (result GetDiagnosticLogsConfigurationSlotOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

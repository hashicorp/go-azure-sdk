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

type GetDiagnosticLogsConfigurationOperationResponse struct {
	HttpResponse *http.Response
	Model        *SiteLogsConfig
}

// GetDiagnosticLogsConfiguration ...
func (c WebAppsClient) GetDiagnosticLogsConfiguration(ctx context.Context, id SiteId) (result GetDiagnosticLogsConfigurationOperationResponse, err error) {
	req, err := c.preparerForGetDiagnosticLogsConfiguration(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetDiagnosticLogsConfiguration", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetDiagnosticLogsConfiguration", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetDiagnosticLogsConfiguration(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetDiagnosticLogsConfiguration", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetDiagnosticLogsConfiguration prepares the GetDiagnosticLogsConfiguration request.
func (c WebAppsClient) preparerForGetDiagnosticLogsConfiguration(ctx context.Context, id SiteId) (*http.Request, error) {
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

// responderForGetDiagnosticLogsConfiguration handles the response to the GetDiagnosticLogsConfiguration request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForGetDiagnosticLogsConfiguration(resp *http.Response) (result GetDiagnosticLogsConfigurationOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

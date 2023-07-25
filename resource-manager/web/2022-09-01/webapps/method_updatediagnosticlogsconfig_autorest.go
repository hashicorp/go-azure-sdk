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

type UpdateDiagnosticLogsConfigOperationResponse struct {
	HttpResponse *http.Response
	Model        *SiteLogsConfig
}

// UpdateDiagnosticLogsConfig ...
func (c WebAppsClient) UpdateDiagnosticLogsConfig(ctx context.Context, id SiteId, input SiteLogsConfig) (result UpdateDiagnosticLogsConfigOperationResponse, err error) {
	req, err := c.preparerForUpdateDiagnosticLogsConfig(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "UpdateDiagnosticLogsConfig", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "UpdateDiagnosticLogsConfig", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForUpdateDiagnosticLogsConfig(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "UpdateDiagnosticLogsConfig", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForUpdateDiagnosticLogsConfig prepares the UpdateDiagnosticLogsConfig request.
func (c WebAppsClient) preparerForUpdateDiagnosticLogsConfig(ctx context.Context, id SiteId, input SiteLogsConfig) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/config/logs", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForUpdateDiagnosticLogsConfig handles the response to the UpdateDiagnosticLogsConfig request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForUpdateDiagnosticLogsConfig(resp *http.Response) (result UpdateDiagnosticLogsConfigOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

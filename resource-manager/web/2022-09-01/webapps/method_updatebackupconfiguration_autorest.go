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

type UpdateBackupConfigurationOperationResponse struct {
	HttpResponse *http.Response
	Model        *BackupRequest
}

// UpdateBackupConfiguration ...
func (c WebAppsClient) UpdateBackupConfiguration(ctx context.Context, id SiteId, input BackupRequest) (result UpdateBackupConfigurationOperationResponse, err error) {
	req, err := c.preparerForUpdateBackupConfiguration(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "UpdateBackupConfiguration", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "UpdateBackupConfiguration", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForUpdateBackupConfiguration(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "UpdateBackupConfiguration", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForUpdateBackupConfiguration prepares the UpdateBackupConfiguration request.
func (c WebAppsClient) preparerForUpdateBackupConfiguration(ctx context.Context, id SiteId, input BackupRequest) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/config/backup", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForUpdateBackupConfiguration handles the response to the UpdateBackupConfiguration request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForUpdateBackupConfiguration(resp *http.Response) (result UpdateBackupConfigurationOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

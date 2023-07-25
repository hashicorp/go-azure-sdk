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

type GetBackupConfigurationOperationResponse struct {
	HttpResponse *http.Response
	Model        *BackupRequest
}

// GetBackupConfiguration ...
func (c WebAppsClient) GetBackupConfiguration(ctx context.Context, id SiteId) (result GetBackupConfigurationOperationResponse, err error) {
	req, err := c.preparerForGetBackupConfiguration(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetBackupConfiguration", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetBackupConfiguration", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetBackupConfiguration(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetBackupConfiguration", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetBackupConfiguration prepares the GetBackupConfiguration request.
func (c WebAppsClient) preparerForGetBackupConfiguration(ctx context.Context, id SiteId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/config/backup/list", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForGetBackupConfiguration handles the response to the GetBackupConfiguration request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForGetBackupConfiguration(resp *http.Response) (result GetBackupConfigurationOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

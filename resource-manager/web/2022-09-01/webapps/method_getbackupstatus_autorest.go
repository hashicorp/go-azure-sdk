package webapps

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetBackupStatusOperationResponse struct {
	HttpResponse *http.Response
	Model        *BackupItem
}

// GetBackupStatus ...
func (c WebAppsClient) GetBackupStatus(ctx context.Context, id BackupId) (result GetBackupStatusOperationResponse, err error) {
	req, err := c.preparerForGetBackupStatus(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetBackupStatus", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetBackupStatus", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetBackupStatus(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetBackupStatus", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetBackupStatus prepares the GetBackupStatus request.
func (c WebAppsClient) preparerForGetBackupStatus(ctx context.Context, id BackupId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForGetBackupStatus handles the response to the GetBackupStatus request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForGetBackupStatus(resp *http.Response) (result GetBackupStatusOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

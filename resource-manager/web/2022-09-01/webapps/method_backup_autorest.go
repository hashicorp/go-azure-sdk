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

type BackupOperationResponse struct {
	HttpResponse *http.Response
	Model        *BackupItem
}

// Backup ...
func (c WebAppsClient) Backup(ctx context.Context, id SiteId, input BackupRequest) (result BackupOperationResponse, err error) {
	req, err := c.preparerForBackup(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "Backup", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "Backup", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForBackup(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "Backup", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForBackup prepares the Backup request.
func (c WebAppsClient) preparerForBackup(ctx context.Context, id SiteId, input BackupRequest) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/backup", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForBackup handles the response to the Backup request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForBackup(resp *http.Response) (result BackupOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

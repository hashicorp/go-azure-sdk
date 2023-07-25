package webapps

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeleteBackupOperationResponse struct {
	HttpResponse *http.Response
}

// DeleteBackup ...
func (c WebAppsClient) DeleteBackup(ctx context.Context, id BackupId) (result DeleteBackupOperationResponse, err error) {
	req, err := c.preparerForDeleteBackup(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "DeleteBackup", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "DeleteBackup", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForDeleteBackup(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "DeleteBackup", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForDeleteBackup prepares the DeleteBackup request.
func (c WebAppsClient) preparerForDeleteBackup(ctx context.Context, id BackupId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsDelete(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForDeleteBackup handles the response to the DeleteBackup request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForDeleteBackup(resp *http.Response) (result DeleteBackupOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

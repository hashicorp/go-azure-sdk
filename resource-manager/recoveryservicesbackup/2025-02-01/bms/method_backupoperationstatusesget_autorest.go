package bms

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BackupOperationStatusesGetOperationResponse struct {
	HttpResponse *http.Response
	Model        *OperationStatus
}

// BackupOperationStatusesGet ...
func (c BmsClient) BackupOperationStatusesGet(ctx context.Context, id BackupOperationId) (result BackupOperationStatusesGetOperationResponse, err error) {
	req, err := c.preparerForBackupOperationStatusesGet(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "bms.BmsClient", "BackupOperationStatusesGet", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "bms.BmsClient", "BackupOperationStatusesGet", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForBackupOperationStatusesGet(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "bms.BmsClient", "BackupOperationStatusesGet", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForBackupOperationStatusesGet prepares the BackupOperationStatusesGet request.
func (c BmsClient) preparerForBackupOperationStatusesGet(ctx context.Context, id BackupOperationId) (*http.Request, error) {
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

// responderForBackupOperationStatusesGet handles the response to the BackupOperationStatusesGet request. The method always
// closes the http.Response Body.
func (c BmsClient) responderForBackupOperationStatusesGet(resp *http.Response) (result BackupOperationStatusesGetOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

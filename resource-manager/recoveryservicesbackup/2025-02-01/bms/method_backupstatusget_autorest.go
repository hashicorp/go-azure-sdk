package bms

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BackupStatusGetOperationResponse struct {
	HttpResponse *http.Response
	Model        *BackupStatusResponse
}

// BackupStatusGet ...
func (c BmsClient) BackupStatusGet(ctx context.Context, id LocationId, input BackupStatusRequest) (result BackupStatusGetOperationResponse, err error) {
	req, err := c.preparerForBackupStatusGet(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "bms.BmsClient", "BackupStatusGet", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "bms.BmsClient", "BackupStatusGet", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForBackupStatusGet(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "bms.BmsClient", "BackupStatusGet", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForBackupStatusGet prepares the BackupStatusGet request.
func (c BmsClient) preparerForBackupStatusGet(ctx context.Context, id LocationId, input BackupStatusRequest) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/backupStatus", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForBackupStatusGet handles the response to the BackupStatusGet request. The method always
// closes the http.Response Body.
func (c BmsClient) responderForBackupStatusGet(resp *http.Response) (result BackupStatusGetOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

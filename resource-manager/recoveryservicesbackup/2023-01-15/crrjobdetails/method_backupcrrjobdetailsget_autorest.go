package crrjobdetails

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BackupCrrJobDetailsGetOperationResponse struct {
	HttpResponse *http.Response
	Model        *JobResource
}

// BackupCrrJobDetailsGet ...
func (c CrrJobDetailsClient) BackupCrrJobDetailsGet(ctx context.Context, id LocationId, input CrrJobRequest) (result BackupCrrJobDetailsGetOperationResponse, err error) {
	req, err := c.preparerForBackupCrrJobDetailsGet(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "crrjobdetails.CrrJobDetailsClient", "BackupCrrJobDetailsGet", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "crrjobdetails.CrrJobDetailsClient", "BackupCrrJobDetailsGet", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForBackupCrrJobDetailsGet(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "crrjobdetails.CrrJobDetailsClient", "BackupCrrJobDetailsGet", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForBackupCrrJobDetailsGet prepares the BackupCrrJobDetailsGet request.
func (c CrrJobDetailsClient) preparerForBackupCrrJobDetailsGet(ctx context.Context, id LocationId, input CrrJobRequest) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/backupCrrJob", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForBackupCrrJobDetailsGet handles the response to the BackupCrrJobDetailsGet request. The method always
// closes the http.Response Body.
func (c CrrJobDetailsClient) responderForBackupCrrJobDetailsGet(resp *http.Response) (result BackupCrrJobDetailsGetOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

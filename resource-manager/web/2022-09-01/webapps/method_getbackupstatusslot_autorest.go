package webapps

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetBackupStatusSlotOperationResponse struct {
	HttpResponse *http.Response
	Model        *BackupItem
}

// GetBackupStatusSlot ...
func (c WebAppsClient) GetBackupStatusSlot(ctx context.Context, id SlotBackupId) (result GetBackupStatusSlotOperationResponse, err error) {
	req, err := c.preparerForGetBackupStatusSlot(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetBackupStatusSlot", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetBackupStatusSlot", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetBackupStatusSlot(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetBackupStatusSlot", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetBackupStatusSlot prepares the GetBackupStatusSlot request.
func (c WebAppsClient) preparerForGetBackupStatusSlot(ctx context.Context, id SlotBackupId) (*http.Request, error) {
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

// responderForGetBackupStatusSlot handles the response to the GetBackupStatusSlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForGetBackupStatusSlot(resp *http.Response) (result GetBackupStatusSlotOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

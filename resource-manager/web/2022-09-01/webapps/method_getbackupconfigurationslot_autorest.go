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

type GetBackupConfigurationSlotOperationResponse struct {
	HttpResponse *http.Response
	Model        *BackupRequest
}

// GetBackupConfigurationSlot ...
func (c WebAppsClient) GetBackupConfigurationSlot(ctx context.Context, id SlotId) (result GetBackupConfigurationSlotOperationResponse, err error) {
	req, err := c.preparerForGetBackupConfigurationSlot(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetBackupConfigurationSlot", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetBackupConfigurationSlot", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetBackupConfigurationSlot(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetBackupConfigurationSlot", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetBackupConfigurationSlot prepares the GetBackupConfigurationSlot request.
func (c WebAppsClient) preparerForGetBackupConfigurationSlot(ctx context.Context, id SlotId) (*http.Request, error) {
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

// responderForGetBackupConfigurationSlot handles the response to the GetBackupConfigurationSlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForGetBackupConfigurationSlot(resp *http.Response) (result GetBackupConfigurationSlotOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

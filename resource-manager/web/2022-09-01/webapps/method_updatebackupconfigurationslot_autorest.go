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

type UpdateBackupConfigurationSlotOperationResponse struct {
	HttpResponse *http.Response
	Model        *BackupRequest
}

// UpdateBackupConfigurationSlot ...
func (c WebAppsClient) UpdateBackupConfigurationSlot(ctx context.Context, id SlotId, input BackupRequest) (result UpdateBackupConfigurationSlotOperationResponse, err error) {
	req, err := c.preparerForUpdateBackupConfigurationSlot(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "UpdateBackupConfigurationSlot", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "UpdateBackupConfigurationSlot", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForUpdateBackupConfigurationSlot(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "UpdateBackupConfigurationSlot", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForUpdateBackupConfigurationSlot prepares the UpdateBackupConfigurationSlot request.
func (c WebAppsClient) preparerForUpdateBackupConfigurationSlot(ctx context.Context, id SlotId, input BackupRequest) (*http.Request, error) {
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

// responderForUpdateBackupConfigurationSlot handles the response to the UpdateBackupConfigurationSlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForUpdateBackupConfigurationSlot(resp *http.Response) (result UpdateBackupConfigurationSlotOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
